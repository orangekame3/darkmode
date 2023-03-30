/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var defaultConfig = `
desktop:
  environment: windows
#  environment: gnome
#  environment: kde
#  environment: xfce
`

var configFile = filepath.Join(configDir(), "darkmode.yaml")

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := os.Stat(configFile); os.IsNotExist(err) {
			os.MkdirAll(filepath.Dir(configFile), 0700)
			f, err := os.Create(configFile)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer f.Close()

			if _, err := f.WriteString(defaultConfig); err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Config file created:", configFile)
		} else {
			fmt.Println("Config file already exists:", configFile)
		}
	},
}

func configDir() string {
	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	return filepath.Join(home, ".config", "darkmode")
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
