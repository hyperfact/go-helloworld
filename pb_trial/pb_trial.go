package main

import (
	"fmt"
	"log"

	eg "pb_trial/example"

	pb "github.com/golang/protobuf/proto"
)

func main() {
	var err error

	fmt.Println("")
	pb.String("")

	_ = eg.FOO_X

	t := &eg.Test{
		Label:         pb.String("label"),
		Type:          pb.Int32(11),
		Reps:          []int64{1, 2, 3, 4},
		Optionalgroup: &eg.Test_OptionalGroup{RequiredField: pb.String("group")},
		Union:         &eg.Test_Name{"name"},
	}

	data, err := pb.Marshal(t)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	newt := &eg.Test{}
	err = pb.Unmarshal(data, newt)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	fmt.Println(newt)
	fmt.Println(newt.GetLabel())
	fmt.Println(newt.GetName())
	fmt.Println(newt.GetNumber())
	fmt.Println(newt.GetType())
	fmt.Println(newt.GetReps())
}
