// package パッケージ宣言

package main

// import文 他のパッケージを参照したい
import (

  // print系のパッケージ
  "fmt"

  // log系のパッケージ
  "log"

  // webサーバ系のパッケージ
  "net/http"
)

func main() {

  //HTTP リクエストを受けてレスポンスを返す系
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    log.Println("received request")
    fmt.Fprintf(w, "Hello Docker!!")
  })

  // サーバのポート系のお話

  log.Println("start server")
  server := &http.Server{Addr: ":8080"}
  if err := server.ListenAndServe(); err != nil {
    log.Println(err)
  }
}
