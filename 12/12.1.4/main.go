package main

import(
	"fmt"

	"os"
)

func main() {
	fmt.Printf("ユーザーID: %d\n", os.Getuid())
	fmt.Printf("ユーザのグループID: %d\n", os.Getgid())
	groups, _ := os.Getgroups()
	fmt.Printf("サブグループID: %v\n", groups)
}
