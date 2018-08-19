package cmd

import (
	"context"
	"fmt"
	"os"

	pb "app/todo"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var rootCmd = &cobra.Command{
	Use: "todo",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
	},
}

type App struct {
	Client pb.TodoServiceClient
	Ctx    context.Context
}

var app *App

func NewApp(host string, port int) *App {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", host, port), grpc.WithInsecure())
	if err != nil {
		panic(fmt.Sprintf("error: %s", err))
	}
	// defer conn.Close()
	client := pb.NewTodoServiceClient(conn)

	return &App{
		Client: client,
		Ctx:    context.Background(),
	}
}

func Execute() {
	_, err := NewConfig()
	if err != nil {

	}
	app = NewApp(viper.Get("host").(string), viper.Get("port").(int))
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
