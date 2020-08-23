package main

import (
	"fmt"

	"go_reset/downloader/testing"
)

type retriever interface {
	Get(string) string
}

func getRetriever() retriever {
	return testing.Retriever{}
}

func main() {
	retriever := getRetriever()
	fmt.Println(retriever.Get("https://www.imooc.com"))
}
