package main

import (
	"archive/zip"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {

	http.HandleFunc("/", handeler)
	http.ListenAndServe(":8080", nil)
}

func handeler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=ascii_sample.zip")

	makezipfile("sample.zip")

}

// zipファイルを作成する
func makezipfile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()

	w, err := zipWriter.Create("a.text")
	if err != nil {
		panic(err)
	}
	io.Copy(w, strings.NewReader("zip web download"))

}
