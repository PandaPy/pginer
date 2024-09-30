package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "pginer",
		Short: "pginer 是一个简单的命令行工具",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, World3!")
		},
	}

	// 执行命令
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
