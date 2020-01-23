// Copyright 2019 Anthony George

// Package main implements a simple gRPC server for the blog service
package server

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	pb "github.com/xanuthatusu/blog/protos"
)

type BlogServer struct {
	// TODO replace with an actual database
	savedPosts []*pb.Post // read-only after initialized
}

func New(filePath string) *BlogServer {
	bs := &BlogServer{}
	err := bs.loadPosts(filePath)
	if err != nil {
		log.Error("Failed to load posts: ", err)
		return nil
	}
	return bs
}

func (bs *BlogServer) GetPost(ctx context.Context, req *pb.GetPostReq) (*pb.Post, error) {
	return bs.getPostByID(req.Id)
}

func (bs *BlogServer) getPostByID(id int32) (*pb.Post, error) {
	for _, post := range bs.savedPosts {
		if post.Id == id {
			return post, nil
		}
	}
	return nil, errors.New("Post not found")
}

func (bs *BlogServer) ListPosts(req *pb.ListPostsReq, stream pb.Blog_ListPostsServer) error {
	for _, post := range bs.savedPosts {
		err := stream.Send(post)
		if err != nil {
			return err
		}
	}
	return nil
}

func (bs *BlogServer) CreatePost(ctx context.Context, post *pb.Post) (*pb.Post, error) {
	log.WithFields(log.Fields{
		"post": post,
	}).Info("Creating post")

	bs.savedPosts = append(bs.savedPosts, post)
	return post, nil
}

func (bs *BlogServer) loadPosts(filePath string) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(data, &bs.savedPosts); err != nil {
		return err
	}
	return nil
}
