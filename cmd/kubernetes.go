package cmd

import (
	"fmt"
	"os"

	"github.com/civo/cli/utility"
	"github.com/spf13/cobra"
)

var kubernetesCmd = &cobra.Command{
	Use:     "kubernetes",
	Aliases: []string{"k3s", "k8s", "kube"},
	Short:   "Details of Civo Kubernetes clusters",
}

var kubernetesApplicationsCmd = &cobra.Command{
	Use:     "applications",
	Aliases: []string{"application", "app", "apps", "app", "application", "addon", "addons", "marketplace", "k8s-apps", "k8s-app", "k3s-apps", "k3s-app"},
	Short:   "Details of Civo Kubernetes applications",
}

func init() {
	rootCmd.AddCommand(kubernetesCmd)
	kubernetesCmd.AddCommand(kubernetesListCmd)
	kubernetesCmd.AddCommand(kubernetesListVersionCmd)
	kubernetesCmd.AddCommand(kubernetesShowCmd)
	kubernetesCmd.AddCommand(kubernetesConfigCmd)
	kubernetesCmd.AddCommand(kubernetesCreateCmd)
	kubernetesCmd.AddCommand(kubernetesRenameCmd)
	kubernetesCmd.AddCommand(kubernetesUpgradeCmd)
	kubernetesCmd.AddCommand(kubernetesScaleCmd)
	kubernetesCmd.AddCommand(kubernetesRemoveCmd)
	kubernetesCmd.AddCommand(kubernetesRecycleCmd)

	home, err := os.UserHomeDir()
	if err != nil {
		utility.Error("%s", err)
		os.Exit(1)
	}

	kubernetesConfigCmd.Flags().BoolVarP(&saveConfig, "save", "s", false, "save the config")
	kubernetesConfigCmd.Flags().BoolVarP(&switchConfig, "switch", "i", false, "switch context to newly-created cluster")
	kubernetesConfigCmd.Flags().BoolVarP(&mergeConfig, "merge", "m", false, "merge the config with existing kubeconfig if it already exists.")
	kubernetesConfigCmd.Flags().StringVarP(&localPathConfig, "local-path", "p", fmt.Sprintf("%s/.kube/config", home), "local path to save the kubeconfig file")

	kubernetesCreateCmd.Flags().StringVarP(&targetNodesSize, "size", "s", "g2.medium", "the size of nodes to create.")
	kubernetesCreateCmd.Flags().IntVarP(&numTargetNodes, "nodes", "n", 3, "the number of nodes to create (the master also acts as a node).")
	kubernetesCreateCmd.Flags().StringVarP(&kubernetesVersion, "version", "v", "latest", "the k3s version to use on the cluster. Defaults to the latest.")
	kubernetesCreateCmd.Flags().StringVarP(&applications, "applications", "a", "", "optional, use names shown by running 'civo kubernetes applications ls'")
	kubernetesCreateCmd.Flags().StringVarP(&removeapplications, "remove-applications", "r", "", "optional, remove default application names shown by running  'civo kubernetes applications ls'")
	kubernetesCreateCmd.Flags().BoolVarP(&waitKubernetes, "wait", "w", false, "a simple flag (e.g. --wait) that will cause the CLI to spin and wait for the cluster to be ACTIVE")
	kubernetesCreateCmd.Flags().BoolVarP(&saveConfigKubernetes, "save", "", false, "save the config")
	kubernetesCreateCmd.Flags().BoolVarP(&switchConfigKubernetes, "switch", "", false, "switch context to newly-created cluster")

	kubernetesRenameCmd.Flags().StringVarP(&kubernetesNewName, "name", "n", "", "the new name for the cluster.")

	kubernetesUpgradeCmd.Flags().StringVarP(&kubernetesNewVersion, "version", "v", "", "change the version of the cluster.")
	kubernetesUpgradeCmd.MarkFlagRequired("version")

	kubernetesScaleCmd.Flags().IntVarP(&kubernetesNewNodes, "nodes", "n", 3, "change the total nodes of the cluster.")
	kubernetesScaleCmd.Flags().BoolVarP(&waitKubernetesNodes, "wait", "w", false, "a simple flag (e.g. --wait) that will cause the CLI to spin and wait for the cluster to be ACTIVE")

	kubernetesRecycleCmd.Flags().StringVarP(&kubernetesNode, "node", "n", "", "the node that needs to be recycled.")
	kubernetesRecycleCmd.MarkFlagRequired("node")

	// Kubernetes Applications
	kubernetesCmd.AddCommand(kubernetesApplicationsCmd)
	kubernetesApplicationsCmd.AddCommand(kubernetesAppListCmd)
	// TODO: show command
	kubernetesApplicationsCmd.AddCommand(kubernetesAppAddCmd)

	kubernetesAppAddCmd.Flags().StringVarP(&kubernetesClusterApp, "cluster", "c", "", "the name of the cluster to install the app.")
}
