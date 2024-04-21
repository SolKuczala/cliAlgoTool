# cliAlgoTool
This can have a couple of solutions and I will choose 1 to showcase my ability.

approaches
1. use built in sort(quicksort) algo from golang with custom sort - is an optimized version of quicksort

2. use own implementation of sort(quicksort)

a: I take de best case for qs everytime I count it
## Possible solutions I came up:
1. quicksort, merge duplicates, create CSV with new order // t: o(n log n) + n (if I create CSV while merging)
2. merge duplicates with the help of a map (and needs to save de code to link it), create new CSV and apply quicksort // t: n + n + n log n

## Desition:
1 - Because of the time and space complexity.

doubts:
- el mapa de golang sortea por key ya cuando guarda?no, is desordenado, si se quiere ordenar,
se tiene que sacar las keys aparte, sortear y ahi usar para mostrar el orden.
- sort en golang es quicksort?si, parece que es una version actualizada
hubo una implementacion dando vueltas que era mas rapido que la version anterior de golang, no se si la nueva es mejor o igual a la que dio vueltas.
- quicksort es el algo mas rapido para alfanumericos?

decisiones:
- creo que voy a hacer mi imple de QS solo para hacer showcase y si me queda tiempo implementar la de golang para tirar un benchmark (porque no estoy segura como implmentar dentro de la libreria de golang- me tomara mas tiempoy no se si llego a algun lado, el otro si me deja llegar a algun lado)