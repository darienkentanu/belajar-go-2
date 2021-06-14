package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"project-go-2/grpc/common/config"
	"project-go-2/grpc/common/model"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

func serviceGarage() model.GaragesClient {
	port := config.SERVICE_GARAGE_PORT
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("couldn't connect to", port, err)
	}
	return model.NewGaragesClient(conn)
}

func serviceUser() model.UsersClient {
	port := config.SERVICE_USER_PORT
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Couldn't connect to", port, err)
	}
	return model.NewUsersClient(conn)
}

func main() {
	user1 := model.User{
		Id:       "d001",
		Name:     "Darien Kentanu",
		Password: "infinity",
		Gender:   model.UserGender(model.UserGender_value["Male"]),
	}

	garage1 := model.Garage{
		Id:   "e001",
		Name: "invoker",
		Coordinate: &model.GarageCoordinate{
			Latitude:   45.234234234,
			Longtitude: 54.234234444,
		},
	}

	// ....

	user := serviceUser()

	fmt.Println("\n", "==============> user test")

	// register user1
	user.Register(context.Background(), &user1)

	// // register user2
	// user.Register(context.Background(), &user2)

	// show all registered users
	res1, err := user.List(context.Background(), &empty.Empty{})
	if err != nil {
		log.Fatal(err.Error())
	}
	res1String, _ := json.Marshal(res1.List)
	log.Println(string(res1String))

	garage := serviceGarage()
	fmt.Println("\n", "==============> garage test A")

	// add garage1 to user1
	garage.Add(context.Background(), &model.GarageAndUserId{
		UserId: user1.Id,
		Garage: &garage1,
	})

	// show all garages of user1
	res2, err := garage.List(context.Background(), &model.GarageUserId{UserId: user1.Id})
	if err != nil {
		log.Fatal(err.Error())
	}

	res2String, _ := json.Marshal(res2.List)
	log.Println(string(res2String))
}
