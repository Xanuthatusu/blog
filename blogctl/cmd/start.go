package cmd

import (
	"fmt"
	"net"

	pb "github.com/xanuthatusu/blog/protos"
	"github.com/xanuthatusu/blog/server"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "starts the blog server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting blog server")

		lis, err := net.Listen("tcp", ":5334")
		if err != nil {
			fmt.Println(err)
			return
		}

		bServer := server.New("posts.json")

		grpcServer := grpc.NewServer()
		pb.RegisterBlogServer(grpcServer, bServer)

		grpcServer.Serve(lis)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
