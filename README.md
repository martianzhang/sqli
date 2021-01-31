# Introduction

`sqli` is a simple wrap of [libinjection](https://github.com/client9/libinjection) C library.

## Example 

Download vendor and build static library

```bash
# get package source code
go get -d github.com/martianzhang/sqli

# build static library
cd ~/go/sr/github.com/martianzhang/sqli && make test
```

Go code example

```go

import "github.com/martianzhang/sqli"

fingerprint, isInjection := sqli.SQLi("select 1")

```

Build & run example

```bash
make example
./example/main
```

## Run test case

```bash
make test
```


