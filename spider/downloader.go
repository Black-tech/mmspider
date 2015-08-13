package spider

import (
	"io"
	"net/http"
	"os"
)

func get(url string) (content []byte, err error) {
	res, err := http.Get(url)
	if err != nil {
		return
	}
	defer res.Body.Close()
	content, err = ioutil.ReadAll(res.Body)
	return
}

func save() {

}
