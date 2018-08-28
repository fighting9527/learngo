package real

import (
	"time"
	"net/http"
	"net/http/httputil"
)

/**
    @author good mood
    2018/7/15 下午3:05
*/

type Retriever struct {
	UserAgent string
	TimeOut time.Duration
}

func (r *Retriever) Get(url string) string  {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	result, err := httputil.DumpResponse(resp, true)
	resp.Body.Close()

	if err != nil {
		panic(err)
	}
	return string(result)
}