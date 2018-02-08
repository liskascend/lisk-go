# Lisk-Go

[![Doc](https://godoc.org/github.com/slamper/lisk-go?status.svg)](http://godoc.org/github.com/slamper/lisk-go)
[![CircleCI](https://circleci.com/gh/Slamper/lisk-go.svg?style=svg)](https://circleci.com/gh/Slamper/lisk-go)
[![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](http://www.gnu.org/licenses/gpl-3.0)

Lisk GO is a Golang library for [Lisk - the cryptocurrency and blockchain application platform](https://github.com/LiskHQ/lisk). It allows developers to create offline transactions and broadcast them onto the network. It also allows developers to interact with the core Lisk API, for retrieval of collections and single records of data located on the Lisk blockchain. Its main benefit is that it does not require a locally installed Lisk node, and instead utilizes the existing peers on the network.

## Install
```
$ go get github.com/slamper/lisk-go
```

## Tests

```
go test -v ./...
```


## Lint

```
gometalinter --config=lint.json ./...
```

Use this command to run the required linters.

## Authors

- Hendrik Hofstadt <dev@slamper.me>