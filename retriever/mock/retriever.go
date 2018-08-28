package mock

import "fmt"

/**
    @author good mood
    2018/7/15 下午3:01
*/

type Retriever struct {
	Contents string
}

func (r *Retriever) String() string {
	return fmt.Sprintf(
		"Retriever: {Contents=%s}", r.Contents,
	)
}

func (r *Retriever) Get(url string) string  {
	return r.Contents
}

func (r *Retriever) Post(url string, form map[string]string) string  {
	r.Contents = form["contents"]
	return "ok"
}

