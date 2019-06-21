package service

import "fmt"

func Test(s string) string {
	fmt.Println("service called", s)
	return "Test" + s
}
