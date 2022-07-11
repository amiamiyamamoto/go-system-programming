package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")
	// json化する元のデータ
	source := map[string]string{
		"Hello": "World",
	}

	// json化する
	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	encoder.SetIndent("", "    ")
	encoder.Encode(source)

	// gzipに圧縮する
	gz_q := gzip.NewWriter()

	// 圧縮に平文渡す
	// 標準出力に平文わたす
	// レスポンスライターに圧縮わたす

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
