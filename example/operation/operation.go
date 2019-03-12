package operation

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Get

func GetOperation(url string) {
	request, _ := http.NewRequest("GET", url, nil)
	client := http.DefaultClient
	response, err := client.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	byte, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(byte))
}

// Post
func PostOperation(path string) {
}

// PostForm

// Url
