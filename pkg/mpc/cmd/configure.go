package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh/terminal"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure the tool",
	Long:  `Set up authentication settings as well as the Base URL of the private cloud deployment.`,
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Enter baseURL: ")
		baseURL, _ := reader.ReadString('\n')

		viper.Set("baseURL", strings.TrimSpace(baseURL))

		fmt.Print("Enter Username: ")
		username, _ := reader.ReadString('\n')

		viper.Set("username", strings.TrimSpace(username))

		fmt.Print("Enter Password: ")
		bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			log.Fatal(err)
		}
		password := string(bytePassword)
		viper.Set("password", strings.TrimSpace(password))

		err = viper.WriteConfig()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("\nDone!")
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
}
