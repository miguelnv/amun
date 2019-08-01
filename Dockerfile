# STEP 1 build executable binary
FROM golang:alpine as builder

# Install git
# RUN apk update && apk add git
RUN apk update && \
    apk add git && \
    apk add ca-certificates && \
    adduser -D -g '' amun

#COPY . $GOPATH/src/amun/
#WORKDIR $GOPATH/src/amun

WORKDIR /amun
COPY go.mod go.sum ./

# get dependencies
#RUN go get -d -v
# Using go mod with go 1.11
RUN go mod download
RUN go mod verify

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# build the binary
# removing debug informations and compile only for linux target and disabling cross compilation
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o amun

#RUN ls -la

# STEP 2 build a small image

# start from scratch
FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd

# expose yaml file
COPY --from=builder /amun/config.yaml /
# Copy our static executable
COPY --from=builder /amun/amun /

USER amun
EXPOSE 9000
ENTRYPOINT ["/amun"]