// http://localhost:18888に対してGETメソッドでアクセスし、ボディを取得する
// $ curl http://localhost:18888をGo言語で実装
package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:18888")
	if err != nil { // Go言語のエラー処理コード　例外処理はない
		panic(err)
	}
	defer resp.Body.Close()                // 後処理コードdeferを付与すると、この関数から抜けたあとにこの文を実行する
	body, err := ioutil.ReadAll(resp.Body) //bodyの内容をバイト列として読み込む
	if err != nil {
		panic(err)
	}

	log.Println(string(body))

}
