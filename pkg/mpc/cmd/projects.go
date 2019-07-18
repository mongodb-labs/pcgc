package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// projectsCmd represents the projects command
var projectsCmd = &cobra.Command{
	Use:     "projects",
	Short:   "Projects operations",
	Long:    "Create, List and manage your MongoDB private cloud projects.",
	Aliases: []string{"groups"},
}

// listCmd represents the list command
var listProjectsCmd = &cobra.Command{
	Use:   "list",
	Short: "List projects",

	Run: func(cmd *cobra.Command, args []string) {
		projects, err := newClient().GetAllProjects()

		if err != nil {
			fmt.Println("Error:", err)
		}

		prettyJSON(projects)
	},
}

// createCmd represents the create command
var createProjectCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a project",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			exitOnErr("create needs a name for the project")
		}
		newProject, err := newClient().CreateOneProject(args[0], orgID)

		exitOnErr(err)
		prettyJSON(newProject)
	},
}

func init() {
	createProjectCmd.Flags().StringVar(&orgID, "org-id", "", "Organization ID for the group")
	rootCmd.AddCommand(projectsCmd)
	projectsCmd.AddCommand(listProjectsCmd)
	projectsCmd.AddCommand(createProjectCmd)
}
