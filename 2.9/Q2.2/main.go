package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	records := [][]string{
		{"ami", "yama", "moto"},
		{"anti", "-", "pattern"},
		{"in", "pu", "t"},
	}

	// ファイル出力の場合
	file, err := os.Create("test.csv")
	if err != nil {
		panic(err)
	}
	csvw := csv.NewWriter(file)

	// 標準出力の場合
	// csvw := csv.NewWriter(os.Stdout)

	for _, record := range records {
		if err := csvw.Write(record); err != nil {
			log.Fatalln("err writing record to csv:", err)
		}
	}
	csvw.Flush()
	if err := csvw.Error(); err != nil {
		log.Fatalln(err)
	}
}
