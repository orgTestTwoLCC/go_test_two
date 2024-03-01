package main

import (
	"fmt"

	"github.com/orgTestOneLCC/go_test_two/main_two"
)

func main() {
	myGreeting := main_two.SayHello2{}
	myGreeting.SayHello()
	fmt.Println("test v0.03")
}
