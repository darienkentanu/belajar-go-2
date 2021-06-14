package main

import (
	// sesuaikan dengan struktur folder project masing-masing
	"bytes"
	"fmt"
	"os"
	"project-go-2/protobuf/model"
	"strings"

	"github.com/golang/protobuf/jsonpb"
)

func main() {
	// more code here
	user1 := &model.User{
		Id:       "u001",
		Name:     "Sylvana Windrunner",
		Password: "f0r Th3 H0rD3",
		Gender:   model.UserGender_FEMALE,
	}

	// userList := &model.UserList{List: []*model.User{user1}}

	garage1 := &model.Garage{
		Id:   "g001",
		Name: "Kalimdor",
		Coordinate: &model.GarageCoordinate{
			Latitude:   23.2212847,
			Longtitude: 53.2222212,
		},
	}

	garageList := &model.GarageList{
		List: []*model.Garage{garage1},
	}

	// garageListByUser := &model.GarageListByUser{
	// 	List: map[string]*model.GarageList{
	// 		user1.Id: garageList,
	// 	},
	// }

	// ================ original
	fmt.Printf("# ==== Original\n %v\n", user1)
	// ================ as string
	fmt.Printf(" ==== As String\n %v\n", user1.String())

	// ================ as json string
	buf := bytes.Buffer{}
	marshaller := &jsonpb.Marshaler{}
	err1 := marshaller.Marshal(&buf, garageList)
	if err1 != nil {
		fmt.Println(err1.Error())
		os.Exit(0)
	}
	jsonString := buf.String()
	fmt.Printf("# ==== As JSON String\n %v\n", jsonString)

	buf2 := strings.NewReader(jsonString)
	protoObject := &model.GarageList{}

	unmarshaller := &jsonpb.Unmarshaler{}
	err2 := unmarshaller.Unmarshal(buf2, protoObject)
	if err2 != nil {
		fmt.Println(err2.Error())
		os.Exit(0)
	}
	fmt.Printf("# ==== As String\n%v\n", protoObject.String())

	protoObject = &model.GarageList{}
	err3 := jsonpb.UnmarshalString(jsonString, protoObject)
	if err3 != nil {
		fmt.Println(err3.Error())
		os.Exit(0)
	}
	fmt.Printf("# ==== As String\n %v\n", protoObject.String())
}
