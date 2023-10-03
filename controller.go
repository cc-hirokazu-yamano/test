package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// ハンドラー関数を定義する
	handler1 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello-1\n")
	}

	// パスとハンドラー関数を結びつける
	http.HandleFunc("/foo/", handler1)

	// Web サーバーを起動する
	log.Fatal(http.ListenAndServe(":8080", nil))
}