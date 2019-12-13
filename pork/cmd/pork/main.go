package main

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ssuareza/go-building-devops-tools/pork"
)

var rootCmd *cobra.Command

func main() {
	rootCmd.Execute()
}

func init() {
	rootCmd = &cobra.Command{
		Use:   "pork",
		Short: "Project Forking Tool for Github",
	}
	rootCmd.AddCommand(pork.SearchCmd)
	rootCmd.AddCommand(pork.DocsCmd)
	rootCmd.AddCommand(pork.CloneCmd)
	rootCmd.AddCommand(pork.ForkCmd)
	viper.SetDefault("location", os.Getenv("HOME"))
	viper.SetConfigName("pork")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
}
