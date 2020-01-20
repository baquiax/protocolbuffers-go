package main

import (
	simplepb "baquiax.me/protobufers-go/src/simple"
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
)

func doSimple() *simplepb.SimpleMessage {
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

	return &sm
}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialize to bytes", err)
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Can't write the file", err)
		return err
	}

	fmt.Println("Data was written")
	return nil
}

func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)

	if err != nil {
		log.Fatalln("Error reading the file", err)
		return err
	}

	err = proto.Unmarshal(in, pb)
	if err != nil {
		log.Fatalln("Error unmarshalling", err)
		return err
	}

	return nil
}

func readAndWriteDemo(sm proto.Message) {
	writeToFile("simple.bin", sm)
	sm2 := &simplepb.SimpleMessage{}
	readFromFile("simple.bin", sm2)
	fmt.Println(sm2)
}

func main() {
	sm := doSimple()
	readAndWriteDemo(sm)
}
