# Amun

This is a program that aims high performance stubbing through configuration

## Install and configure Go

### Download Go

```bash
curl -sLO https://dl.google.com/go/go1.11.5.linux-amd64.tar.gz ; \
    chmod +x go1.11.5.linux-amd64.tar.gz ; \
    tar -C /usr/local -xzf go1.11.5.linux-amd64.tar.gz
```

### Go location

```bash
export GOROOT=/usr/local/go
export GOPATH=$HOME/work/go
export PATH=$HOME/bin:/usr/local/bin:$GOPATH/go/bin:$GOROOT/bin:$PATH
```

## Build and Run It

### Locally

```bash
> go run $GOPATH/src/amun -file-path=$GOPATH/src/amun/config.yaml
```

### Dockerized

```bash
> docker build -t "amun:0.1" .
> docker run -d -p 9000:9000 --name "amun" amun:0.1
```
