package main

import (
	"fmt"
	"strings"

	"analizadorLL1GO/controlador"
	"analizadorLL1GO/modelo"
)

var gramatica1 = modelo.Gramatica1

//var gramatica2 = modelo.Gramatica2
//var gramatica3 = modelo.Gramatica3
//var gramatica4 = modelo.Gramatica4

func main() {
	fmt.Printf(strings.Repeat("_", 30) + "\n")
	fmt.Printf("GRAMATICA ORIGINAL" + "\n")
	controlador.ImprimirGramatica(gramatica1)

	var gramaticaLista = controlador.EliminarRecursion(gramatica1)
	fmt.Printf(strings.Repeat("_", 30) + "\n")
	fmt.Printf("ELIMINAR RECURSION" + "\n")
	controlador.ImprimirGramatica(gramaticaLista)

	fmt.Printf(strings.Repeat("_", 30) + "\n")
	fmt.Printf("CONJUNTO DE PRIMEROS" + "\n")
	controlador.ImprimirGramatica(controlador.Primeros(gramaticaLista))

	fmt.Printf(strings.Repeat("_", 30) + "\n")
	fmt.Printf("CONJUNTO DE SIGUIENTES" + "\n")
	controlador.ImprimirGramatica(controlador.Primeros(gramaticaLista))

	fmt.Printf(strings.Repeat("_", 30) + "\n")
	fmt.Printf("CONJUNTO DE PREDICCION" + "\n")
	var conj_prediccion = controlador.Conjunto_prediccion(gramaticaLista)
	controlador.ImprimirGramatica(conj_prediccion)

	fmt.Printf(strings.Repeat("_", 30) + "\n")
	if controlador.EsLL1(conj_prediccion) {
		fmt.Printf("Es una gramatica LL1")
	} else {
		fmt.Printf("NO es una gramatica LL1")
	}

}
