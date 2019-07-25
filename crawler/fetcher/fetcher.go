package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	resp, e := http.Get(url)
	if e != nil {
		return nil, e
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code %d",
			resp.StatusCode)
	}
	bodyReader := bufio.NewReader(resp.Body)
	i := determineEncoding(bodyReader)
	reader := transform.NewReader(bodyReader, i.NewDecoder())
	return ioutil.ReadAll(reader)
}

//gopm get -g -v golang.org/x/text
//go get -v github.com/gpmgo/gopm
func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		log.Printf("Fetcher error : %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
