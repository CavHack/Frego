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

func
