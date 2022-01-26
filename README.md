# cdk8s for declarative management

This repo reproduces the [runnable example](https://github.com/vmware-tanzu/cartographer/tree/main/examples/runnable-tekton)
found in Cartographer.

Generating the `imports` directory is through the `cdk8s-cli` via `cdk8s import <my-yaml>.yaml -l go`.

TODO: Explore dependency graphs

### Setup

Install cert-manager:
```
kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v1.6.1/cert-manager.yaml
```

Install Tekton Pipelines:
```
kubectl apply --filename https://storage.googleapis.com/tekton-releases/pipeline/latest/release.yaml
```

Install Cartographer:
```
kubectl apply -f https://github.com/vmware-tanzu/cartographer/releases/latest/download/cartographer.yaml
```

Hack because unclear how JSON patching works:
```
kubectl apply -f https://raw.githubusercontent.com/vmware-tanzu/cartographer/main/examples/runnable-tekton/00-setup/tekton-task.yml
```

### Usage

Only an installation of Go is needed. `go run main.go -h` shows the inputs for this example.

`kubectl apply -f dist/*.yaml` each time the inputs change.

### Notes

Conversion to go code: JSII - build reserved key word so need temporarily replace `build` in OpenAPI spec
