package main

import "fmt"

func main() {
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
