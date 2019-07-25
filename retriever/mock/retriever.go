package mock

import "fmt"

type Retriever1 struct {
	Contents string
}

func (r Retriever1) String() string {
	return fmt.Sprintf("Retriever:{Contents:=%s}", r.Contents)
}

func (r Retriever1) Post(url string, from map[string]string) string {
	r.Contents = from["Contents"]
	return "ok"
}

func (r Retriever1) Get(url string) string {
	return r.Contents
}
