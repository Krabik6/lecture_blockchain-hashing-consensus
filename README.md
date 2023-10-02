# Mining example
This is a simple example of how work mining proof of work.

## Requirements:
### Local machine
- [Go](https://golang.org/doc/install)


### Docker
- [Docker](https://docs.docker.com/desktop/install/windows-install/)

## Setup:

### Docker
Build the docker image:
```bash
docker build -t mining-example .
```

Run the docker image:
```bash
docker run -it mining-example
```

### Go
Build the project:
```bash
go run main.go
```