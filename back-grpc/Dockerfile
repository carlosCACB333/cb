FROM golang:1.21.3 as base


FROM base as development
WORKDIR /app
ENV STAGE=development
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

COPY back-grpc/go.mod back-grpc/go.sum ./

#PROTOBUF COMPILER
RUN apt-get update && apt-get install -y protobuf-compiler
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# GRPC GATEWAY

RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go
RUN go get google.golang.org/grpc/cmd/protoc-gen-go-grpc

# DEVTOOLS
RUN go install -v golang.org/x/tools/gopls@latest
RUN go install -v github.com/ramya-rao-a/go-outline@latest
RUN go install -v github.com/go-delve/delve/cmd/dlv@latest
RUN go install -v honnef.co/go/tools/cmd/staticcheck@latest
RUN go install -v github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest
RUN go get -u golang.org/x/lint/golint
RUN go mod download
COPY back-grpc .
CMD [ "air"]


FROM base as builder
WORKDIR /app
COPY back-grpc/go.mod back-grpc/go.sum ./
RUN go mod download
COPY back-grpc .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


FROM alpine:3.18.3 as production
RUN apk --no-cache add ca-certificates
WORKDIR /root/
ENV STAGE=production
COPY --from=builder /app/main .
EXPOSE 8086
EXPOSE 8087
CMD ["./main"]
