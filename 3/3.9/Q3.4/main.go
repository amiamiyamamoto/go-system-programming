package main

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// zipをWebサーバからダウンロード
func main() {

	http.HandleFunc("/", handeler)
	http.ListenAndServe(":8080", nil)
}

func handeler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=ascii_sample.zip")

	file := makezipfile("sample.zip")
	// defer file.Close()
	// iti, _ := file.Seek(0, 1)

	//ここから

	// file, err := os.Create("sample.zip")

	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()

	// zipWriter := zip.NewWriter(file)
	// // defer zipWriter.Close()

	// watxt, err := zipWriter.Create("a.text")
	// if err != nil {
	// 	panic(err)
	// }
	// io.Copy(watxt, strings.NewReader("zip web download"))
	// // iti, _ := file.Seek(0, 1)
	// // fmt.Println("makezipfile: ", iti)
	// // iti2, _ := file.Seek(0, 0)
	// zipWriter.Close()

	// file.Seek(0, 0)
	// fmt.Println("makezipfile: ", iti2)

	//ここまで
	// fmt.Println("main:", iti)

	io.Copy(w, file)

}

// zipファイルを作成する
func makezipfile(filename string) *os.File {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	// file.Close()

	zipWriter := zip.NewWriter(file)

	w, err := zipWriter.Create("a.text")
	if err != nil {
		panic(err)
	}
	io.Copy(w, strings.NewReader("zip web download"))
	zipWriter.Close()
	// iti, _ := file.Seek(0, 1)
	// fmt.Println("makezipfile: ", iti)
	iti2, _ := file.Seek(0, 0)
	fmt.Println("makezipfile: ", iti2)
	return file

}
