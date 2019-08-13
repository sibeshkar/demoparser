build:
	go build -o ./bin/demoparser ./cmd/

run:
	./bin/demoparser -fps=20 -speedup=1.0

clean:
	rm -rf imgs/*