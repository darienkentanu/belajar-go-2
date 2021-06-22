package main

import (
	"context"
	"log"
	"net"

	"project-go-2/grpc/common/config"
	"project-go-2/grpc/common/model"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

var localStorage *model.GarageListByUser

func init() {
	localStorage = &model.GarageListByUser{}
	localStorage.List = make(map[string]*model.GarageList)
}

type GaragesServer struct{}

func (GaragesServer) Add(ctx context.Context, param *model.GarageAndUserId) (*empty.Empty, error) {
	userId := param.UserId
	garage := param.Garage

	if _, ok := localStorage.List[userId]; !ok {
		localStorage.List[userId] = &model.GarageList{}
		localStorage.List[userId].List = make([]*model.Garage, 0)
	}
	localStorage.List[userId].List = append(localStorage.List[userId].List, garage)

	log.Println("Adding garage", garage.String(), "for user", userId)

	return &empty.Empty{}, nil
}

func (GaragesServer) List(ctx context.Context, param *model.GarageUserId) (*model.GarageList, error) {
	userId := param.UserId
	return localStorage.List[userId], nil
}

func main() {
	srv := grpc.NewServer()
	var garageSrv GaragesServer
	model.RegisterGaragesServer(srv, garageSrv)

	log.Println("Starting RPC server at", config.SERVICE_GARAGE_PORT)

	l, err := net.Listen("tcp", config.SERVICE_GARAGE_PORT)
	if err != nil {
		log.Fatalf("couldn't listen to %s: %v", config.SERVICE_GARAGE_PORT, err)
	}

	log.Fatal(srv.Serve(l))
}