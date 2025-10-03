package main

import "fmt"

var build = "devlop"

func main() {
	err := run()
	if err != nil {
		fmt.Println(err)
	}
}

func run() error {
	fmt.Println("sales service is running", build)
	return nil
}
