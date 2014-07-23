RM=rm

COMPILE_DIR=$(shell pwd )
GOPATH=$(COMPILE_DIR)
all:
	@echo $(GOPATH)
	@cd ./src/
	go build ./...
	go install ./...
clean:
	$(RM) -rf ./bin/*
run:all
	echo "run" 
