# cliAlgoTool
## How to run it
```make run-local```

### Possible choices for sorting algorithm
1. Use built in sort(quicksort) algo from golang - is an optimized version of quicksort for v1.22.

2. Use own implementation of sort(quicksort)

### Possible approaches:
1. Quicksort, then merge duplicates, then create CSV with new order.
```tc: o(n log n)(best case) + o(n) + o(n)```
2. Merge duplicates with the help of a map, then create new CSV, then quicksort ```tc: o(n) + o(n) + o(n log n)(best case)```
3. Merge duplicates with the help of a map , then quicksort, then create CSV with new order	```tc: o(n) + o(n log n)(best case) + o(n)```

### Decision:
Option ```3```: Because of the time and space complexity and it was easier to implement and understand to me.
With golang built-in solution, given the time constraint and that is more efficient than my implementation.

### Considerations
- I built the algorithm as a specific solution to the CSV given for the sake of time.
- I needed to remove a trailing comma manually from the CSV in order to be able to parse it; I couldn't find a good solution to normalize the CSV prior to parse it.
- I used logrus for logs as it's easy to set up
- I could have use assert pkg, but I wanted to keep it as vanilla as possible. 
