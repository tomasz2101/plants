deps:
	dep ensure
run:
	go run main.go
build_docker:
	docker build -t go ./docker
build:
	docker run --name rpi_build --rm -it -v ~/go/src/plants:/root/go/src/plants --workdir=/root/go/src/plants go sh -c "CC=arm-linux-gnueabi-gcc CGO_ENABLED=1 GOOS=linux GOARCH=arm GOARM=6 go build"
send:
	scp ./plants pi@10.5.5.94:/tmp/testing