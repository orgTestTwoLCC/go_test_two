package main_two

import (
	"fmt"

	"github.com/orgTestOneLCC/dep_repo"
)

func (g Greeting) SayHello2() {
	myGreeting := main_one.Greeting{}

	myGreeting.SayHello()
}