//Frego

package Frego


import (
	
	"crypto/x509"
	"crypto/tls"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"

)


var errorCounter = prometheus.NewCounter(prometheus.CounterOpts {


		Namespace: "mesos",
		Subsystem: "collector",
		Name:	   "errors_total",
		Help:	   "Total Number of internal mesos-collector errors.",




	})

func init() {

}



func mkHttpClient(url string, timeout time.Duration, auth authInfo, certPool *x509.CertPool ) *httpClient {

	transport := &http.Transport{
			TLSClientConfig: &tls.Config{RootCAs: certPool},
	
	}


	var redirectFunc func(req *http.Request, via[]*http.Request) error 

	if auth.username != ""  && auth.password != "" {

		//Auth information is only available in the current context -> use lambda
		redirectFunc = func(req *http.Request, via[]*http.Request) error {

			req.SetBasicAuth(auth.username, auth.password)
			return nil

		}

	}


	return &httpClient {

			http.Client{Timeout: timeout, Transport: transport, CheckRedirect: redirectFunc}, 
			url,
			auth,

	}

}

func csvInputToList(input string)[]string {

	var entryList []string
	if input == "" {
		return entryList
	}

	sanitizedString := strings.Replace(input, "", "", -1)
	entryList = strings.Split(sanitizedString, ",")
	return entryList



}

func main() {

	masterURL := fs.String("master", "", "Expose metrics from master running on this URL")
	slaveURL  := fs.String("slave", "", "Expose metrics from slave running on this URL")
	fs := flag.NewFlagSet("Frego", flag.ExitOnError)
	timeout := fs.Duration("timeout", 5*time.Second, "Master polling timeout" )
	exportedTaskLabels = fs.String("exportedTaskLabels", "", "Comma separated list of task labels to include the corresponding metric")


}