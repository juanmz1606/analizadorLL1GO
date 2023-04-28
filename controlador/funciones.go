package controlador

import (
	"fmt"
	"strings"
	"unicode"
)

func BuscarProduccion(nombre_produccion string, gramatica []map[string][]string) []string {
	elementos_produccion := make([]string, 0)
	for _, prod := range gramatica {
		if key, ok := prod[nombre_produccion]; ok {
			elementos_produccion = append(elementos_produccion, key...)
		}
	}
	return elementos_produccion
}

func RemoveElement(slice []string, element string) []string {
	result := make([]string, 0, len(slice))
	for _, s := range slice {
		if s != element {
			result = append(result, s)
		}
	}
	return result
}

func ListaNoTerminales(gramatica []map[string][]string) []string {
	lista_nt := make([]string, 0)
	for _, produccion := range gramatica {
		for key := range produccion {
			if !stringInSlice(key, lista_nt) {
				lista_nt = append(lista_nt, key)
			}
		}
	}
	return lista_nt
}

func ListaTerminales(gramatica []map[string][]string) []string {
	lista_t := []string{"$"}
	for _, produccion := range gramatica {
		for _, value := range produccion {
			for _, derivado := range value {
				arr_derivado := strings.Split(derivado, " ")
				for _, el := range arr_derivado {
					if (unicode.IsLower([]rune(el)[0]) || !unicode.IsLetter([]rune(el)[0])) && el != "λ" {
						if !stringInSlice(el, lista_t) {
							lista_t = append(lista_t, el)
						}
					}
				}
			}
		}
	}
	return lista_t
}

func Conjunto_prediccion(gramatica []map[string][]string) []map[string][]string {
	lista_primeros := Primeros(gramatica)
	lista_siguientes := Siguientes(gramatica)

	lista_t := ListaTerminales(gramatica)

	conjunto_pred := []map[string][]string{}
	valores_tabla_conjunto_pred := [][]string{}

	for _, produccion := range gramatica {
		conjunto_prediccion_prod_actual := []string{}
		lista_tabla_prod_actual := []string{Keys(produccion)[0]}

		for key, value := range produccion {

			for i := 0; i < len(lista_t); i++ {
				lista_tabla_prod_actual = append(lista_tabla_prod_actual, " ")
			}

			for _, derivado := range value {

				arr_value := strings.Split(derivado, " ")

				if arr_value[0] == "λ" {
					siguientes_esta_prod := BuscarProduccion(key, lista_siguientes)
					conjunto_prediccion_prod_actual = append(conjunto_prediccion_prod_actual, siguientes_esta_prod...)
				} else if unicode.IsLower(rune(arr_value[0][0])) || !unicode.IsLetter(rune(arr_value[0][0])) {
					conjunto_prediccion_prod_actual = append(conjunto_prediccion_prod_actual, arr_value[0])
				} else {
					primeros_esta_prod := BuscarProduccion(arr_value[0], lista_primeros)
					conjunto_prediccion_prod_actual = append(conjunto_prediccion_prod_actual, primeros_esta_prod...)
				}
			}
		}
		valores_tabla_conjunto_pred = append(valores_tabla_conjunto_pred, lista_tabla_prod_actual)
		conjunto_pred = append(conjunto_pred, map[string][]string{Keys(produccion)[0]: conjunto_prediccion_prod_actual})
	}

	return conjunto_pred
}

// isLower verifica si un caracter es una letra minúscula
func isLower(c string) bool {
	return c >= "a" && c <= "z"
}

// isAlnum verifica si un caracter es alfanumérico
func isAlnum(c string) bool {
	return (c >= "a" && c <= "z") || (c >= "A" && c <= "Z") || (c >= "0" && c <= "9")
}

// indexOf devuelve el índice de un elemento en un arreglo de strings, o -1 si no lo encuentra
func indexOf(arr []string, elem string) int {
	for i, v := range arr {
		if v == elem {
			return i
		}
	}
	return -1
}

