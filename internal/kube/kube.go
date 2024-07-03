package kube

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/DenisOzindzheDev/cs-cli/internal/models"
	"github.com/DenisOzindzheDev/cs-cli/pkg/kubeclients"
	"github.com/DenisOzindzheDev/cs-cli/pkg/tableprint"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
)

func GetSecrets(namespace string) models.VaultSecret {
	clients, err := kubeclients.CreateClients()
	if err != nil {
		fmt.Printf("Clients error: %v", err)
		return models.VaultSecret{}
	}

	dataMap := make(map[string]string)
	wg := &sync.WaitGroup{}
	for k, v := range clients {
		wg.Add(1)
		go func() {
			defer wg.Done()
			secrets, err := v.CoreV1().Secrets(namespace).List(context.TODO(), metav1.ListOptions{})
			if err != nil {
				fmt.Printf("Cannot list secrets for %s : %v\n", k, err)
			}
			for _, secret := range secrets.Items {
				if secret.Name == "vault-secret" {
					fmt.Printf("\n\nVault secret found in %v\n", k)
					for k, v := range secret.Data {
						value := string(v)
						dataMap[k] = value
						// fmt.Printf("\n%v  %v", k, value)
					}
				}
			}
		}()
	}
	wg.Wait()
	tableprint.TablePrintStringMap(os.Stdout, dataMap, "Key", "Value")
	return models.VaultSecret{HC_VAULT_ROLE_ID: dataMap["HC_VAULT_ROLE_ID"], HC_VAULT_SECRET_ID: dataMap["HC_VAULT_SECRET_ID"]}
}
