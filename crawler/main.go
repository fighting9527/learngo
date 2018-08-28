package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"io"
	"golang.org/x/text/encoding"
	"bufio"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/net/html/charset"
)

/**
    @author good mood
    2018/7/27 下午10:24
*/

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
	}

	e := determineEncoding(resp.Body)

	utf8Reader := transform.NewReader(resp.Body,
		simplifiedchinese.GBK.NewDecoder())

	all, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", all)
}

func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}

	charset.DetermineEncoding()
}