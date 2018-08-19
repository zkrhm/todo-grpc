package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use: "todo-server",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.Flags().StringP("host", "H", "localhost", "the hostname to bind")
	viper.BindPFlag("host", rootCmd.Flags().Lookup("host"))

	rootCmd.Flags().IntP("port", "p", 5555, "the port to bind")
	viper.BindPFlag("port", rootCmd.Flags().Lookup("port"))
}

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