func ListKeys(m map[string][]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func contains(arr []string, element string) bool {
	for _, v := range arr {
		if v == element {
			return true
		}
	}
	return false
}

func Keys(m map[string][]string) []string {
	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func GetIndex(val string, arr []string) int {
	for i, v := range arr {
		if v == val {
			return i
		}
	}
	return -1
}

func stringInSlice(s string, slice []string) bool {
	for _, value := range slice {
		if value == s {
			return true
		}
	}
	return false
}

func EsLL1(conjunto_prediccion []map[string][]string) bool {
	esLL1 := true

	for _, produccion := range conjunto_prediccion {
		repetidosLL1 := false
		for _, value := range produccion {
			prodSinRepetidos := make(map[string]bool)

			for _, derivado := range value {
				arrDerivado := strings.Split(derivado, " ")
				primero := arrDerivado[0]

				if prodSinRepetidos[primero] {
					repetidosLL1 = true
					break
				}

				prodSinRepetidos[primero] = true
			}

			if repetidosLL1 {
				esLL1 = false
				break
			}
		}

		if !esLL1 {
			break
		}
	}

	return esLL1
}

// Función que realiza el análisis sintáctico de una cadena con una gramática LL(1) no recursiva.
func BuscarProduccionAnalizador(nombre_produccion string, gramatica []map[string][]string) map[string][]string {
	for _, produccion := range gramatica {
		for no_terminal, derivados := range produccion {
			if no_terminal == nombre_produccion {
				return map[string][]string{no_terminal: derivados}
			}
		}
	}
	return nil
}

// Función que retorna una copia del slice dado, pero en orden inverso.
func Reverse(s []string) []string {
	r := make([]string, len(s))
	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
		r[i], r[j] = s[j], s[i]
	}
	return r
}

// Función que retorna true si el slice contiene el elemento dado, y false en caso contrario.
func Contiene(s []string, elem string) bool {
	for _, v := range s {
		if v == elem {
			return true
		}
	}
	return false
}

func ImprimirGramatica(gramatica []map[string][]string) {
	for _, element := range gramatica {
		for key, value := range element {
			fmt.Printf("%v -> [%v]\n", key, strings.Join(value, ","))
		}
	}
}

func EliminarRecursion(gramatica []map[string][]string) []map[string][]string {
	for _, prod := range gramatica {
		alfas := []string{}
		betas := []string{}
		for key, values := range prod {
			for _, value := range values {
				if string(value[0]) == key {
					for _, char := range values {
						if string(char[0]) == key {
							alfas = append(alfas, strings.TrimSpace(char[1:]))
						} else {
							betas = append(betas, char)
						}
					}
					nombreNuevaProd := key + "p"
					prod[key] = []string{}
					elementosNuevaProd := []string{}

					for _, beta := range betas {
						prod[key] = append(prod[key], strings.TrimSpace(beta)+" "+nombreNuevaProd)
					}
					for _, alfa := range alfas {
						elementosNuevaProd = append(elementosNuevaProd, strings.TrimSpace(alfa)+" "+nombreNuevaProd)
					}
					elementosNuevaProd = append(elementosNuevaProd, "λ")
					gramatica = append(gramatica, map[string][]string{nombreNuevaProd: elementosNuevaProd})
				}
			}
		}
	}
	return gramatica
}

func Siguientes(gramatica []map[string][]string) []map[string][]string {
	lista_primeros := Primeros(gramatica)
	lista_siguientes := []map[string][]string{}

	indice_produccion_actual := 0

	for indice_produccion_actual < len(gramatica) {
		nt_prod_actual := ListKeys(gramatica[indice_produccion_actual])[0]
		siguientes_prod_actual := []string{}
		for _, produccion := range gramatica {
			for nt, derivados := range produccion {
				for _, derivado := range derivados {
					arr_derivado := strings.Split(derivado, " ")
					if contains(arr_derivado, nt_prod_actual) {
						indice_nt_actual := indexOf(arr_derivado, nt_prod_actual)

						if indice_produccion_actual == 0 {
							siguientes_prod_actual = append(siguientes_prod_actual, "$")
						}

						if !(indice_nt_actual == len(arr_derivado)-1) {

							if isLower(arr_derivado[indice_nt_actual+1]) || !isAlnum(arr_derivado[indice_nt_actual+1]) {

								if !contains(siguientes_prod_actual, arr_derivado[indice_nt_actual+1]) {
									siguientes_prod_actual = append(siguientes_prod_actual, arr_derivado[indice_nt_actual+1])
								}
							} else {
								lista_primeros_del_siguiente := BuscarProduccion(arr_derivado[indice_nt_actual+1], lista_primeros)
								if contains(lista_primeros_del_siguiente, "λ") {
									lista_primeros_del_siguiente = RemoveElement(lista_primeros_del_siguiente, "λ")
									siguientes_raiz := BuscarProduccion(nt, lista_siguientes)
									for _, siguiente := range siguientes_raiz {
										if !contains(lista_primeros_del_siguiente, siguiente) {
											lista_primeros_del_siguiente = append(lista_primeros_del_siguiente, siguiente)
										}
									}
								}
								for _, sig := range lista_primeros_del_siguiente {
									if !contains(siguientes_prod_actual, sig) {
										siguientes_prod_actual = append(siguientes_prod_actual, sig)
									}
								}
							}
						} else {
							siguientes_raiz_l := BuscarProduccion(nt, lista_siguientes)
							for _, sig := range siguientes_raiz_l {
								if !contains(siguientes_prod_actual, sig) {
									siguientes_prod_actual = append(siguientes_prod_actual, sig)
								}
							}
						}

					}
				}
			}
		}
		lista_siguientes = append(lista_siguientes, map[string][]string{nt_prod_actual: siguientes_prod_actual})
		indice_produccion_actual += 1
	}
	return lista_siguientes
}

func Primeros(gramatica []map[string][]string) []map[string][]string {
	lista_primeros := make([]map[string][]string, 0)

	no_terminales := ListaNoTerminales(gramatica)
	terminales := ListaTerminales(gramatica)

	for _, element := range gramatica {
		for key, value := range element {
			primeros_prod_actual := make([]string, 0)

			for _, i := range value {
				characters := strings.Split(i, " ")
				valor_actual := characters[0]
				prod_actual := make([]string, 0)

				if stringInSlice(characters[0], terminales) || characters[0] == "λ" {
					if !stringInSlice(characters[0], primeros_prod_actual) {
						primeros_prod_actual = append(primeros_prod_actual, characters[0])
					}
				} else {
					for stringInSlice(valor_actual, no_terminales) {
						prod_actual = BuscarProduccion(valor_actual, gramatica)
						for _, value := range prod_actual {
							arr_value := strings.Split(value, " ")
							if stringInSlice(arr_value[0], terminales) || arr_value[0] == "λ" {
								if !stringInSlice(arr_value[0], primeros_prod_actual) {
									primeros_prod_actual = append(primeros_prod_actual, arr_value[0])
								}
							}
							valor_actual = arr_value[0]
						}
					}
				}
			}
			lista_primeros = append(lista_primeros, map[string][]string{key: primeros_prod_actual})
		}
	}
	return lista_primeros
}
