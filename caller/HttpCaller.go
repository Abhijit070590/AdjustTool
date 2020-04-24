package caller

import (
	"io/ioutil"
	"log"
	"net/http"
)

type HttpCaller struct {
}

func (c *HttpCaller) Call(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	return body
}
