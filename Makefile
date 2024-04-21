build-local:
	go mod tidy && go build .

run-local: build-local
	./clialgotool -input ./csv/hb_test.csv -print

test-local:
	go test -v ./...

run-docker:
	docker build . -t test && docker run --rm test