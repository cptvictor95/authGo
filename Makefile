build:
	go build -o bin/authGo ./cmd/authGo

run: build
	./bin/authGo

clean:
	rm bin/*
