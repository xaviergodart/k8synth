package main

import (
	"context"
	"flag"
	"fmt"
	"k8synth/midi"
	"path/filepath"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	namespace := flag.String("namespace", "", "kube namespace")
	flag.Parse()

	midi, err := midi.New()
	if err != nil {
		panic(err.Error())
	}

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	ctx := context.Background()
	for {
		pods, err := clientset.CoreV1().Pods(*namespace).List(ctx, metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

		// we link each mod to a midi channel and send a midi note
		// in fifth with each others
		note := 60
		for channel, _ := range pods.Items {
			midi.NoteOn(0, uint8(channel), uint8(note), uint8(100))
			note += 5
		}

		time.Sleep(2 * time.Second)
	}
}
