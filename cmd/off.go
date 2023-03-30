/*
Copyright © 2023 Takafumi Miyanaga miya.org.0309@gmail.com
*/
package cmd

import (
	"errors"
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// offCmd represents the off command
var offCmd = &cobra.Command{
	Use:   "off",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if err := switchToLightMode(); err != nil {
			fmt.Println("Error switching to light mode:", err)
		} else {
			fmt.Println("Switched to light mode")
		}
	},
}

func switchToLightMode() error {
	desktopEnv := viper.GetString("desktop.environment")
	lightTheme := viper.GetString("desktop.light-theme")

	switch desktopEnv {
	case "windows":
		cmd := exec.Command("powershell.exe", "-Command", "New-ItemProperty -Path HKCU:\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Themes\\Personalize -Name SystemUsesLightTheme -Value 1 -Type Dword -Force; New-ItemProperty -Path HKCU:\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Themes\\Personalize -Name AppsUseLightTheme -Value 1 -Type Dword -Force")
		if err := cmd.Run(); err != nil {
			return err
		}
	case "gnome":
		cmd := exec.Command("gsettings", "set", "org.gnome.desktop.interface", "gtk-theme", lightTheme)
		return cmd.Run()
	default:
		return errors.New("unsupported desktop environment: " + desktopEnv)
	}
	return nil
}

func init() {
	rootCmd.AddCommand(offCmd)
}
