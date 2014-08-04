BINPATH=$(GOPATH)\bin

all:
	go install
clean:
	@go clean
	@rm -rf $(BINPATH)\tmp.html
	@rm -rf $(BINPATH)\trstatus.exe
run:all
	cd $(BINPATH)
	$(BINPATH)\trstatus.exe>$(BINPATH)\tmp.html
