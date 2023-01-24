package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"net/http"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	podsGauge = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "pods_total",
		Help: "Number of pods in the cluster",
	})
)

func main() {


	fmt.Println("# Golang Snippets")
	fmt.Println("## Prometheus")

	kubeconfig := os.Getenv("KUBECONFIG")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// Start a goroutine to continuously update the pods gauge
	go func() {
		for {
			pods, err := clientset.CoreV1().Pods("").List(context.TODO(), v1.ListOptions{})
			if err != nil {
				fmt.Println(err)
				continue
			}

			podsGauge.Set(float64(len(pods.Items)))
			time.Sleep(1 * time.Second)
		}
	}()

	// Start a HTTP server to expose the Prometheus metrics
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8386", nil)
}
