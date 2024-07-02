package kube

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/DenisOzindzheDev/cs-cli/internal/models"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func GetSecrets(namespace string) models.Vault {
	// parse k8s
	// HOME ./kube config
	kubeconfig := ""
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
	//load k8s
	// todo implement WG
	for contextName := range config.Contexts {
		clientConfig := clientcmd.NewNonInteractiveClientConfig(*config, contextName, &clientcmd.ConfigOverrides{}, nil)
		restConfig, err := clientConfig.ClientConfig()
		if err != nil {
			fmt.Printf("Cannot load client config for %s context", contextName)
		}

		clientSet, err := kubernetes.NewForConfig(restConfig)
		if err != nil {

		}
		//find namespace in clientSet and if matches get vault-secret and extract it
		secrets, err := clientSet.CoreV1().Secrets(namespace).List(context.TODO(), metav1.ListOptions{})
		if err != nil {

		}
		// fmt.Println(secrets.Items) debug
		for _, secret := range secrets.Items {
			if secret.Name == "vault-secret" {
				// fmt.Println(secret.Data)
				for k, v := range secret.Data {
					//byte to string
					value := string(v)
					// log.Print(value)
					// decodeValue, err := base64.StdEncoding.DecodeString(value)
					// if err != nil {
					// 	fmt.Println("Something went wrong decoding")
					// }
					dataMap[k] = value
					fmt.Printf("Vault secret %v  %v\n", k, value)
				}
			}
		}

	}

	// in each context find namespace and map into struct

	return models.Vault{HC_VAULT_ROLE_ID: dataMap["HC_VAULT_ROLE_ID"], HC_VAULT_SECRET_ID: dataMap["HC_VAULT_SECRET_ID"]}
}
