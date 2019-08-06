git submodule update --recursive --remoteIMAGES = go_builder
TARGET = plants
include $(dir $(lastword ${MAKEFILE_LIST}))/tools/Makefile

MQTT_USERNAME?=internal
MQTT_PASSWORD?=internal
MQTT_HOSTNAME?=localhost
MQTT_PORT?=1883

dep_init:
	dep init

deps:
	dep ensure

run:
	go run main.go -u $(MQTT_USERNAME) -p $(MQTT_PASSWORD) -d testing -H $(MQTT_HOSTNAME) -P $(MQTT_PORT)

test:
	docker run --name rpi_build --rm -it \
	-v ${PWD}:/root/go/src/${TARGET} --workdir=/root/go/src/${TARGET} \
	tomasz2101/go_builder sh -c "CC=arm-linux-gnueabi-gcc CGO_ENABLED=1 GOOS=linux GOARCH=arm GOARM=6 go build"
send:
	scp ./${TARGET} pi@10.5.5.79:/tmp/${TARGET}
