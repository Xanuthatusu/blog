// Copyright 2019 Anthony George

// Package main implements a client to connect to the basic blog server
package main

import (
	"context"
	"time"

	"google.golang.org/grpc"

	log "github.com/sirupsen/logrus"
	pb "github.com/xanuthatusu/blog/protos"
)

func main() {
	conn, err := grpc.Dial("localhost:5334", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Could not connect to server: ", err)
	}

	defer conn.Close()

	client := pb.NewBlogClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	post := &pb.GetPostReq{
		Id: 1,
	}

	res, err := client.GetPost(ctx, post)
	if err != nil {
		log.Fatal("Error when trying to get a post: ", err)
	}
	log.Info(res)
}
