package main

func readFile(path string) chan string {
	promise := make(chan string)
	promise <- path
	return promise
}
