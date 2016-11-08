// Código de exemplo para o Capítulo 2.2 para "Build Web Application with Golang"
// Propósito: Compreender a atribuição e manipulação de tipos de dados básicos.
package main

import (
	"errors"
	"fmt"
)

// constantes
const Pi = 3.1415926

// booleanos por padrão são `false`
var isActive bool                   // variável global
var enabled, disabled = true, false // omitindo os tipos de variáveis

// definições agrupadas
const (
	i         = 1e4
	MaxThread = 10
	prefix    = "astaxie_"
)

var (
	frenchHello string      // forma básica de definir string
	emptyString string = "" // define uma string e atribui vazio
)

func show_multiple_assignments() {
	fmt.Println("show_multiple_assignments()")
	var v1 int = 42

	// Define três variáveis como "int", e inicializa seus valores.
	// vname1 é v1, vname2 é v2, vname3 é v3
	var v2, v3 int = 2, 3

	// `:=` só funciona em funções
	// `:=` a maneira curta de definir variáveis sem
	//  especificar o tipo usando a palavra-chave `var`.
	vname1, vname2, vname3 := v1, v2, v3

	// `_` ignora o valor atribuído.
	_, b := 34, 35

	fmt.Printf("vname1 = %v, vname2 = %v, vname3 = %v\n", vname1, vname2, vname3)
	fmt.Printf("v1 = %v, v2 = %v, v3 = %v\n", v1, v2, v3)
	fmt.Println("b =", b)
}

func show_bool() {
	fmt.Println("show_bool()")
	var available bool // variável local
	valid := false     // atribuição curta
	available = true   // atribuindo valor a variável

	fmt.Printf("valid = %v, !valid = %v\n", valid, !valid)
	fmt.Printf("available = %v\n", available)
}

func show_different_types() {
	fmt.Println("show_different_types()")
	var (
		unicodeChar rune
		a           int8
		b           int16
		c           int32
		d           int64
		e           byte
		f           uint8
		g           int16
		h           uint32
		i           uint64
	)
	var cmplx complex64 = 5 + 5i

	fmt.Println("Default values for int types")
	fmt.Println(unicodeChar, a, b, c, d, e, f, g, h, i)

	fmt.Printf("Value is: %v\n", cmplx)
}
func show_strings() {
	fmt.Println("show_strings()")
	no, yes, maybe := "no", "yes", "maybe" // inicialização rápida
	japaneseHello := "Ohaiyou"
	frenchHello = "Bonjour" // forma básica de atribuição

	fmt.Println("Random strings")
	fmt.Println(frenchHello, japaneseHello, no, yes, maybe)

	// O acento grave, `, não irá escapar nenhum caracter na string
	fmt.Println(`This
	is on
	multiple lines`)
}

func show_string_manipulation() {
	fmt.Println("show_string_manipulation()")
	var s string = "hello"

	//Não é possível fazer isso com strings
	//s[0] = 'c'

	s = "hello"
	c := []byte(s) // converte string para o tipo []byte
	c[0] = 'c'
	s2 := string(c) // converte de volta para string

	m := " world"
	a := s + m

	d := "c" + s[1:] // não é possível alterar o valor da string pelo índice, mas é possível obter os valores.
	fmt.Printf("%s\n", d)

	fmt.Printf("s = %s, c = %v\n", s, c)
	fmt.Printf("s2 = %s\n", s2)
	fmt.Printf("combined strings\na = %s, d = %s\n", a, d)
}

func show_errors() {
	fmt.Println("show_errors()")
	err := errors.New("Example error message\n")
	if err != nil {
		fmt.Print(err)
	}
}

func show_iota() {
	fmt.Println("show_iota()")
	const (
		x = iota // x == 0
		y = iota // y == 1
		z = iota // z == 2
		w        // Se não houver nenhuma declaração depois de uma constante,
		// é usada a última expressão, então aqui temos que w = iota implicitamente.
		// Portanto w == 3, e tanto y quanto x também poderiam omitit a parte "= iota" da mesma forma.
	)

	const v = iota // cada vez que iota encontra a palavra chave `const`,  ele é resetado para `0`, então v = 0.

	const (
		e, f, g = iota, iota, iota // e=0,f=0,g=0 valores de iota são iguais quando estão na mesma linha.
	)
	fmt.Printf("x = %v, y = %v, z = %v, w = %v\n", x, y, z, w)
	fmt.Printf("v = %v\n", v)
	fmt.Printf("e = %v, f = %v, g = %v\n", e, f, g)
}

// Função e variáveis capitalizadas são públicas para outros pacotes.
// Todo o restante é privado
func This_is_public()  {}
func this_is_private() {}

