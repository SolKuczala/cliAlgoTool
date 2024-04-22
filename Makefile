build-local:
	go mod tidy
	go build .

run-local: build-local
	./clialgotool -input ./csv/hb_test.csv -print

test-local-unit:
	go test -v ./...

test-local-integration: build-local
	./clialgotool -input ./integration/test_input.csv -output ./integration/testrun_output.csv
	diff ./integration/expected_output.csv ./integration/testrun_output.csv
	rm ./integration/testrun_output.csv

run-docker-test:
	docker build . -t test
	docker run --rm test

run-benchmark:
	go test -bench=. -benchmem ./...