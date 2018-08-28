package main

import (
	"learngo/retriever/mock"
	"fmt"
	"learngo/retriever/real"
	"time"
)

/**
    @author good mood
    2018/7/15 下午2:57
*/

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

const url  = "http://www.imooc.com"

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func post(poster Poster)  {
	poster.Post("http://www.imooc.com",
		map[string]string{
			"name": "cao",
			"course": "golang",
		})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})
	return s.Get(url)
}
func main() {
	var r Retriever
	retriver := &mock.Retriever{"this is a fake imooc.com"}

	r = retriver
	inspect(r)

	r = &real.Retriever{UserAgent:"Mozilla/5.0", TimeOut:time.Minute}

	inspect(r)

	// Type assertion
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}

	fmt.Println("Try a session")
	fmt.Println(session(retriver))

	//fmt.Println(download(r))
}

func inspect(r Retriever)  {
	fmt.Println("Inspecting", r)
	fmt.Printf(" > %T %v\n", r, r)
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}