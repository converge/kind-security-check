package main

import (
	"github.com/converge/kind-security-check/internal/k8s"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

func main() {

	// setup zerolog format
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, NoColor: false})

	// todo: temporary issues mapping
	var issues = map[string]string{}

	// instantiate kubernetes client config
	configDefaultValues := clientcmd.NewDefaultClientConfigLoadingRules().GetDefaultFilename()
	config, err := clientcmd.BuildConfigFromFlags("", configDefaultValues)
	if err != nil {
		log.Error().Err(err).Msg("")
	}

	// instantiate clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Error().Err(err).Msg("")
	}

	k8sClient, err := k8s.NewKubernetesClient(clientset)
	if err = k8sClient.CheckPodsInDefaultNamespace(); err != nil {
		issues["defaultNamespace"] = err.Error()
	}

	if err = k8sClient.CheckExposeControlPlane(); err != nil {
		issues["exposeControlPlane"] = err.Error()
	}

	if len(issues) > 0 {
		for _, issue := range issues {
			log.Error().Msg(issue)
		}
		return
	}

	log.Info().Msg("Kubernetes cluster is secure")

}
