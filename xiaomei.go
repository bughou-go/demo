package main

import (
	"github.com/bughou-go/xiaomei/cli/db"
	"github.com/bughou-go/xiaomei/cli/deploy"
	"github.com/bughou-go/xiaomei/cli/develop"
	"github.com/bughou-go/xiaomei/cli/oam"
	"github.com/bughou-go/xiaomei/cli/setup"
	"github.com/bughou-go/xiaomei/config"

	"github.com/spf13/cobra"
)

func main() {
	cobra.EnableCommandSorting = false

	root := &cobra.Command{
		Use:   `xiaomei`,
		Short: `be small and beautiful.`,
	}
	root.AddCommand(develop.Cmds()...)
	root.AddCommand(db.Cmds()...)
	root.AddCommand(oam.Cmds(s)...)
	root.AddCommand(deploy.Cmd(s), setup.Cmd())

	root.Execute()
}
