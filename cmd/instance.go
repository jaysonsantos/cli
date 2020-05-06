package cmd

import (
	"github.com/spf13/cobra"
)

var instanceCmd = &cobra.Command{
	Use:     "instance",
	Aliases: []string{"instances"},
	Short:   "Details of Civo instances",
}

func init() {
	rootCmd.AddCommand(instanceCmd)
	instanceCmd.AddCommand(instanceListCmd)
	instanceCmd.AddCommand(instanceCreateCmd)
	instanceCmd.AddCommand(instanceShowCmd)
	instanceCmd.AddCommand(instanceUpdateCmd)
	instanceCmd.AddCommand(instanceRemoveCmd)
	instanceCmd.AddCommand(instanceRebootCmd)
	instanceCmd.AddCommand(instanceSoftRebootCmd)
	instanceCmd.AddCommand(instanceConsoleCmd)
	instanceCmd.AddCommand(instanceStopCmd)
	instanceCmd.AddCommand(instanceStartCmd)
	instanceCmd.AddCommand(instanceUpgradeCmd)
	instanceCmd.AddCommand(instanceMoveIPCmd)
	instanceCmd.AddCommand(instanceSetFirewallCmd)
	instanceCmd.AddCommand(instancePublicIPCmd)
	instanceCmd.AddCommand(instancePasswordCmd)
	instanceCmd.AddCommand(instanceTagCmd)

	/*
		Flags for instance update cmd
	*/
	instanceUpdateCmd.Flags().StringVarP(&notes, "notes", "n", "", "notes stored against the instance")
	instanceUpdateCmd.Flags().StringVarP(&reverseDNS, "reverse-dns", "r", "", "the reverse DNS entry for the instance")
	instanceUpdateCmd.Flags().StringVarP(&hostname, "hostname", "s", "", "the instance's hostname")

	/*
		Flags for instance create cmd
	*/
	instanceCreateCmd.Flags().BoolVarP(&wait, "wait", "w", false, "wait until the instance's is ready")
	instanceCreateCmd.Flags().StringVarP(&hostnameCreate, "hostname", "s", "", "the instance's hostname")
	instanceCreateCmd.Flags().StringVarP(&size, "size", "i", "", "the instance's size")
	instanceCreateCmd.Flags().StringVarP(&template, "template", "t", "", "the instance's template")
	instanceCreateCmd.Flags().StringVarP(&snapshot, "snapshot", "n", "", "the instance's snapshot")
	instanceCreateCmd.Flags().StringVarP(&publicip, "publicip", "p", "", "the instance's public ip")
	instanceCreateCmd.Flags().StringVarP(&initialuser, "initialuser", "u", "", "the instance's initial user")
	instanceCreateCmd.Flags().StringVarP(&sshkey, "sshkey", "k", "", "the instance's ssh key you can use the Name or the ID")
	instanceCreateCmd.Flags().StringVarP(&network, "network", "r", "", "the instance's network you can use the Name or the ID")
	instanceCreateCmd.Flags().StringVarP(&tags, "tags", "g", "", "the instance's tags")

	/*
		Flags for stop cmd
	*/
	instanceStopCmd.Flags().BoolVarP(&waitStop, "wait", "w", false, "wait until the instance's is stoped")

}
