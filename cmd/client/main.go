package main

import (
	"github.com/1920853199/passwd/cmd/client/opts"
	"github.com/spf13/cobra"
)

func main() {
	// 1. 创建root cmd
	rootCmd := NewRootCmd()
	rootCmd.Execute()
}

// Root Cmd
func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "passwd-cli",
		Short: "A service-based password management tool.",
		Long:  `A service-based password management tool.`,
	}
	// 添加子命令
	childCommands := []*cobra.Command{
		opts.NewGetCmd(),   // get
		opts.NewSetCmd(),   // set
		opts.NewDelCmd(),   // del
		opts.NewClearCmd(), // clear
		opts.NewAllCmd(),   // all
		opts.NewSshCmd(),   // ssh
	}
	cmd.AddCommand(childCommands...)

	return cmd
}
