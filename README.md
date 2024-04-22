# cliAlgoTool
## How to run it
```make run-local```

### Possible approaches
This problem have a couple of solutions and I will choose 1 to showcase my ability.

1. use built in sort(quicksort) algo from golang with custom sort - is an optimized version of quicksort

2. use own implementation of sort(quicksort)

### Considerations
- I built the algorithm as a specific solution to the CSV given for the sake of time.
- I needed to remove a trailing comma manually from the CSV in order to be able to parse it; I couldn't find a good solution to normalize the CSV prior to parse it.


### Possible solutions I came up:
1. Quicksort, then merge duplicates, then create CSV with new order.
```tc: o(n log n)(best case) + o(n) + o(n)```
2. Merge duplicates with the help of a map, then create new CSV, then quicksort ```tc: o(n) + o(n) + o(n log n)(best case)```
3. Merge duplicates with the help of a map , then quicksort, then create CSV with new order	```tc: o(n) + o(n log n)(best case) + o(n)```

### Decision:
Option ```3```: Because of the time and space complexity and it was easier to implement and understand to me.
