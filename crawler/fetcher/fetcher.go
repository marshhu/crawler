package fetcher

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"log"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"golang.org/x/net/html/charset"
)

//Fetch to get url body data
func Fetch(url string) ([]byte,error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil,err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil,fmt.Errorf("wrong status code: d%",resp.StatusCode)
	}
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes,"")
	return e
}
