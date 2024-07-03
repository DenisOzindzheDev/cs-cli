package kubeclients

import (
	"fmt"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// Create clients from os kubernetes configuration file and returns map of contexts and interfaces
func CreateClients() (map[string]kubernetes.Interface, error) {
	kubeconfigPath := os.Getenv("KUBECONFIG")
	if kubeconfigPath == "" {
		fmt.Printf("Warning: kubeconfig is not set in the environment using the default configuration\n")
		if home := homedir.HomeDir(); home != "" {
			kubeconfigPath = filepath.Join(home, ".kube", "config")
			fmt.Printf("using %s\n", kubeconfigPath)
		} else {
			kubeconfigPath = ""
		}
	}

	loadingRules := &clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfigPath}
	configOverrides := &clientcmd.ConfigOverrides{}

	kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)

	rawConfig, err := kubeconfig.RawConfig()
	if err != nil {
		fmt.Printf("Error: cannot load raw kubeconfig, %v\n", err)
		return nil, fmt.Errorf("Error: cannot load raw kubeconfig, %v\n", err)
	}

	clients := make(map[string]kubernetes.Interface)
	for contextName := range rawConfig.Contexts {
		contextConfig := clientcmd.NewNonInteractiveClientConfig(rawConfig, contextName, configOverrides, loadingRules)
		config, err := contextConfig.ClientConfig()
		if err != nil {
			fmt.Printf("Error: cannot load client config, for %s : %v\n", contextName, err)
			continue
		}

		client, err := kubernetes.NewForConfig(config)
		if err != nil {
			fmt.Printf("Error: cannot create client for context %s : %v", contextName, err)
			continue
		}
		clients[contextName] = client
	}
	return clients, nil
}
