package main

func main() {
	//

	for {
		select {
		case data := <-reader
		case <-exit:
			// ループを抜ける
			break
		}
	}
}
