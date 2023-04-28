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

	fmt.Printf("\n")
	fmt.Printf(strings.Repeat("_", 30) + "\n")
	cadena := "id + id * id"
	analisisSintacticoRecursivo(cadena, gramatica1)
}

func analisisSintacticoRecursivo(cadena string, gramatica []map[string][]string) bool {

	if controlador.EsLL1(gramatica1) {
		conjunto_pred := controlador.Conjunto_prediccion(gramatica1)

		// Construir la tabla de análisis sintáctico
		tabla := make(map[string]map[string]string)
		for _, produccion := range conjunto_pred {
			for no_terminal, predicciones := range produccion {
				tabla[no_terminal] = make(map[string]string)
				for _, prediccion := range predicciones {
					tabla[no_terminal][prediccion] = controlador.Keys(controlador.BuscarProduccionAnalizador(no_terminal, gramatica1))[0]
				}
			}
		}

		// Analizador sintáctico LL(1)
		pila := []string{"$", controlador.Keys(controlador.BuscarProduccionAnalizador("E", gramatica1))[0]}
		cadena = cadena + " $"
		i := 0
		for len(pila) > 0 {
			if pila[len(pila)-1] == cadena[i:i+1] {
				pila = pila[:len(pila)-1]
				i++
			} else {
				no_terminal := pila[len(pila)-1]
				entrada := cadena[i : i+1]
				if _, ok := tabla[no_terminal][entrada]; ok {
					pila = pila[:len(pila)-1]
					if tabla[no_terminal][entrada] != "λ" {
						produccion := strings.Split(tabla[no_terminal][entrada], " ")
						pila = append(pila, controlador.Reverse(produccion)...)
					}
				} else {
					fmt.Println("La cadena no pertenece a la gramática")
					return false
				}
			}
		}
		if len(pila) == 0 && i == len(cadena)-1 {
			fmt.Println("La cadena pertenece a la gramática")
			return true
		}
	} else {
		fmt.Println("La gramática no es LL(1)")
		return false
	}
	return false
}
