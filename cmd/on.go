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

	switch desktopEnv {
	case "windows":
		cmd := exec.Command("powershell.exe", "-Command", "New-ItemProperty -Path HKCU:\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Themes\\Personalize -Name SystemUsesLightTheme -Value 0 -Type Dword -Force; New-ItemProperty -Path HKCU:\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Themes\\Personalize -Name AppsUseLightTheme -Value 0 -Type Dword -Force")
		if err := cmd.Run(); err != nil {
			return err
		}
	case "gnome":
		cmd := exec.Command("gsettings", "set", "org.gnome.desktop.interface", "gtk-theme", "Adwaita-dark")
		return cmd.Run()
	case "kde":
		// Set ColorScheme to Breeze Dark in kdeglobals
		cmd := exec.Command("kwriteconfig5", "--file", "~/.config/kdeglobals", "--group", "General", "--key", "ColorScheme", "Breeze Dark")
		if err := cmd.Run(); err != nil {
			return err
		}
		// Reconfigure KWin
		cmd = exec.Command("qdbus", "org.kde.KWin", "/KWin", "reconfigure")
		return cmd.Run()
	case "xfce":
		cmd := exec.Command("xfconf-query", "-c", "xsettings", "-p", "/Net/ThemeName", "-s", "Adwaita-dark")
		if err := cmd.Run(); err != nil {
			return err
		}
	default:
		return errors.New("unsupported desktop environment: " + desktopEnv)
	}
	return nil
}

func init() {
	rootCmd.AddCommand(onCmd)
}
