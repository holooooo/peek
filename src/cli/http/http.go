package main

import (
	"context"
	v1 "peek/src/gen/peek/v1"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	c := v1.NewManagerServiceClient(conn)
	agent, err := c.CreateAgent(context.TODO(), &v1.Target{
		Namespace: "sm-verify",
		Pod:       "sm-verify-page--131898-7d568466c6-bbn4q",
	})
	if err != nil {
		panic(err)
	}
	logrus.Infof("%v", agent)
}