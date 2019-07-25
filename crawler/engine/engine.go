package engine

import (
	"golang/crawler/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, m := range seeds {
		requests = append(requests, m)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("Fetching %s", r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetchter: error fetching url %s:%v", r.Url, err)
			continue
		}
		parseResult := r.ParseFunc(body)
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}

}
