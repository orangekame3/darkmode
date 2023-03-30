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

var defaultConfig = `desktop:
  environment: %s
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
		env, _ := cmd.Flags().GetString("env")
		switch env {
		case "windows":
			defaultConfig = fmt.Sprintf(defaultConfig, "windows")
		case "gnome":
			defaultConfig = fmt.Sprintf(defaultConfig, "gnome")
		case "kde":
			defaultConfig = fmt.Sprintf(defaultConfig, "kde")
		case "xfce":
			defaultConfig = fmt.Sprintf(defaultConfig, "xfce")
		default:
			fmt.Println("Invalid environment specified:", env)
			return
		}
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
			os.Remove(configFile)
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
			fmt.Println("Config file updated:", configFile)
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
	initCmd.Flags().StringP("env", "e", "windows", "set your platform (windows/gnome/kde/xfce)")
}
