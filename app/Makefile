VERSION='1.0.0'
CONFIG_FILE=production_config.json
OUTPUT_DIR=bin
BINARY=mango-app


docker-make: clean build docker

build:
	@if ! [ -e ${OUTPUT_DIR} ] ; then mkdir ${OUTPUT_DIR} ; fi
	@cp ${CONFIG_FILE} ${OUTPUT_DIR}/
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-w -s' -o ${OUTPUT_DIR}/${BINARY} .

docker:
	@docker build -t ${BINARY}:${VERSION} .

docker-run:
	@docker run --rm -d --name=${BINARY} -p 8080:8080 ${BINARY}:${VERSION}

docker-stop:
	@docker kill ${BINARY}

run:
	@go run main.go

clean:
	@if [ -e ${OUTPUT_DIR} ] ; then rm -rf ${OUTPUT_DIR} ; fi

info:
	@echo VERSION : ${VERSION}
	@echo CONFIG_FILE : ${CONFIG_FILE}
	@echo OUTPUT_DIR : ${OUTPUT_DIR}
	@echo BINARY : ${BINARY}