# terratest-example

## Introduction

A collection of examples to investigate the use of the Terratest tool.

## Requirements

Golang - `brew install go`

Kubectl - `brew install kubectl`

## Running a test

`cd test`

`go test -v -run TestKubernetesExample -timeout 30m`

Example output:
```
=== RUN   TestKubernetesExample
=== PAUSE TestKubernetesExample
=== CONT  TestKubernetesExample
TestKubernetesExample 2020-10-22T10:32:40+02:00 logger.go:66: Running command kubectl with args [--namespace stratus apply -f ../examples/terraform-kubernetes-example/hello-world-deployment.yml]
TestKubernetesExample 2020-10-22T10:32:40+02:00 logger.go:66: deployment.apps/hello-world-deployment unchanged
TestKubernetesExample 2020-10-22T10:32:40+02:00 logger.go:66: service/hello-world-service unchanged
TestKubernetesExample 2020-10-22T10:32:40+02:00 retry.go:72: Wait for service hello-world-service to be provisioned.
TestKubernetesExample 2020-10-22T10:32:40+02:00 client.go:33: Configuring kubectl using config file /Users/mmwinther/.kube/config with context
TestKubernetesExample 2020-10-22T10:32:40+02:00 node.go:33: Getting list of nodes from Kubernetes
TestKubernetesExample 2020-10-22T10:32:40+02:00 client.go:33: Configuring kubectl using config file /Users/mmwinther/.kube/config with context
TestKubernetesExample 2020-10-22T10:32:40+02:00 service.go:86: Service is now available
    terraform_kubernetes_example_test.go:42: Running port-forward
TestKubernetesExample 2020-10-22T10:32:40+02:00 retry.go:72: HTTP GET to URL http://localhost:5000
TestKubernetesExample 2020-10-22T10:32:40+02:00 http_helper.go:32: Making an HTTP GET call to URL http://localhost:5000
TestKubernetesExample 2020-10-22T10:32:40+02:00 retry.go:84: HTTP GET to URL http://localhost:5000 returned an error: Get "http://localhost:5000": dial tcp [::1]:5000: connect: connection refused. Sleeping for 3s and will try again.
TestKubernetesExample 2020-10-22T10:32:43+02:00 retry.go:72: HTTP GET to URL http://localhost:5000
TestKubernetesExample 2020-10-22T10:32:43+02:00 http_helper.go:32: Making an HTTP GET call to URL http://localhost:5000
    terraform_kubernetes_example_test.go:35: Killing kubectl port-forward process
    terraform_kubernetes_example_test.go:23: Cleaning up deployment
TestKubernetesExample 2020-10-22T10:32:44+02:00 logger.go:66: Running command kubectl with args [--namespace stratus delete -f ../examples/terraform-kubernetes-example/hello-world-deployment.yml]
TestKubernetesExample 2020-10-22T10:32:44+02:00 logger.go:66: deployment.apps "hello-world-deployment" deleted
TestKubernetesExample 2020-10-22T10:32:44+02:00 logger.go:66: service "hello-world-service" deleted
--- PASS: TestKubernetesExample (30.64s)
PASS
ok  	github.com/mmwinther/terratest-example	31.069s
```
