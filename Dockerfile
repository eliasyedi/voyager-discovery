# syntax=docker/dockerfile:1
FROM golang:1.22.3 AS build-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN	goarch=amd64 goos=linux go build -o api .

#there should be a testing stage


#release stage
FROM scratch AS run-release-stage
EXPOSE 8080
COPY --from=build-stage /app/api /opt/api/
ENTRYPOINT ["/opt/api"]

