# syntax=docker/dockerfile:1
FROM golang:1.22.3 AS build-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN	goarch=amd64 goos=linux go build -o api .

#there should be a testing stage


#release stage
FROM go:1.22.3-alpine AS run-release-stage
WORKDIR /app
RUN apk add libc6-compat
EXPOSE 8080
COPY --from=build-stage /app/api . 
COPY *.env . 
ENTRYPOINT ["/app/api"]
