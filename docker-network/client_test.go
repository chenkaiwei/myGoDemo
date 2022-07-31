package docker_network

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestClient(t *testing.T) {

	resp, err := http.Get("http://localhost:1122")
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}
