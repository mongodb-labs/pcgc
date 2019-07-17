package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/mongodb-labs/pcgc/pkg/httpclient"
	"github.com/mongodb-labs/pcgc/pkg/opsmanager"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// projectsCmd represents the projects command
var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "Projects operations",
	Long:  "Create, List and manage your mongo private cloud projects.",
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List projects",

	Run: func(cmd *cobra.Command, args []string) {
		resolver := httpclient.NewURLResolverWithPrefix(viper.GetString("baseURL"), "/api/public/v1.0")
		client := opsmanager.NewClientWithAuthentication(resolver, viper.GetString("username"), viper.GetString("password"))
		projects, err := client.GetAllProjects()

		if err != nil {
			fmt.Println("Error:", err)
		}
		json, err := json.MarshalIndent(projects, "", "\t")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(json))
	},
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a project",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Coming soon")
	},
}

func init() {
	rootCmd.AddCommand(projectsCmd)
	projectsCmd.AddCommand(listCmd)
	projectsCmd.AddCommand(createCmd)
}
