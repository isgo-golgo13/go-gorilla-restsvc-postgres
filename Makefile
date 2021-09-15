SERVICE_NAME=service

.PHONY : clean 
.DEFAULT_GOAL : all

all:
	
compile: 
	go build github.com/isgo-golgo13/go-gorilla-restsvc-postgres/cmd/service

clean: 
	rm -f ${SERVICE_NAME}

test:
	go test -v