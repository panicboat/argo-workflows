package clustertemplate

import (
	"github.com/spf13/cobra"

	"github.com/argoproj/argo-workflows/v3/cmd/argo/commands/client"
	"github.com/argoproj/argo-workflows/v3/cmd/argo/commands/common"
	clusterworkflowtmplpkg "github.com/argoproj/argo-workflows/v3/pkg/apiclient/clusterworkflowtemplate"
)

func NewGetCommand() *cobra.Command {
	var output = common.NewPrintWorkflowOutputValue("")

	command := &cobra.Command{
		Use:   "get CLUSTER WORKFLOW_TEMPLATE...",
		Short: "display details about a cluster workflow template",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, apiClient, err := client.NewAPIClient(cmd.Context())
			if err != nil {
				return err
			}
			serviceClient, err := apiClient.NewClusterWorkflowTemplateServiceClient()
			if err != nil {
				return err
			}
			for _, name := range args {
				wftmpl, err := serviceClient.GetClusterWorkflowTemplate(ctx, &clusterworkflowtmplpkg.ClusterWorkflowTemplateGetRequest{
					Name: name,
				})
				if err != nil {
					return err
				}
				printClusterWorkflowTemplate(wftmpl, output.String())
			}
			return nil
		},
	}

	command.Flags().VarP(&output, "output", "o", "Output format. "+output.Usage())
	return command
}
