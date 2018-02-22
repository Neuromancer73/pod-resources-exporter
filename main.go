package main

import (
	"fmt"
	"time"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	//"k8s.io/api/core/v1"
)

func main() {

	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	for {
		pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}

		fmt.Printf("Timestamp: %d\n", time.Now())
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

		namespaces, namespacesErr := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})

		if namespacesErr != nil {
			fmt.Printf("\tError while getting namespaces: %v\n", namespacesErr)
		} else {

			fmt.Printf("Pods list\n")

			for _, namespace := range namespaces.Items {

				fmt.Printf("\tNamespace: %v\n", namespace.Name)

				pods, podsErr := clientset.CoreV1().Pods(namespace.Name).List(metav1.ListOptions{})

				if podsErr != nil {
					fmt.Printf("\t\tError while getting pods from: %v\n", namespacesErr)
				} else {
					fmt.Printf("\t\tPods in namespace: %v\n", namespace.Name)
					for _, pod := range pods.Items {
						fmt.Printf("\t\t%v\n", pod.Name)
					}
				}
			}
		}

		fmt.Printf("----------------------------------------------------")

		if errors.IsNotFound(err) {
			fmt.Printf("Pod not found\n")
		} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
			fmt.Printf("Error getting pod %v\n", statusError.ErrStatus.Message)
		} else if err != nil {
			panic(err.Error())
		} else {
			fmt.Printf("Found pod\n")
		}

		time.Sleep(10 * time.Second)
	}
}
