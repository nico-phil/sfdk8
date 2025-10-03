package main

import "fmt"

func main() {
	err := run()
	if err != nil {
		fmt.Println(err)
	}
}

func run() error {
	fmt.Println("auth service running")
	return nil
}
