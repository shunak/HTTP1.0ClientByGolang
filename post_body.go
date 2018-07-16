// $ curl -T server.go -H "Content-Type: text/plain" http://localhost:18888 をGoで実装
// (curl コマンドでファイルから読み込んだ任意のコンテンツを送信)
package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	file, err := os.Open("post_body.go")
	// file, err := os.Open("../godev/server.go")
	if err != nil {
		panic(err)
	}
	resp, err := http.Post("http://localhost:18888", "text/plain", file)
	if err != nil {
		panic(err)
	}
	log.Println("Status:", resp.Status)
}
