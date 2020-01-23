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

func TestCreatePost(t *testing.T) {
	bs := newServer("../posts.json")

	post := &pb.Post{
		Id:    2,
		Title: "Unit test post",
		DatePosted: &pb.Date{
			Year:       2019,
			Month:      7,
			Day:        4,
			DateString: "07-04-2019",
		},
	}

	resp, err := bs.CreatePost(context.Background(), post)
	if err != nil {
		t.Errorf("Error in CreatePost: %v", err)
	}
	if resp.Title != post.Title {
		t.Errorf("Post.Title was incorrect, got: %s, want: %s", resp.Title, post.Title)
	}
}
