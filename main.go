package main

import (
	complexpb "baquiax.me/protobufers-go/src/complex"
	enumpb "baquiax.me/protobufers-go/src/enum_example"
	simplepb "baquiax.me/protobufers-go/src/simple"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
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

func toJSON(pb proto.Message) string {
	marshallers := jsonpb.Marshaler{}
	out, err := marshallers.MarshalToString(pb)
	if err != nil {
		log.Fatalln("Can't convert JSON", err)
		return ""
	}

	return out
}

func fromJSON(in string, pb proto.Message) {
	err := jsonpb.UnmarshalString(in, pb)
	if err != nil {
		log.Fatalln("Can't unmarshall string to poto.Message")
	}
}

func jsonDemo(pb proto.Message) {
	smAsString := toJSON(pb)
	fmt.Println(smAsString)

	sm2 := &simplepb.SimpleMessage{}
	fromJSON(smAsString, sm2)
	fmt.Println("Sucess created", sm2)
}

func doEnum() {
	ep := enumpb.EnumMessage{
		Id:           42,
		DayOfTheWeek: enumpb.DayOfTheWeek_THURSDAY,
	}

	ep.DayOfTheWeek = enumpb.DayOfTheWeek_MONDAY
	fmt.Println(ep)
}

func doComplex() {
	cm := complexpb.ComplexMessage{
		OneDummy: &complexpb.DummyMessage{
			Id:   1,
			Name: "Dummy message",
		},
		MultipleDummy: []*complexpb.DummyMessage{
			{
				Id:   1,
				Name: "Dummy message",
			},
			{
				Id:   2,
				Name: "Second message",
			},
			{
				Id:   3,
				Name: "Third message",
			},
		},
	}

	fmt.Println(cm)
}

func main() {
	sm := doSimple()
	readAndWriteDemo(sm)
	jsonDemo(sm)
	doEnum()
	doComplex()
}
