package exercise1

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

func addPerson(addressBook *AddressBook, person *Person) *AddressBook {

	addressBook.People = append(addressBook.People, person)
	writeDown(addressBook)
	return addressBook
}

func writeDown(addressBook *AddressBook) {
	bytes, err := proto.Marshal(addressBook)
	if err != nil {
		log.Fatalln("Error getting bytes", err)
	}

	ioutil.WriteFile("book.bin", bytes, 0644)
}

func readBookFromDisk() *AddressBook {
	bytes, err := ioutil.ReadFile("book.bin")
	if err != nil {
		fmt.Println("Creating the book", err)
	}

	book := AddressBook{}

	err = proto.Unmarshal(bytes, &book)
	if err != nil {
		log.Fatalln("Error marshalling")
	}

	return &book
}

// DoExercise1 Run example
func DoExercise1() {
	addressBook := readBookFromDisk()
	person := &Person{
		Id:          int32(time.Now().Nanosecond()),
		Name:        "Alex",
		Email:       "alex@baquiax.me",
		LastUpdated: &timestamp.Timestamp{},
		Phones: []*Person_PhoneNumber{
			{
				Number: "50212345678",
				Type:   Person_MOBILE,
			},
		},
	}

	addPerson(addressBook, person)

	fmt.Println("Persons in the book: ", len(addressBook.People))
}
