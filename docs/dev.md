# Development Document

Welcome to contribute code to this project. Below are some information you need to know.

## Project structure

* client: Monitoring service client
* control: Monitoring service master control
* docs: Project documentation
* out: Output product

## Start development

Please execute `go mod download` to download dependencies under the project directory.

## Run the code

You cannot use `go run` directly because the project uses a daemon to call the business process, and running it directly will cause an exception.

Please [compile and run](#compile-and-package) using the command shown in the last step

## Compile and package

You can use the following command to directly compile the product of the current platform
```shell
cd client
go build -o ../out/monitor-client -ldflags '-s -w'

cd control
go build -o ../out/monitor-control -ldflags '-s -w'
```

or you can use the following command to cross-compile the product of other platforms

```shell
cd client
# Linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../out/monitor-client -ldflags '-s -w'
# Windows
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ../out/monitor-client -ldflags '-s -w'

cd control
# Linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../out/monitor-control -ldflags '-s -w'
# Windows
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ../out/monitor-control -ldflags '-s -w'
```