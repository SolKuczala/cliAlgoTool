# cliAlgoTool
This can have a couple of solutions and I will choose 1 to showcase my ability.

## considerations
We need to normalize the data in some way, the sorting algorithm of golang will perform a lexicographically ordering and because the data of the column we need to sort contains numbers like
"9" and "10", it will consider 10 smaller than 9 because it starts with a "1".
We either add the necessary "0" to the numbers of 1 digit or compare the chars first and numbers second after converting to ints.
(first me agrega n, sec me agrega 2 nlogn) 
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

## How to run it
- 


enc := make([]byte, 256)
		return hex.Encode(enc, byte(bays[i])) < hex.Encode(enc, byte(bays[j]))
