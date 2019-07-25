package parser

import (
	"fmt"
	"golang/crawler/fetcher"
	"testing"
)

func TestParseCityList(t *testing.T) {
	body, err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", body)
}
