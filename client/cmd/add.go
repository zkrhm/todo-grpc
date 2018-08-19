package cmd

import (
	pb "app/todo"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use: "add",
	Run: func(cmd *cobra.Command, args []string) {
		todo := strings.Join(args, " ")
		ack, err := app.Client.Save(app.Ctx, &pb.TodoToSave{Content: todo, Owner: "a user"})
		if err != nil {
			panic(fmt.Sprintf("error %s", err))
		}

		fmt.Println("ack : %s", ack)
	},
}