func set_default_values() {
	// valores padrões para os tipos.
	const (
		a int     = 0
		b int8    = 0
		c int32   = 0
		d int64   = 0
		e uint    = 0x0
		f rune    = 0   // rune é na verdade int32
		g byte    = 0x0 // byte é na verdade uint8
		h float32 = 0   // tamanho é 4 byte
		i float64 = 0   // tamanho é 8 byte
		j bool    = false
		k string  = ""
	)
}
func show_arrays() {
	fmt.Println("show_arrays()")
	var arr [10]int // an array de type int
	arr[0] = 42     // array é indexado em 0 (0-based)
	arr[1] = 13     // atribuíndo valor a um elemento

	a := [3]int{1, 2, 3} // define um array de int com 3 elementos

	b := [10]int{1, 2, 3}
	// define um array de int com 10 elementos,
	// onde os primeiros três são atribuídos, o restante sendo setados com o padrão 0.

	c := [...]int{4, 5, 6} // use `…` para substituir o tamanho, Go irá calcular pra você.

	// define um array de duas dimensões com 2 elementos, e cada um deles com 4 elementos.
	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}

	// Declaração curta.
	easyArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}

	fmt.Println("arr =", arr)
	fmt.Printf("The first element is %d\n", arr[0]) // Recupera o valor do elemento, retorna 42
	fmt.Printf("The last element is %d\n", arr[9])  //retornará o valor padrão, que é 0 no caso.

	fmt.Println("array a =", a)
	fmt.Println("array b =", b)
	fmt.Println("array c =", c)

	fmt.Println("array doubleArray =", doubleArray)
	fmt.Println("array easyArray =", easyArray)
}
func show_slices() {
	fmt.Println("show_slices()")
	// define um slice com 10 elementos do tipo byte
	var ar = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}

	// define dois slices do tipo []byte
	var a, b []byte

	// os pontos do array do terceiro ao sexto.
	a = ar[2:5]
	// agora a possui os elementos ar[2]、ar[3] and ar[4]

	// b é outro slice de ar
	b = ar[3:5]
	// agora b possui os elementos ar[3] e ar[4]

	// define um array
	var array = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	// define dois slices
	var aSlice, bSlice []byte

	// algumas operações importantes
	aSlice = array[:3] // mesmo que aSlice = array[0:3] aSlice tem a,b,c
	aSlice = array[5:] // mesmo que aSlice = array[5:10] aSlice tem f,g,h,i,j
	aSlice = array[:]  // mesmo que aSlice = array[0:10] aSlice tem todos os elementos do array

	// slice de slice
	aSlice = array[3:7]  // aSlice possui os elementos d,e,f,g，len=4，cap=7
	bSlice = aSlice[1:3] // bSlice contém aSlice[1], aSlice[2], então possui os elementos e,f
	bSlice = aSlice[:3]  // bSlice contém aSlice[0], aSlice[1], aSlice[2], então possui os elementos d,e,f
	bSlice = aSlice[0:5] // slice pode ser expandido até a faixa de cap, agora bSlice contém d,e,f,g,h
	bSlice = aSlice[:]   // bSlice possui os mesmos elementos de aSlice, que são d,e,f,g

	fmt.Println("slice ar =", ar)
	fmt.Println("slice a =", a)
	fmt.Println("slice b =", b)
	fmt.Println("array =", array)
	fmt.Println("slice aSlice =", aSlice)
	fmt.Println("slice bSlice =", bSlice)
	fmt.Println("len(bSlice) =", len(bSlice))
}
func show_map() {
	fmt.Println("show_map()")
	// use string como tipo da chave, int como tipo do valor, e use `make` para inicializar.
	var numbers map[string]int
	// another way to define map
	numbers = make(map[string]int)
	numbers["one"] = 1 // atribuíndo valor pela chave
	numbers["ten"] = 10
	numbers["three"] = 3

	// Inicializando um map
	rating := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2}

	fmt.Println("map numbers =", numbers)
	fmt.Println("The third number is: ", numbers["three"]) // get values
	// saída: The third number is: 3

	// map possui dois valores de retorno. O segundo valor retorna false quando a chave não existe，então ok é  false, caso contrário ok é true.
	csharpRating, ok := rating["C#"]
	if ok {
		fmt.Println("C# is in the map and its rating is ", csharpRating)
	} else {
		fmt.Println("We have no rating associated with C# in the map")
	}

	delete(rating, "C") // deleta o elemento com a chave "c"
	fmt.Printf("map rating = %#v\n", rating)
}
//executa todas as funções
func main() {
	show_multiple_assignments()
	show_bool()
	show_different_types()
	show_strings()
	show_string_manipulation()
	show_errors()
	show_iota()
	set_default_values()
	show_arrays()
	show_slices()
	show_map()
}
