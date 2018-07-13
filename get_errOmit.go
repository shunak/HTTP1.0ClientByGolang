package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, _ := http.Get("http://localhost:18888")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
	log.Println("Status:", resp.Status)         //文字列で200 OK
	log.Println("StatusCode:", resp.StatusCode) //数字で200
	log.Println("Headers:", resp.Header)
}
