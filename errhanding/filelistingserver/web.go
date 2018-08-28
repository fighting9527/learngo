package main

import (
	"net/http"
	"learngo/errhanding/filelistingserver/filelisting"
	"os"
	"github.com/astaxie/beego/logs"
	"log"
	_ "net/http/pprof"
)

/**
    @author good mood
    2018/7/15 下午6:16
*/

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {

		defer func() {
			if  r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				http.Error(
					writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError,
				)
			}

		}()

		err := handler(writer, request)
		if err != nil {
			logs.Warn("Error handling request")
			if userErr, ok := err.(userError); ok {
				http.Error(
					writer,
					userErr.Message(),
					http.StatusBadRequest,
				)
				return
			}
			code := http.StatusOK
			switch  {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(
				writer,
				http.StatusText(code),
				code,
			)
		}
	}
}

type userError interface {
	error
	Message() string
}


func main() {
	http.HandleFunc("/",
		errWrapper(filelisting.HandleFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
