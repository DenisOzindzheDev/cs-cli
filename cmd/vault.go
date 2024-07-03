/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/DenisOzindzheDev/cs-cli/internal/kube"
	"github.com/DenisOzindzheDev/cs-cli/internal/rest"
	"github.com/spf13/cobra"
)

var (
	channelNamespace = "channel-namespace"
	vaultNamespace   = "vault-namespace"
	dataPath         = "data-path"
)

var vaultCmd = &cobra.Command{
	Use:   "vault",
	Short: "Extract vault data",
	Long:  `Allows to extract data from vault`,
	Run: func(cmd *cobra.Command, args []string) {
		vault := kube.GetSecrets(channelNamespace)
		rest.ExtractData(vault.HC_VAULT_ROLE_ID, vault.HC_VAULT_SECRET_ID, dataPath, vaultNamespace)
	},
}

func init() {

	rootCmd.AddCommand(vaultCmd)

	vaultCmd.Flags().StringVarP(&channelNamespace, "channel-namespace", "n", "", "channel namespace")
	vaultCmd.Flags().StringVarP(&vaultNamespace, "vault-namespace", "v", "", "vault namespace(optional)")
	vaultCmd.Flags().StringVarP(&dataPath, "data-path", "p", "", "vault path")

	vaultCmd.MarkFlagRequired("channel-namespace")
	vaultCmd.MarkFlagRequired("data-path")

}
