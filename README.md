# Lisk-Go

[![Doc](https://godoc.org/github.com/liskascend/lisk-go?status.svg)](http://godoc.org/github.com/liskascend/lisk-go)
[![CircleCI](https://circleci.com/gh/liskascend/lisk-go.svg?style=svg)](https://circleci.com/gh/liskascend/lisk-go)
[![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](http://www.gnu.org/licenses/gpl-3.0)

Lisk-Go is a Golang library for [Lisk - the cryptocurrency and blockchain application platform](https://github.com/LiskHQ/lisk) released by [Lisk Ascend](https://liskascend.com). It allows developers to create offline transactions and broadcast them onto the network. It also allows developers to interact with the core Lisk API, for retrieval of collections and single records of data located on the Lisk blockchain. Its main benefit is that it does not require a locally installed Lisk node, and instead utilizes the existing peers on the network.

Currently this library is in **Beta** and the API will experience changes.

Transaction and Crypto code have >95% test coverage and are extensively tested.

Features:
- [X] Support for the latest 1.0 Lisk API
- [X] 100% coverage of the API functions
- [X] Very detailed documentation
- [X] Usage of Go contexts for requests
- [X] Pretty printing of API errors + detailed errors for internal functions
- [X] Use of Go's native data types + a lot of helper structs
- [X] Modular layout
- [X] Full Test Coverage of Crypto
- [X] Full Test Coverage of Transaction Logic/Serialization
- [ ] Full Test Coverage of API Wrapper

## Install
```
$ go get github.com/liskascend/lisk-go
```

## Usage

For detailed documentation consider the GoDoc linked above.

This project is consists of 3 modules/packages:
* `api` - Module used to communicate with the Lisk 1.0 API
* `crypto` - Module which implements the core cryptography functions required for using Lisk
* `transactions` - Module which implements transaction and payload serialization and validation

#### Sending a simple transaction

The library offers comfortable util constructors for all supported transaction types. 

They are in the transactions package and prefixed with `New`

The following example creates+signs a transaction and broadcasts it to the network.
```
// Create the client
client := api.NewClient()
// Create the transaction using the constructor utils
transaction, err := transactions.NewTransactionWithData("104666L", 0, "wagon stock borrow episode laundry kitten salute link globe zero feed marble", "", 0, "abc")
if err != nil {
	// handle error
	return
}
res, err := client.SendTransaction(context.Background(), transaction)
if err != nil {
	// handle error
	return
}
```

This is the equivalent but done manually:
```
// Create the client
client := api.NewClient()

timestamp := GetCurrentTimeWithOffset(0)

transaction := &Transaction{
	Type:        TransactionTypeNormal,
	Amount:      0,
	RecipientID: "104666L",
	Timestamp:   timestamp,
	Asset:       transactions.DataAsset("abc"),
}

secret := "wagon stock borrow episode laundry kitten salute link globe zero feed marble"

pubKey := crypto.GetPublicKeyFromSecret(secret)
transaction.SenderPublicKey = pubKey

privKey := crypto.GetPrivateKeyFromSecret(secret)
transaction.Sign(privKey)

res, err := client.SendTransaction(context.Background(), transaction)
if err != nil {
	// handle error
	return
}
```

Manual usage of the transaction struct + assets can be used for more complex use-cases.

This library offers intensive validation of the transaction which is automatically performed before serialization 
or when ``isValid()`` is called on transactions or assets.

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

- Hendrik "Slamper" Hofstadt <dev@slamper.me> ``16863632246347444618L``
