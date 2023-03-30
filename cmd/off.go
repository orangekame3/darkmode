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

	switch desktopEnv {
	case "windows":
		cmd := exec.Command("powershell.exe", "-Command", "New-ItemProperty -Path HKCU:\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Themes\\Personalize -Name SystemUsesLightTheme -Value 1 -Type Dword -Force; New-ItemProperty -Path HKCU:\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Themes\\Personalize -Name AppsUseLightTheme -Value 1 -Type Dword -Force")
		if err := cmd.Run(); err != nil {
			return err
		}
	case "gnome":
		cmd := exec.Command("gsettings", "set", "org.gnome.desktop.interface", "gtk-theme", "Adwaita")
		return cmd.Run()
	case "kde":
		// Set ColorScheme to Breeze Dark in kdeglobals
		cmd := exec.Command("kwriteconfig5", "--file", "~/.config/kdeglobals", "--group", "General", "--key", "ColorScheme", "Breeze")
		if err := cmd.Run(); err != nil {
			return err
		}

		// Reconfigure KWin
		cmd = exec.Command("qdbus", "org.kde.KWin", "/KWin", "reconfigure")
		return cmd.Run()
	case "xfce":
		cmd := exec.Command("xfconf-query", "-c", "xsettings", "-p", "/Net/ThemeName", "-s", "Adwaita")
		if err := cmd.Run(); err != nil {
			return err
		}
	default:
		return errors.New("unsupported desktop environment: " + desktopEnv)
	}
	return errors.New("unsupported desktop environment: " + desktopEnv)
}

func init() {
	rootCmd.AddCommand(offCmd)
}
