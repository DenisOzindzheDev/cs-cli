/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/DenisOzindzheDev/cs-cli/internal/http"
	"github.com/DenisOzindzheDev/cs-cli/internal/kube"
	"github.com/spf13/cobra"
)

var (
	channelNamespace = "channel-namespace"
	vaultNamespace   = "vault-namespace"
	dataPath         = "data-path"
)

// vaultCmd represents the vault command
var vaultCmd = &cobra.Command{
	Use:   "vault",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("vault called")
		fmt.Println(channelNamespace, vaultNamespace, dataPath)

		//extract k8s secrets
		vault := kube.GetSecrets(channelNamespace)
		//call vault
		http.ExtractData(vault.HC_VAULT_ROLE_ID, vault.HC_VAULT_SECRET_ID, "platformeco/data/apim-channel-prod/auth", vaultNamespace)

	},
}

func init() {
	//root nested
	rootCmd.AddCommand(vaultCmd)
	//flags
	vaultCmd.Flags().StringVarP(&channelNamespace, "channel-namespace", "n", "", "channel namespace")
	vaultCmd.Flags().StringVarP(&vaultNamespace, "vault-namespace", "v", "", "vault namespace")
	vaultCmd.Flags().StringVarP(&dataPath, "data-path", "p", "", "vault path")
}
