package kube

import (
	"context"
	"fmt"
	"path/filepath"
	"sync"

	"github.com/DenisOzindzheDev/cs-cli/internal/models"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func GetSecrets(namespace string) models.Vault {

	kubeconfig := ""
	wg := &sync.WaitGroup{}
	dataMap := make(map[string]string)

	if home := homedir.HomeDir(); home != "" {
		kubeconfig = filepath.Join(home, ".kube", "config") // ~/.kubeconfig
	} else {
		kubeconfig = ""
	}
	config, err := clientcmd.LoadFromFile(kubeconfig)
	if err != nil {
		fmt.Println("Cannot load kubeconfig")
	}

	// todo implement WG
	for contextName := range config.Contexts {
		wg.Add(1)
		go func() {
			defer wg.Done()
			clientConfig := clientcmd.NewNonInteractiveClientConfig(*config, contextName, &clientcmd.ConfigOverrides{}, nil)
			restConfig, err := clientConfig.ClientConfig()
			if err != nil {
				fmt.Printf("Cannot load client config for %s context\n", contextName)
			}

			clientSet, err := kubernetes.NewForConfig(restConfig)
			if err != nil {
				fmt.Printf("Cannot load client set for %s context\n", contextName)
			}
			secrets, err := clientSet.CoreV1().Secrets(namespace).List(context.TODO(), metav1.ListOptions{})
			if err != nil {
				fmt.Printf("Cannot list secrets for %s context", contextName)
			}
			for _, secret := range secrets.Items {
				if secret.Name == "vault-secret" {
					fmt.Printf("\n\nVault secret found in %v", contextName)
					for k, v := range secret.Data {
						value := string(v)
						dataMap[k] = value
						fmt.Printf("\n%v  %v", k, value)
					}
				}
			}
		}()
	}
	wg.Wait()

	return models.Vault{HC_VAULT_ROLE_ID: dataMap["HC_VAULT_ROLE_ID"], HC_VAULT_SECRET_ID: dataMap["HC_VAULT_SECRET_ID"]}
}
