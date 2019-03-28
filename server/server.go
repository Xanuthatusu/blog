// Copyright 2019 Anthony George

// Package main implements a simple gRPC server for the blog service
package main

import (
	"context"
	"encoding/json"
	"errors"
	"google.golang.org/grpc"
	"io/ioutil"
	"net"

	log "github.com/sirupsen/logrus"
	pb "github.com/xanuthatusu/blog/protos"
)

type blogServer struct {
	// TODO replace with an actual database
	savedPosts []*pb.Post // read-only after initialized
}

func (bs *blogServer) GetPost(ctx context.Context, post *pb.Post) (*pb.Post, error) {
	return bs.getPostByID(post.Id)
}

func (bs *blogServer) getPostByID(id int32) (*pb.Post, error) {
	for _, post := range bs.savedPosts {
		if post.Id == id {
			return post, nil
		}
	}
	log.Warn("Post not found")
	return nil, errors.New("Post not found")
}

func (bs *blogServer) loadPosts(filePath string) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Error("Failed to load posts:", err)
		return
	}
	if err = json.Unmarshal(data, &bs.savedPosts); err != nil {
		log.Error("Failed to load posts:", err)
		return
	}
}

func newServer() *blogServer {
	bs := &blogServer{}
	bs.loadPosts("posts.json")
	return bs
}

func main() {
	lis, err := net.Listen("tcp", "localhost:5334")
	if err != nil {
		log.Error("Error trying to listen: ", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterBlogServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
