package main

import (
	"context"
	"testing"

	pb "github.com/xanuthatusu/blog/protos"
)

func TestServer(t *testing.T) {
	bs := newServer("../posts.json")

	resp, err := bs.GetPost(context.Background(), &pb.Post{Id: 1})
	if err != nil {
		t.Errorf("Error in GetPost: %v", err)
	}
	if resp.Title != "test post" {
		t.Errorf("Post.Title was incorrect, got: %s, want: %s", resp.Title, "test post")
	}
}
