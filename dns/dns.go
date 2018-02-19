package dns

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/mercadolibre/golang-restclient/rest"
)

//DNSURL is the url for DNS status
const DNSURL string = "http://localhost:8080/status,http://localhost:8080/status"

//DNS JSON model
type DNS struct {
	Hostname string  `json:"hostname,omitempty"`
	FreeMem  uint64  `json:"freemem,omitempty"`
	Average  float64 `json:"cpuUsage,omitempty"`
	DNSAlive string  `json:"dnsalive,omitempty"`
}

//GetDNS returns DNS status JSON
func GetDNS(w http.ResponseWriter, r *http.Request) {
	var dnsGroup []DNS
	var url string

	//Check environment variable
	if os.Getenv("DNSURL") != "" {
		url = os.Getenv("DNSURL")
	} else {
		url = DNSURL
	}
	//Creates a list from the url variable
	list := strings.Split(url, ",")
	var f [2]*rest.FutureResponse

	//Makes the requests
	rest.ForkJoin(func(c *rest.Concurrent) {
		f[0] = c.Get(list[0])
		f[1] = c.Get(list[1])
	})

	w.Header().Set("Content-Type", "application/json")

	//Save every response from the previous requests and convert it in JSON format
	for i := range f {
		if f[i].Response().StatusCode == http.StatusOK {
			res := f[i].Response().String()
			b := []byte(res)
			var dns DNS
			err := json.Unmarshal(b, &dns)
			if err != nil {
				log.Fatal(err)
			}
			dnsGroup = append(dnsGroup, dns)
		}
	}
	json.NewEncoder(w).Encode(dnsGroup)
}
