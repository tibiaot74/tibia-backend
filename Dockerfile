## Build
FROM golang:1.19-alpine3.16 as build
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN GOOS=linux go build -o app . 

## Deploy
FROM alpine:latest
WORKDIR /app
COPY --from=build /build/app .
RUN chmod +x ./app
EXPOSE 7474
ENTRYPOINT ["./app"]
