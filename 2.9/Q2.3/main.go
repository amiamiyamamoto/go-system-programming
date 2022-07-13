package main

import (
	"bytes"
	"encoding/json"
	"io"
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

	// Clientへ返す
	// io.WriteString(w, buffer.String())
	io.WriteString(w, "aaa")

	// 標準出力する →　成功
	// io.Copy(os.Stdout, &buffer)

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
