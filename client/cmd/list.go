package cmd

import (
	"fmt"

	pb "app/todo"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		todos, err := app.Client.Search(app.Ctx, &pb.SearchReq{Owner: "", Completed: false})
		if err != nil {
			panic(fmt.Sprintf("List CMD err: %s", err))
		}

		for i, todo := range todos.Todos {
			fmt.Println("%d : %s", i, todo)
		}
	},
	Args: cobra.NoArgs,
}
