FROM golang:alpine3.11 as build-stage
RUN apk add build-base
WORKDIR /project
COPY . .
RUN GOOS=linux CGO_ENABLED=1 go build -mod vendor -ldflags "-s -w" -o server cmd/main.go

FROM alpine:3.11.3
WORKDIR /project
COPY --from=build-stage /project /project
ENTRYPOINT ["./server"]
