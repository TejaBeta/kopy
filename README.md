# kopy [![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg)](https://github.com/gomods/athens)

**`Work in progress`**

A `kubectl` plugin to copy resources from a particular namespace from one context to another context. 

## Help Menus

```

$kopy --help
kopy is a kubectl plugin to copy resources from one context to another context

kopy is a kubectl plugin or a cli to copy K8s resources
from a particular namespace from the current context to
another context. However, the context switching is kind
of opinionated, as the tool requires all the contexts 
within the same config.

Usage:
  kopy [flags]

Flags:
      --a                            All the resources within the namespace
  -d, --destination-context string   Destination Context name to copy resources into(required)
  -h, --help                         help for kopy
  -n, --ns string                    Namespace within the current context

```

**`Ideas are welcome, please drop me a message`**

## License
This project is distributed under the [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0), see [LICENSE](./LICENSE) for more information.
