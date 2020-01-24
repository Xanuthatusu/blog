package server_test

import (
	"context"
	"fmt"
	"io"
	"net"
	"os"
	"testing"

	pb "github.com/xanuthatusu/blog/protos"
	"github.com/xanuthatusu/blog/server"
	"google.golang.org/grpc"
)

var (
	lis net.Listener
	ctx = context.Background()
)

func startServer() (net.Listener, error) {
	var err error
	lis, err = net.Listen("tcp", ":5334")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	bServer := server.New("../posts.json")

	grpcServer := grpc.NewServer()

	pb.RegisterBlogServer(grpcServer, bServer)

	go func() {
		grpcServer.Serve(lis)
	}()

	return lis, nil
}

func TestMain(m *testing.M) {
	var err error

	lis, err = startServer()
	if err != nil {
		fmt.Println(err)
		return
	}

	exitCode := m.Run()

	err = lis.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	os.Exit(exitCode)
}

func dialServer() (*grpc.ClientConn, pb.BlogClient, error) {
	conn, err := grpc.Dial(":5334", grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}

	return conn, pb.NewBlogClient(conn), nil
}

func TestCreatePost(t *testing.T) {
	conn, bClient, err := dialServer()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

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

	resp, err := bClient.CreatePost(context.Background(), post)
	if err != nil {
		t.Errorf("Error in CreatePost: %v", err)
	}
	if resp.Title != post.Title {
		t.Errorf("Post.Title was incorrect, got: %s, want: %s", resp.Title, post.Title)
	}
}

func TestListPosts(t *testing.T) {
	conn, bClient, err := dialServer()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	posts, err := bClient.ListPosts(ctx, &pb.ListPostsReq{})
	if err != nil {
		t.Fatal(err)
	}

	for {
		post, err := posts.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			t.Fatal(err)
		}

		fmt.Println(post)
	}
}
