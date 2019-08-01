# Amun

This is a program that aims high performance stubbing through configuration

## Install and configure Go

Install [Go](https://golang.org/dl/)

## Build and Run It

### Locally

```bash
git clone https://github.com/miguelnv/amun.git
cd amun
go build
./amun -file-path=custom/config.yaml
```

### Dockerized

```bash
docker build -t "amun:0.1" .
docker run -d -p 9000:9000 --name "amun" amun:0.1
```

## Usage

### Matching parameter and header

```bash
curl http://localhost:9000/test/123?action=test -H "X-test: val2"
```

### With 200 milliseconds of latency

```bash
curl http://localhost:9000/test/123?action=test -H "X-test: val2" -H "X-Amun-Latency: 200ms"
```

### With 1 second of latency

```bash
curl http://localhost:9000/test/123?action=test -H "X-test: val2" -H "X-latency: 1s"
```

## Profiling it

### with wrk2

```bash
./wrk -t2 -c100 -d30s -L -R2000 http://127.0.0.1:9000/test/123\?action\=test -H "X-test: val2"
```
