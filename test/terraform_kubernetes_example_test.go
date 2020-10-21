package test

import (
	"fmt"
	"testing"
	"time"

	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/k8s"
)

func TestKubernetesExample(t *testing.T) {
	t.Parallel()

	kubeResourcePath := "../examples/terraform-kubernetes-example/hello-world-deployment.yml"
	serviceName := "hello-world-service"  // Defined in line 28 of hello-world-deployment.yml

	options := k8s.NewKubectlOptions("", "", "default")

	defer k8s.KubectlDelete(t, options, kubeResourcePath)

	k8s.KubectlApply(t, options, kubeResourcePath)

	k8s.WaitUntilServiceAvailable(t, options, serviceName, 120, 1*time.Second)
	service := k8s.GetService(t, options, serviceName)
	url := fmt.Sprintf("http://%s", k8s.GetServiceEndpoint(t, options, service, 5000))

	http_helper.HttpGetWithRetry(t, url, nil, 200, "Hello world!", 30, 3*time.Second)

}