FROM golang:1.18-alpine as builder

WORKDIR /work

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o app ./cmd/app/main.go


FROM alpine:latest

ENV TZ Asia/Shanghai
RUN export GIN_MODE=release

WORKDIR /server

COPY --from=builder /work/app ./
COPY --from=builder /work/resource ./resource

EXPOSE 3000

CMD ["./app", "-f", "/resource/application.prod.yaml"]
