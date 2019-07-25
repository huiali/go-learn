package parser

import (
	"golang/crawler/engine"
	"regexp"
)

const profileRe = `<td><span class="label">年龄:</span>([\d]+)岁`

func ParseProfile(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(profileRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, match := range matches {
		result.Items = append(result.Items, string(match[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(match[1]),
			ParseFunc: engine.NilParser,
		})
	}
	return result
}
