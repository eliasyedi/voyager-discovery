

BINARY_NAME=voyager


build: 
	GOARCH=amd64 GOOS=linux go build -o ./target/BINARY_NAME


run: build
	./target/BINARY_NAME


clean:
	rm ./target/BINARY_NAME 
