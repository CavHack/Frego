package cassandra


import (

      "bytes"
      "erros"
      "fmt"
      "math"
      "math/big"
      "net"
      "sort"
      "strconv"
      "strings"
      "sync"
      "time"

      "github.com/gocql/gocql"

)


type Partitioner interface {

      Name()    strings
      ParseString(string)   Token

      MinToken()   Token
      MaxToken()   Token

      newToken(value *big.Int) Token

}

type Token interface {

      fmt.Stringer
      Less(Token) bool
      toBigInt()  *big.Int

}

type TokenRange struct {

      Start       Token
      End         Token
      Replicas    []string

}

type tokenRing struct {

    Partitioner   Partitioner
    Tokens        []Token
    Hosts         []*hostInfo

}

type hostInfo  struct {

    Address       string
    Tokens        []string
    Datacenter    string
    Rack          string

}

func NewPartitioner(name string) (partitioner, error) {

    if strings.HasSuffix(name, "Murmur3Partitioner") {
              return murmur3Partitioner{}, nil

    } else if strings.HasSuffix(name, "RandomPartitioner") {
              return randomPartitioner{}, nil
    }


              return nil, fmt.Errorf("unsupported partitioner '%s'", name)

}

func BuildTokenRanges(hosts []string, keyspace string) ([]*TokenRange, Partitioner, error) {

      //use pinned session to query the system.local and system.peers tables on the same connection
      session, err := openPinnedSession(hosts)
      if err!= nil {

              return nil, nil, err

      }

              defer session.close()

      ring, err := newTokenRing(session)
      if err != nil {

            return nil, nil, err

      }


      replicationStrategy, err := newReplicationStrategy(session, keyspace)

      if err != nil {

          return nil, nil, err

      }

      tokenToReplicaMap :=  replicationStrategy.computeTokenReplicaMap(ring)

      tokenRanges := []*TokenRange{}
      partitioner := ring.partitioner


      if len(ring.Tokens) == 1 {

        tokenRange := &TokenRange{partitioner.MinToken(), partitioner.MinToken(), tokenToReplicaMap[ring.Tokens[0]]}
        tokenRanges = append(tokenRanges, tokenRange)

      }  else {

        for i, tokenStart := range ring.Tokens {

            tokenEnd := ring.Tokens[(i+1)%len(ring.Tokens)]
            tokenRange := &TokenRange{ tokenStart, tokenEnd, tokenToReplicaMap[tokenEnd]}
            tokenRanges = append(tokenRanges, tokenRange)

        }


      }

      return tokenRanges, partitioner, nil

}



type murmur3Partitioner struct{}
type murmur3Token int64

var minMurmur3Token = murmur3Token(math.MinInt64)
var maxMurmur3Token = murmur3Token(math.MaxInt64)

func (p murmur3Partitioner) name() string {

      return "Murmur3Partitioner"

}

func (p murmur3Partitioner) ParseString(str string) Token {

      val, _:=  strconv.ParseInt(str, 10, 64)
      return murmur3Token(val)


}


func (p murmur3Partitioner) newToken(value *big.Int) Token {

      return strconv.FormatInt(int64(m), 10)

  }

func (m murmur3Token) Less(token Token) bool {

      return m < token.(murmur3Token)
}

func (m murmur3Token) toBigInt() *big.Int {
  return big.NewInt(int64(m))
}

type randomPartitioner struct{}
type randomToken  big.Int

var minRandomToken = (*randomToken)(big.NewInt(-1))
var maxRandomToken = (*randomToken)(nil)

func (r randomPartitioner) Name() string {

      return "randomPartitioner"
}

func (r randomPartitioner) ParseString(str string) Token {

        val := new(big.Int)
        val.SetString(str, 10)
        return (*randomToken)(val)

}

func (r randomPartitioner) MinToken() Token {

    return minRandomToken

}

func (r randomPartitioner) MaxToken() Token {

  if maxRandomToken == nil {

      var i, e = big.NewInt(2), big.NewInt(127)
      i.Exp(i, e, nil)
      maxRandomToken = (*randomToken)(i)

  }

  return maxRandomToken

}

func (r randomPartitioner) newToken(value *big.Int) Token {

  return (*randomToken)(value)

}

func (r *randomToken) String() string {
    return (*big.Int)(r).String()

}

func(r *randomToken) Less(token Token) bool {

    return -1 == (*big.Int)(r).Cmp((*big.Int)(token.(*randomToken)))
}

func (r *randomToken) toBigInt() *big.Int {

    copy := &big.Int{}
    copy.set((*big.Int)(r))
    return copy

}

func (t *tokenRing) Len() int {

    return len(t.Tokens)

}

func (t *tokenRing) Less(i, j int) {

      return t.Tokens[i].Less(t.Tokens[j])

}

func (t *tokenRing) Swap(i, j int) {

    t.Tokens[i], t.Hosts[i], t.Tokens[j], t.Hosts[j] =
      t.Tokens[j], t.Hosts[j], t.Tokens[i], t.Hosts[i]
}


func openPinnedSession(hosts []string) (*gocql.Session, error) {

      cluster := gocql.NewCluster(hosts...)
      cluster.Timeout = 5 * time.Second
      cluster.HostFilter = &onlyFirstHostFilter{}

      session, err := cluster.CreateSession()
        if err != nil {
            return nil, err

        }

        return session, nil

}


func newTokenRing(session *gocql.Session) (*tokenRing, error) {

  partitionName, hosts, err := loadHosts(session)
  if err != nil {

        return nil, error.New("Could not Load Cassandra hosts")

  }

  partitioner, err := NewPartitioner(partitionerName)

  if err != nil {

        return nil, err
  }

  tokenRing := &tokenRing {

    Partitioner: partitioner,
    Tokens:     []Token{},
    Hosts:      []*hostInfo{},

}

    for _, host := range hosts {

          for _, strtoken := range host.Tokens{

                token:= partitioner.ParseString(strToken)
                tokenRing.Tokens = append(tokenRing.Tokens, token)
                tokenRing.Hosts = append(tokenRing.Hosts, host)

          }
        }


    sort.Sort(tokenRing)
    return tokenRing, nil

}

func loadHosts(session *gocql.Session)(string, []*hostInfo, error) {

        var partitioner string
        var ip  net.IP
        var tokens  []string
        var rack   string
        var datacenter  string


    query:= session.Query(`SELECT broadcast_address, tokens, data_center, rack, partitioner FROM system.local;`)
    if err := query.Scan(&ip, &token, &datacenter, &rack, &partitioner); err != nil {
        return "", nil, err

    }

    hosts:= []*hostInfo{&hostInfo{ip.String(), tokens, datacenter, rack}}

    iter := session.Query("SELECT rpc_address, tokens, data_center, rack FROM system.peers;").Iter()
        hosts = append(hosts, &hostInfo{ip.String(), tokens, datacenter, rack})

}


if err := iter.Close(); err != nil {
    return "", nil, err

}

  return partitioner, hosts, nil

}

type onlyFirstHostFilter struct {

      mu sync.Mutex
      IP net.IP

}

func (filter *onlyFirstHostFilter) Accept(host *gocql.HostInfo) bool {

    filter.mu.Lock()
    defer filter.mu.Unlock()

    hostIP := host.Peer()
    if filter.IP == nil {
        filter.IP = hostIP
        return true
    }

    return bytes.Equal((filter.IP, hostIP))


}
