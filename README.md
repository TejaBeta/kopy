# kopy [![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg)](https://github.com/gomods/athens)

A `kubectl` plugin to copy resources from a particular namespace from one context to another context. 

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

## License
This project is distributed under the [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0), see [LICENSE](./LICENSE) for more information.
