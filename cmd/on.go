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

// onCmd represents the on command
var onCmd = &cobra.Command{
	Use:   "on",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if err := switchToDarkMode(); err != nil {
			fmt.Println("Error switching to dark mode:", err)
		} else {
			fmt.Println("Switched to dark mode")
		}
	},
}

func switchToDarkMode() error {
	desktopEnv := viper.GetString("desktop.environment")
	onTheme := viper.GetString("desktop.on-theme")

	switch desktopEnv {
	case "windows":
		cmd := exec.Command("powershell.exe", "-Command", "New-ItemProperty -Path HKCU:\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Themes\\Personalize -Name SystemUsesLightTheme -Value 0 -Type Dword -Force; New-ItemProperty -Path HKCU:\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Themes\\Personalize -Name AppsUseLightTheme -Value 0 -Type Dword -Force")
		if err := cmd.Run(); err != nil {
			return err
		}
	case "gnome":
		cmd := exec.Command("gsettings", "set", "org.gnome.desktop.interface", "gtk-theme", "Adwaita")
		if onTheme != "" {
			cmd = exec.Command("gsettings", "set", "org.gnome.desktop.interface", "gtk-theme", onTheme)
		}
		return cmd.Run()
	default:
		return errors.New("unsupported desktop environment: " + desktopEnv)
	}
	return nil
}

func init() {
	rootCmd.AddCommand(onCmd)
}
