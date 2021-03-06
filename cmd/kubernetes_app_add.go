package cmd

import (
	"fmt"
	"os"

	"github.com/civo/civogo"
	"github.com/civo/cli/config"
	"github.com/civo/cli/utility"
	"github.com/spf13/cobra"
)

var kubernetesClusterApp string

var kubernetesAppAddCmd = &cobra.Command{
	Use:     "add",
	Example: "civo kubernetes application add NAME:PLAN --cluster CLUSTER_NAME",
	Aliases: []string{"install"},
	Args:    cobra.MinimumNArgs(1),
	Short:   "Add the marketplace application to a Kubernetes cluster",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := config.CivoAPIClient()
		if err != nil {
			utility.Error("Creating the connection to Civo's API failed with %s", err)
			os.Exit(1)
		}

		kubernetesFindCluster, err := client.FindKubernetesCluster(kubernetesClusterApp)
		if err != nil {
			utility.Error("Finding a Kubernetes cluster failed with %s", err)
			os.Exit(1)
		}

		appList, err := client.ListKubernetesMarketplaceApplications()
		if err != nil {
			utility.Error("Listing all Kubernetes cluster Applications failed with %s", err)
			os.Exit(1)
		}

		result := utility.RequestedSplit(appList, args[0])
		configKubernetes := &civogo.KubernetesClusterConfig{
			Applications: result,
		}

		kubeCluster, err := client.UpdateKubernetesCluster(kubernetesFindCluster.ID, configKubernetes)
		if err != nil {
			utility.Error("Installing the application in the Kubernetes cluster failed with %s", err)
			os.Exit(1)
		}

		ow := utility.NewOutputWriterWithMap(map[string]string{"ID": kubeCluster.ID, "Name": kubeCluster.Name})

		switch outputFormat {
		case "json":
			ow.WriteSingleObjectJSON()
		case "custom":
			ow.WriteCustomOutput(outputFields)
		default:
			fmt.Printf("The application was installed in the Kubernetes cluster %s\n", utility.Green(kubeCluster.Name))
		}
	},
}
