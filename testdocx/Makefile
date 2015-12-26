BIN = testhtml
PWD = $(shell pwd)
DIR_NAME = $(shell dirname ${PWD})
ORG_PATH = github.com/mingqing
REPO_PATH = ${ORG_PATH}/${BIN}

BUILD_OS_TARGETS = "linux"

GOPATH := ${DIR_NAME}/gopath

all: clean build run

build:
	mkdir build
	mkdir -p ${GOPATH}/src/${ORG_PATH}
	ln -s ${PWD} ${GOPATH}/src/${REPO_PATH}
	GOPATH=$(GOPATH) go build -o build/${BIN}

clean:
	rm -rf build
	rm -f ${GOPATH}/src/${REPO_PATH}

run:
	build/${BIN}
