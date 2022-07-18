package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
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
	// encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ")
	encoder.Encode(source)

	//gzip化する
	gzip_writer := gzip.NewWriter(w)
	defer gzip_writer.Close()

	// レスポンスライターと標準出力をまとめる
	writer := io.MultiWriter(gzip_writer, os.Stdout)

	// レスポンス返却、標準出力する
	io.Copy(writer, &buffer)

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
