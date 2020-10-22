package test

import (
	"fmt"
	"os/exec"
	"testing"
	"time"

	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/k8s"
)

func TestKubernetesExample(t *testing.T) {
	t.Parallel()

	kubeResourcePath := "../examples/terraform-kubernetes-example/hello-world-deployment.yml"
	serviceName := "hello-world-service"  // Defined in line 28 of hello-world-deployment.yml
	deploymentName := "hello-world-deployment"  // Defined in line 6 of hello-world-deployment.yml

	options := k8s.NewKubectlOptions("", "", "stratus")

	t.Cleanup(func() {
		t.Log("Cleaning up deployment")
		k8s.KubectlDelete(t, options, kubeResourcePath)
	})

	k8s.KubectlApply(t, options, kubeResourcePath)

	k8s.WaitUntilServiceAvailable(t, options, serviceName, 120, 1*time.Second)

	// Port forward
	portForwardCommand := exec.Command("kubectl", "-n", "stratus", "port-forward", fmt.Sprintf("deployment/%s", deploymentName), "5000")

	t.Cleanup(func() {
		t.Log("Killing kubectl port-forward process")
		// Make sure the process is killed before the test ends
		if err := portForwardCommand.Process.Kill(); err != nil {
			t.Log("failed to kill process: ", err)
		}
	})

	t.Log("Running port-forward")
	err := portForwardCommand.Start()  // Run the process in the background

	if err != nil {
		t.Error("Failed to start kubectl port-forward", err)
	}

	http_helper.HttpGetWithRetry(t, "http://localhost:5000", nil, 200, "Hello world!", 10, 3*time.Second)
}