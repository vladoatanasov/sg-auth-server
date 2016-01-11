buildit:
	go get && go build -ldflags "-w" -o bin/auth_server
clean:
	rm -fr bin/*
buildclean: clean buildit
cleanbuild: clean buildit
test:
	go get && go test
