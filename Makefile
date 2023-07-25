.PHONY: all producer consumer clean

all: producer consumer

producer: 
	go build -o bin/producer cmd/producer/main.go

consumer:
	go build -o bin/consumer cmd/consumer/main.go

clean: 
	rm -rf bin
