package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/mongodb-labs/pcgc/pkg/httpclient"
	"github.com/mongodb-labs/pcgc/pkg/opsmanager"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// projectsCmd represents the projects command
var projectsCmd = &cobra.Command{
	Use:     "projects",
	Short:   "Projects operations",
	Long:    "Create, List and manage your mongo private cloud projects.",
	Aliases: []string{"groups"},
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List projects",

	Run: func(cmd *cobra.Command, args []string) {
		withResolver := opsmanager.WithResolver(httpclient.NewURLResolverWithPrefix("http://OPS-MANAGER-INSTANCE", opsmanager.PublicAPIPrefix))
		withDigestAuth := httpclient.WithDigestAuthentication(viper.GetString("username"), viper.GetString("password"))
		withHTTPClient := opsmanager.WithHTTPClient(httpclient.NewClient(withDigestAuth))
		client := opsmanager.NewClient(withResolver, withHTTPClient)
		projects, err := client.GetAllProjects()

		if err != nil {
			fmt.Println("Error:", err)
		}
		json, err := json.MarshalIndent(projects, "", "\t")
		if err != nil {
			er(err)
		}

		fmt.Println(string(json))
	},
}

// createCmd represents the create command
var (
	orgID string

	createCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a project",

		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("len(args)", len(args))
			fmt.Println("args", args)
			fmt.Println("orgID", orgID)
			if len(args) < 1 {
				er("create needs a name for the project")
			}
			//resolver := httpclient.NewURLResolverWithPrefix(viper.GetString("baseURL"), opsmanager.PublicAPIPrefix)
			//client := opsmanager.NewClientWithAuthentication(resolver, viper.GetString("username"), viper.GetString("password"))
			//
			//var createRequest opsmanager.Project
			//if orgID != "" {
			//	createRequest = opsmanager.Project{
			//		Name:  args[0],
			//		OrgID: orgID,
			//	}
			//} else {
			//	createRequest = opsmanager.Project{
			//		Name: args[0],
			//	}
			//}
			//
			//fmt.Printf("%+v\n", createRequest)
			//newProject, err := client.CreateOneProject(createRequest)
			//
			//if err != nil {
			//	useful.PanicOnUnrecoverableError(err)
			//}
			//
			//json, err2 := json.MarshalIndent(newProject, "", "\t")
			//if err2 != nil {
			//	useful.PanicOnUnrecoverableError(err2)
			//}
			//
			//fmt.Println(string(json))
		},
	}
)

func init() {
	createCmd.Flags().StringVar(&orgID, "orgID", "", "Organization ID for the group")
	rootCmd.AddCommand(projectsCmd)
	projectsCmd.AddCommand(listCmd)
	projectsCmd.AddCommand(createCmd)
}
