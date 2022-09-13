from pprint import pprint
conjunto = [1,2,3]
k = 2  
lista = []
iteraciones = [0]
def subconjuntos(l, k):
    if k == len(l):
        if not l in lista:
            lista.append(l)
        return
    for i in l:
        aux = l[:]
        aux.remove(i)
        result = subconjuntos(aux, k)
        iteraciones[0] += 1
        if not result in lista and result:
            lista.append( result)

subconjuntos(conjunto, k)
print (lista)
print ('cant iteraciones: ' + str(iteraciones[0]))
for j in lista:
	for i in j:
		sum = i
		print (sum)