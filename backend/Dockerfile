FROM golang:1.21.1 as base


FROM base as development
WORKDIR /app
ENV STAGE=development
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
COPY backend/go.mod backend/go.sum ./
# DEVTOOLS
RUN go install -v golang.org/x/tools/gopls@latest
RUN go install -v github.com/ramya-rao-a/go-outline@latest
RUN go install -v github.com/go-delve/delve/cmd/dlv@latest
RUN go install -v honnef.co/go/tools/cmd/staticcheck@latest
RUN go install -v github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest
RUN go get -u golang.org/x/lint/golint
RUN go mod download
COPY backend .
CMD [ "air" ]


FROM base as builder
WORKDIR /app
COPY backend/go.mod backend/go.sum ./
RUN go mod download
COPY backend .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


FROM alpine:3.18.3 as production
RUN apk --no-cache add ca-certificates
WORKDIR /root/
ENV STAGE=production
COPY --from=builder /app/main .
COPY --from=builder /app/public ./public
COPY --from=builder /app/views ./views
CMD ["./main"]
