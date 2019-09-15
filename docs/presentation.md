# Golang

## Introduction

Go is a relatively new programming language (first appeard in year 2009).
It is a statically typed, compiled programming language designed at Google
by Robert Griesemer, Rob Pike, and Ken Thompson.

Go is influenced by C and other languages (Pascal, Smalltalk, Java, Python).

Emphasis on greater simplicity, safety and productivity.

Theoretical pureness vs. real world practical situations.

Design to use modern hardware architectures (multi-cores).

---
## Famous applications

Some notable open-source applications written in Go include:

- Docker - a set of tools for deploying Linux containers

- Ethereum - the go-ethereum implementation of the Ethereum Virtual Machine blockchain for the Ether cryptocurrency

- InfluxDB - an open source database to handle time series data with HA and high performance requirements

- Kubernetes - a container management system

- OpenShift -  a cloud computing platform as a service by Red Hat

- Terraform - an open-source, multiple cloud infrastructure provisioning tool from HashiCorp

---
## Features

Go is easy to learn and retain.
It compiles a single binary for a single operating system (runtime, imported packages, entire application)

Go has:

- multiple return values
- modern standard library (networking, HTML templating, cryptography, data encoding, ...)
- goroutines and channels
- function types
- anonymous functions
- closures
- testing, code coverage, race-condition detection
- embedded coding convention (go fmt)
- effective and simple garbage-collection system (only 2 options)
- fast compilation process

---
## No good for

Go doesn't have:

- ternary operator (?:)
- generics

Go doesn't fit for:

- embedded systems
- desktop applications

---
## Code organization

Go code is developed in a **workspace**.
A workspace is a directory with three subdirectories:

- src contains Go source files (organized in projects and packages)
- pkg contains package objects
- bin contains executable binary files

To set up your workspace, you need to set the GOPATH environment variable.

## Set environment variables for Go

Go expects variable GOPATH to exist.

GOPATH tells Go where your workspace is.

---
## Start a program

Two ways:

1. ```go run main.go```

2. ```go build main.go```
   ```./main```

---

## References

Some important links:

- https://tour.golang.org
- https://play.golang.org


