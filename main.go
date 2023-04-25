package main

import "fmt"
import "strings"
import "unicode"

func main() {
	fmt.Printf(strings.Repeat("_", 30)+"\n")
	fmt.Printf("GRAMATICA ORIGINAL"+"\n")
	imprimirGramatica(gramatica2)

	var gramatica2Lista = eliminarRecursion(gramatica2)
	fmt.Printf(strings.Repeat("_", 30)+"\n")
	fmt.Printf("ELIMINAR RECURSION"+"\n")
	imprimirGramatica(gramatica2Lista)

	fmt.Printf(strings.Repeat("_", 30)+"\n")
	fmt.Printf("CONJUNTO DE PRIMEROS"+"\n")
	imprimirGramatica(primeros(gramatica2Lista))

	fmt.Printf(strings.Repeat("_", 30)+"\n")
	fmt.Printf("CONJUNTO DE SIGUIENTES"+"\n")
	imprimirGramatica(siguientes(gramatica2Lista))
}

var gramatica1 = []map[string][]string{
    {"E": {"E + T", "T"}},
    {"T": {"T * F", "F"}},
    {"F": {"id", "( E )"}},
}

var gramatica2 = []map[string][]string{
    {"S": {"S xx", "A B C D"}},
    {"A": {"p", "λ", "B D", "A p B"}},
    {"B": {"q C H", "q B H", "λ"}},
    {"H": {"xyxx"}},
    {"D": {"d", "λ"}},
    {"C": {"idd S fx", "id"}},
}

var gramatica3 = []map[string][]string{
    {"B": {"D L"}},
    {"D": {"id; D", "λ"}},
    {"L": {"S ; L", "λ"}},
    {"S": {"a+a"}},
}

var gramatica4 = []map[string][]string{
    {"S": {"Q A"}},
    {"A": {"or Q A", "λ"}},
    {"Q": {"R B"}},
    {"B": {"R B", "λ"}},
    {"R": {"U", "x", "y"}},
    {"U": {"z"}},
}

func imprimirGramatica(gramatica []map[string][]string) {
    for _, element := range gramatica {
        for key, value := range element {
            fmt.Printf("%v -> [%v]\n", key, strings.Join(value, ","))
        }
    }
}

func eliminarRecursion(gramatica []map[string][]string) []map[string][]string {
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

func primeros(gramatica []map[string][]string) []map[string][]string {
    lista_primeros := make([]map[string][]string, 0)

    no_terminales := listaNoTerminales(gramatica)
    terminales := listaTerminales(gramatica)

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
                        prod_actual = buscarProduccion(valor_actual, gramatica)
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

func siguientes(gramatica []map[string][]string) []map[string][]string {
    lista_primeros := primeros(gramatica)
    lista_siguientes := []map[string][]string{}

    indice_produccion_actual := 0

    for indice_produccion_actual < len(gramatica) {
        nt_prod_actual := listKeys(gramatica[indice_produccion_actual])[0]
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
                                lista_primeros_del_siguiente := buscarProduccion(arr_derivado[indice_nt_actual+1], lista_primeros)
                                if contains(lista_primeros_del_siguiente, "λ") {
                                    lista_primeros_del_siguiente = removeElement(lista_primeros_del_siguiente, "λ")
                                    siguientes_raiz := buscarProduccion(nt, lista_siguientes)
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
                            siguientes_raiz_l := buscarProduccion(nt, lista_siguientes)
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

func listKeys(m map[string][]string) []string {
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

func removeElement(slice []string, element string) []string {
    result := make([]string, 0, len(slice))
    for _, s := range slice {
        if s != element {
            result = append(result, s)
        }
    }
    return result
}

func listaNoTerminales(gramatica []map[string][]string) []string {
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

func listaTerminales(gramatica []map[string][]string) []string {
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

func buscarProduccion(nombre_produccion string, gramatica []map[string][]string) []string {
    elementos_produccion := make([]string, 0)
    for _, prod := range gramatica {
        if key, ok := prod[nombre_produccion]; ok {
            elementos_produccion = append(elementos_produccion, key...)
        }
    }
    return elementos_produccion
}

func stringInSlice(s string, slice []string) bool {
    for _, value := range slice {
        if value == s {
            return true
        }
    }
    return false
}

