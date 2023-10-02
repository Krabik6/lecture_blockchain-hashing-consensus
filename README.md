# Скриптики на GO
- mining
- sha256
- simplehash (5 symbols)

# Setup
Скачайте и установите git: [Download Git](https://github.com/git-guides/install-git#install-git-on-windows)

Cклонируйте репозиторий:
```bash
git clone https://github.com/Krabik6/lecture_blockchain-hashing-consensus.git
```

## Go
Install: [Download and install](https://go.dev/doc/install)
```bash
go version
```
### Simplehash (5 symbols)
```bash
go run ./simplehash/main.go
```
### Sha256
```bash
go run ./sha256/main.go
```
### Mining:
```bash
go run ./mining/main.go
```


## Docker

Install: [Install Docker Desktop on Windows](https://docs.docker.com/desktop/install/windows-install/)

### Simplehash:

Build:
```bash
docker build -t simplehash-image --target simplehash .
```

Run:
```bash
docker run -it simplehash-image
```

### Sha256:

Build:
```bash
docker build -t sha256-image --target sha256 .
```

Run:
```bash
docker run -it sha256-image
```

### Mining:

Build:
```bash
docker build -t mining-image --target mining .
```

Run:
```bash
docker run -it mining-image
```