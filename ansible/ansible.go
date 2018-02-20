package ansible

import (
	"log"
	"net/http"
	"os"

	"gopkg.in/resty.v1"
)

//URL variable
const URL string = "https://jsonplaceholder.typicode.com/posts"

//GetAnsible returns the JSON object from URL variable
func GetAnsible(w http.ResponseWriter, r *http.Request) {
	var url string
	//Check environment variable
	if os.Getenv("ANSIBLE_URL") != "" {
		url = os.Getenv("ANSIBLE_URL")
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
