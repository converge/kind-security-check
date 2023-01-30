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

	err = k8sClient.CheckPodsInDefaultNamespace()
	if err != nil {
		log.Error().Err(err).Msg("")
	}

	err = k8sClient.CheckExposeControlPlane()
	if err != nil {
		log.Error().Err(err).Msg("")
	}

	if err == nil {
		log.Info().Msg("Kubernetes cluster is secure")
	}

}
