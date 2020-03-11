FROM golang:1.14.0-alpine3.11 as build-stage
WORKDIR /project
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o server main.go

FROM alpine:3.11.3
WORKDIR /project
COPY --from=build-stage /project /project
ENTRYPOINT ["./server"]
