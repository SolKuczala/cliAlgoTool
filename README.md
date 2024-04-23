# cliAlgoTool
## How to run it
```bash
make run-local
```
It will build and run the binary taking the csv provided to make the output to `csv` folder.

It also prints the output as it has the `-print` option.
To run it manually, see target in Makefile.

### Usage
Once the binary is built run:

```bash
./clialgotool <flag> [path-to-csv-file]
```

#### Flags:
- `-input` Takes a path to a CSV to process.
- `-output` Takes a name of a file to write the output. If it is not provided it will use a default name.
- `-print` Prints the output to console.

### To run tests and benchmarks
```bash
make test-local-unit
make test-local-integration
make run-benchmark
```

#### Dockerfile
To try the docker build target, make sure docker is in your sudoer group.
This was just and addition to show the usage.

### Possible choices for sorting algorithm
1. Use Golang built in sort(quicksort) algorithm - it has an optimized version of quicksort for v1.22.

2. Use own implementation of sort(quicksort).

### Possible approaches:
1. Quicksort, then merge duplicates, then create CSV with new order.
Time complexity, best case: `o(n log n) + o(n) + o(n)`

2. Merge duplicates with the help of a map, then create new CSV, then quicksort
Time complexity, best case: `o(n) + o(n) + o(n log n)`

3. Merge duplicates with the help of a map , then quicksort, then create CSV with new order.
Time complexity, best case: `o(n) + o(n log n) + o(n)`

### Decision:
I decided to go for option `3` because the time and space complexity are better and it was easier to implement and reason about.

Of the two sorting algorithm choices I decided to go with golang's built-in quicksort, given the time constraint and the fact that it is more efficient than what my own implementation would be.

### Considerations
- I built the algorithm as a specific solution to the CSV given for the sake of time. For a production codebase this should be implemented to support changes in the future, like not depending on a specific column arrangement and normalizing the product codes and locations in a more generic way.
- I had to remove a trailing comma manually from the CSV in order to be able to parse it.