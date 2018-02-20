package iaas

import (
	"log"
	"net/http"
	"os"

	resty "gopkg.in/resty.v1"
)

//URL const variable
const URL string = "https://jsonplaceholder.typicode.com/posts"

//GetIaas returns the JSON object from URL variable
func GetIaas(w http.ResponseWriter, r *http.Request) {
	var url string
	//Check environment variable
	if os.Getenv("IAAS_URL") != "" {
		url = os.Getenv("IAAS_URL")
	} else {
		url = URL
	}
	w.Header().Set("Content-Type", "application/json")
	resp, err := resty.R().Get(url)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(resp.Body())
}
