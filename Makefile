build:
	go build -o ./bin/demoparser ./cmd/

run:
	./bin/demoparser -logLevel=debug -protoFile=demo/proto.rbs

clean:
	rm -rf imgs/*