// Copyright 2019 Anthony George

// Package main implements a simple gRPC server for the blog service
package main

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net"

	"google.golang.org/grpc"

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
	return nil, errors.New("Post not found")
}

func (bs *blogServer) CreatePost(ctx context.Context, post *pb.Post) (*pb.Post, error) {
	log.WithFields(log.Fields{
		"post": post,
	}).Info("Creating post")

	bs.savedPosts = append(bs.savedPosts, post)
	return post, nil
}

func (bs *blogServer) loadPosts(filePath string) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(data, &bs.savedPosts); err != nil {
		return err
	}
	return nil
}

func newServer(filePath string) *blogServer {
	bs := &blogServer{}
	err := bs.loadPosts(filePath)
	if err != nil {
		log.Error("Failed to load posts: ", err)
		return nil
	}
	return bs
}

func main() {
	lis, err := net.Listen("tcp", "localhost:5334")
	if err != nil {
		log.Error("Error trying to listen: ", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterBlogServer(grpcServer, newServer("posts.json"))
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}
