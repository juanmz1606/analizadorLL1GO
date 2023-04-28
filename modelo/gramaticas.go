package modelo

var Gramatica1 = []map[string][]string{
	{"E": {"E + T", "T"}},
	{"T": {"T * F", "F"}},
	{"F": {"id", "( E )"}},
}

var Gramatica2 = []map[string][]string{
	{"S": {"S xx", "A B C D"}},
	{"A": {"p", "λ", "B D", "A p B"}},
	{"B": {"q C H", "q B H", "λ"}},
	{"H": {"xyxx"}},
	{"D": {"d", "λ"}},
	{"C": {"idd S fx", "id"}},
}

var Gramatica3 = []map[string][]string{
	{"B": {"D L"}},
	{"D": {"id; D", "λ"}},
	{"L": {"S ; L", "λ"}},
	{"S": {"a+a"}},
}

var Gramatica4 = []map[string][]string{
	{"S": {"Aa", "Bb"}},
	{"A": {"ab", "λ"}},
	{"B": {"ba", "λ"}},
}
