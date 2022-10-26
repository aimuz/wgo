WeChat Go SDK
=============

[![lint](https://github.com/aimuz/wgo/actions/workflows/lint.yml/badge.svg)](https://github.com/aimuz/wgo/actions/workflows/lint.yml)
[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://pkg.go.dev/github.com/aimuz/wgo)
[![Go Report Card](https://goreportcard.com/badge/github.com/aimuz/wgo)](https://goreportcard.com/report/github.com/aimuz/wgo)
[![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/6613/badge)](https://bestpractices.coreinfrastructure.org/projects/6613)

**In Development**

A Go SDK for the WeChat Open Platform API.

Usage
-----

### Client

```go
import "github.com/aimuz/wgo"
```

```go
client := wgo.NewClient(wgo.WithAPPIDAndSecret("<APPID>", "<SECRET>"))
```

Roadmap
-------

 [x] TODO

Contributors
------------

Pull requests are always welcome! Please base pull requests against the `main` branch and follow the [contributing guide](CONTRIBUTING.md).

Feedback
--------

We’d love to hear your thoughts on this project. Feel free to drop us a note!
    
- [Issues](https://github.com/aimuz/wgo/issues)

License
-------

```text
Copyright 2022 The WGo Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
