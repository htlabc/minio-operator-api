package clientset

import (
	"context"
	"flag"
	"fmt"
	"github.com/htlabc/minio-operator-api/clientset/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"testing"
)

func Test_clientset(t *testing.T) {
	var kubeconfig = flag.String("kubeconfig", "E:\\工作目录\\kubeconfig\\dev-yn-qujing-1.txt", "Absolute path to the kubeconfig file. Required only when running out of cluster.")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Println(err)
	}

	clientset, err := versioned.NewForConfig(config)
	if err != nil {
		fmt.Println("test")
	}

	fmt.Println(clientset)

	obj, err := clientset.MinioV2().Tenants("minio").Get(context.Background(), "minio-tenant", metav1.GetOptions{})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(obj)

}
