package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"net/http"
	"strings"
)

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error Status code:%v", resp.StatusCode)
	}
	bodyReader := bufio.NewReader(resp.Body)
	e := DeterminEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return io.ReadAll(utf8Reader)
}

func DeterminEncoding(r *bufio.Reader) encoding.Encoding {
	bytes,err:= r.Peek(1024)

	if err != nil{
		fmt.Println("fetch error:%v",err)
		return unicode.UTF8
	}
	e,_,_ := charset.DetermineEncoding(bytes,")
	return e
}

func main() {
	url := "https://www.thepaper.cn/"
	//fmt.Println("body:", string(body))
	body, err:= Fetch(url)
	if err !=nil{
		panic(err)
	}
	numlinks := strings.Count(string(body), "<a")
	fmt.Printf("homepage has %d links.\n", numlinks)
	exist := strings.Contains(string(body), "疫情")
	fmt.Printf("是否存在疫情：%v\n", exist)

}
