# STEP 1 build executable binary

FROM golang:alpine as builder

# Install git
# RUN apk update && apk add git
RUN apk update && \
    apk add git && \
    apk add ca-certificates && \
    adduser -D -g '' amun

COPY . $GOPATH/src/amun/
WORKDIR $GOPATH/src/amun

# get dependencies
RUN go get -d -v

# build the binary
# removing debug informations and compile only for linux target and disabling cross compilation
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o $GOPATH/bin/amun

# build the binary
# RUN go build -o $GOPATH/bin/amun

# STEP 2 build a small image

# start from scratch
FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd

# expose yaml file
COPY --from=builder /go/src/amun/config.yaml /config.yaml
# Copy our static executable
COPY --from=builder /go/bin/amun /go/bin/amun

USER amun
EXPOSE 9000
ENTRYPOINT ["/go/bin/amun"]