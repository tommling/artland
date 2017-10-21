ifndef VERBOSE
.SILENT:
endif

all: server

server:
	go build server.go

clean:
	rm server

.PHONY: all, clean
