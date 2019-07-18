package cmd

import (
	"github.com/spf13/cobra"
)

// projectsCmd represents the automation command
var automationCmd = &cobra.Command{
	Use:   "automation",
	Short: "Automation operations",
	Long:  "Manage projects automation configs.",
}

// automationStatusCmd represents  status command
var automationStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Automation status",

	Run: func(cmd *cobra.Command, args []string) {
		automationStatus, err := newClient().GetAutomationStatus(projectID)

		if err != nil {
			er(err)
		}

		prettyJSON(automationStatus)
	},
}

// automationStatusCmd represents  status command
var automationRetrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Automation retrieve",

	Run: func(cmd *cobra.Command, args []string) {
		automationStatus, err := newClient().GetAutomationConfig(projectID)

		if err != nil {
			er(err)
		}

		prettyJSON(automationStatus)
	},
}

func init() {
	automationStatusCmd.Flags().StringVar(&projectID, "project-id", "", "Organization ID for the group")
	_ = automationStatusCmd.MarkFlagRequired("project-id")
	automationRetrieveCmd.Flags().StringVar(&projectID, "project-id", "", "Organization ID for the group")
	_ = automationRetrieveCmd.MarkFlagRequired("project-id")
	rootCmd.AddCommand(automationCmd)
	automationCmd.AddCommand(automationStatusCmd)
	automationCmd.AddCommand(automationRetrieveCmd)
}
