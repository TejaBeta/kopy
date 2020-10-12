# kopy ![Go](https://github.com/TejaBeta/kopy/workflows/Go/badge.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/tejabeta/kopy)](https://goreportcard.com/report/github.com/tejabeta/kopy) [![License](https://img.shields.io/badge/License-Apache%202.0-green.svg)](./LICENSE)

A `kubectl` plugin to copy resources from a particular namespace from one context to another context.

## Idea, inception, and use-case

Imagine there is a provisioned **`K8s`** environment **`dev`**, and a user/developer wants to perform regression on a snippet of code in a particular namespace, but the environment is crowded with other team members testing, and working with several other microservices at the same time.

In order to avoid collision with other users/namespace resource changes within `K8s` environment, and to make life much more simplier, a tool to `Kopy` all the resources from a particular `K8s` environment to another `K8s` environment, this another environment could be anything from a large scale `K8s` cluster to a minikube hosted `K8s` environment running on a `spot-instance` just for one developer.

`Kopy` is a tool that just fits the above purpose. Hassle free, don't need to have complex CI/CD pipelines or multiple scripts to do the job. A simple, tiny, opensource tool to get things done.

Along with the above, there are several other use-cases one can explore.

Basic idea behind building this tool is to make a developer's life easy to test, and work along with serveral other microservices using `K8s` in an isolated environment. 

## Help Menus

```

kopy is a kubectl plugin to copy resources from one context to another context

kopy is a kubectl plugin or a cli to copy K8s resources
from one context to another context. Only requirement
here is to have both the contexts inside your kubectl 
config and appropriate accesses to the clusters

Usage:
  kopy [flags]

Flags:
  -d, --destination-context string   Destination Context name to copy resources into(required)
  -h, --help                         help for kopy
  -n, --ns string                    Namespace within the current context
  -s, --source-context string        Source Context name to copy resources from. If empty takes current context.

```

**`Ideas and contributions are always welcome 💪`**

## Future Improvements
- Support to allow multiple namespaces as input
- Copy resources into a different auto-generated/user-specified namespace to avoid collision

## Limitations
- Copy of stateful-sets is not supported
- Kubeconfig should have the source and destination contexts embedded
- Kopy has no capabilities to delete/update existing namespace or it's resources, tool leaves it to user with a human friendly message

## License
This project is distributed under the [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0), see [LICENSE](./LICENSE) for more information.
