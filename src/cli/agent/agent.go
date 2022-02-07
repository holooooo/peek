package main

import (
	"flag"
	"path/filepath"
	"peek/src/service/agent"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var (
	kubeconfig *string
)

func init() {
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
}

func getKubeconfigOrDie() *kubernetes.Clientset {
	config, err := rest.InClusterConfig()
	if config == nil {
		config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig)
	}
	if err != nil {
		panic(err.Error())
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return client
}

func main() {
	clientSet := getKubeconfigOrDie()
	// TODO make this configurable
	err := agent.Serve(clientSet, agent.Config{
		ServerAddress: "0.0.0.0:9000",
	})
	panic(err)
}