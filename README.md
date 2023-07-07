# Golang

## Golang and MUX router REST API

## Install Golang - Linux

```bash
sudo tar -C /usr/local -xzf go1.13.5.linux-amd64.tar.gz

vim ~/.profile

(append :/usr/local/go/bin to PATH)
```

## Install Mux library

```bash
go get -u github.com/gorilla/mux
```

## Install Chi library

```bash
go get -u github.com/go-chi/chi
```

## Install Firestore library

```bash
go get -u cloud.google.com/go/firestore
```

## Install MySQL library

```bash
go get -u github.com/go-sql-driver/mysql
```

## Install MongoDB library

```bash
go get -u go.mongodb.org/mongo-driver
```

## Install SQLite 3 Library

```bash
go get -u github.com/mattn/go-sqlite3
```

## Install Testify library

```bash
go get github.com/stretchr/testify
```

## Export Environment variable 

```bash
export GOOGLE_APPLICATION_CREDENTIALS='/path/to/project-private-key.json'
```

## How to get the private key JSON file:
## From the Firebase Console: Project Overview -> Project Settings -> Service Accounts -> Generate new private key

## Build

```bash
go build
```
## Test (specific test)

```bash
go test -run NameOfTest
```

## Test (all the tests within the service folder)

```bash
go test service/*.go
```

## Run

```bash
go run .
```

```bash
go run *.go
```
