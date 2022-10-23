WeChat Go SDK
=============

[![lint](https://github.com/aimuz/wgo/actions/workflows/lint.yml/badge.svg)](https://github.com/aimuz/wgo/actions/workflows/lint.yml)
[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://pkg.go.dev/github.com/aimuz/wgo)
[![Go Report Card](https://goreportcard.com/badge/github.com/aimuz/wgo)](https://goreportcard.com/report/github.com/aimuz/wgo)

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

