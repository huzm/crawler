package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	url := "https://www.thepaper.cn/"
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("fetch url error:%v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error Status code:%v", resp.StatusCode)
		return
	}
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("read content failed:%v", err)
		return
	}
	//fmt.Println("body:", string(body))
	numlinks := strings.Count(string(body), "<a")
	fmt.Printf("homepage has %d links.\n", numlinks)
	exist := strings.Contains(string(body), "疫情")
	fmt.Printf("是否存在疫情：%v\n", exist)

}
