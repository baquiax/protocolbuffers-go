package main

import (
	simplepb "baquiax.me/protobufers-go/src/simple"
	"fmt"
)

func doSimple() {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "Alex",
		SampleList: []int32{1, 4, 7, 8},
	}
	fmt.Println(sm)

	sm.Name = "Josue"

	fmt.Println(sm)

	fmt.Println("The ID is: 	", sm.GetId())
}

func main() {
	doSimple()
}
