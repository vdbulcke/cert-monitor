FROM golang:1.15.3-alpine3.12 as build
WORKDIR /go/src/certificate-monitor

# Install the Protocol Buffers compiler and Go plugin
RUN apk add protobuf git make zip
RUN go get github.com/golang/protobuf/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc

ENV GO111MODULE=on
# <- COPY go.mod and go.sum files to the workspace
COPY go.mod . 
COPY go.sum .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download

COPY . .


RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -o  cert-monitor ./main.go

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
## non privileged user
USER 1111 
WORKDIR /app/
COPY --from=0 /go/src/certificate-monitor/cert-monitor .
ENTRYPOINT ["/app/cert-monitor", "-config",  "/app/config.yaml"] 