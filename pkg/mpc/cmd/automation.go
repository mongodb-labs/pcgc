package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
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

		exitOnErr(err)

		prettyJSON(automationStatus)
	},
}

// automationStatusCmd represents  status command
var automationRetrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Automation retrieve",

	Run: func(cmd *cobra.Command, args []string) {
		automationStatus, err := newClient().GetAutomationConfig(projectID)

		exitOnErr(err)

		prettyJSON(automationStatus)
	},
}

func aliasNormalizeFunc(_ *pflag.FlagSet, name string) pflag.NormalizedName {
	switch name {
	case "group-id":
		name = "project-id"
	}
	return pflag.NormalizedName(name)
}

func init() {
	automationStatusCmd.Flags().StringVar(&projectID, "project-id", "", "Project ID, group-id can also be used")
	_ = automationStatusCmd.MarkFlagRequired("project-id")
	automationStatusCmd.Flags().SetNormalizeFunc(aliasNormalizeFunc)

	automationRetrieveCmd.Flags().StringVar(&projectID, "project-id", "", "Project ID, group-id can also be used")
	_ = automationRetrieveCmd.MarkFlagRequired("project-id")
	automationRetrieveCmd.Flags().SetNormalizeFunc(aliasNormalizeFunc)

	rootCmd.AddCommand(automationCmd)
	automationCmd.AddCommand(automationStatusCmd)
	automationCmd.AddCommand(automationRetrieveCmd)
}
