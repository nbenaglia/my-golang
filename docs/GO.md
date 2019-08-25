# Golang

## Introduction

Go is a relatively new programming language (first appeard in year 2009).
It is a statically typed, compiled programming language designed at Google
by Robert Griesemer, Rob Pike, and Ken Thompson.

Go is influenced by C, but with an emphasis on greater simplicity and safety.

Some notable open-source applications written in Go include:

- Docker, a set of tools for deploying Linux containers
- Ethereum, The go-ethereum implementation of the Ethereum Virtual Machine blockchain for the Ether cryptocurrency
- InfluxDB, an open source database specifically to handle time series data with high availability and high performance requirements.  
- Kubernetes container management system
- OpenShift, a cloud computing platform as a service by Red Hat
- Terraform, an open-source, multiple cloud infrastructure provisioning tool from HashiCorp.


Go isn’t considered a functional language, but it has some features that are common to functional languages, including:

- function types
- anonymous functions
- closures

## Code organization

Go code is developed in a **workspace**.
A workspace is a directory with three subdirectories:

- src contains Go source files (organized in projects and packages)
- pkg contains package objects
- bin contains executable binary files

To set up your workspace, you need to set the GOPATH environment variable.
