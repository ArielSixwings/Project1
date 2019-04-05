package main

import (
	"bufio"
	"container/heap"
	"container/list"
	"fmt"
	"os"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

const N = 20 //--> Algoritmo O(n!) portanto n não pode ser muito grande
const INF = 1000000007

var Distance [][]int64

//using ll = long long;

type pair struct {
	first  int64
	second int64
}

var Edge int                       //E
var Base uint                      //verticeInicial
var NumberVertex uint              //V
var graph = make([][]pair, N)      //vector<pair<int,int>> graph[N];
var currentedges = make([]pair, N) //vector<pair<int,int>> graph[N];

var BestPath = list.New()
var Path = list.New()
var found = list.New()

func Dijkstra(Initial int) {

	// Definimos como infinita a Distance de todo current ligado ao 'incial'
	var i uint
	for i = 0; i < NumberVertex; i++ {
		Distance[Initial][i] = INF

	}
	Line := &IntHeap{Initial}
	heap.Init(Line)
	//heap.Push(Line, Initial)
	Distance[Initial][Initial] = 0
	for Line.Len() > 0 { //While it is not empty
		var TopLine = (*Line)[0]
		heap.Pop(Line)
		for i := 0; i < len(graph[TopLine]); i++ {
			currentedges = graph[TopLine]
			if currentedges[i].first+Distance[Initial][TopLine] < Distance[Initial][currentedges[i].second] {
				Distance[Initial][currentedges[i].second] = Distance[Initial][TopLine] + currentedges[i].first

				heap.Push(Line, currentedges[i].second) // E botar na fila

			}
		}
	}

}
func caixeiroViajante(current uint, bitmask uint) int {
	found.PushBack(Base)
	if current == Base && bitmask == ((1<<NumberVertex)-1) {

		return 0
	}
	bitmask = bitmask + (1 << current)
	var UltraInf int64 = INF * INF
	var i uint
	for i = 0; i < uint(NumberVertex); i++ {
		var whatnigga bool = (i == current)
		bitmask = bitmask << shift(1)
		var compare = bitmask & 1
		var whatfuck bool = !(compare || i == uint(Base))
		if whatnigga && whatfuck {

			var auxLL int = 0
			auxLL = caixeiroViajante(uint(i), bitmask)
			var aux64LL int64 = int64(auxLL)
		}
	}
	Path.PushBack(current)

	return int(UltraInf) // Retornando par de vector caminho + valor do caminho
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input number of vertex and edges")
	_, err := fmt.Scanf("%d", &NumberVertex)
	_, err = fmt.Scanf("%d", &Edge)

	fmt.Println("Digite o numero dr vertices e arestas, depois para cada aresta\n")
	fmt.Println("a aresta de origem, destino e o peso da aresta\n")

	var dist, e, v int
	for i := 0; i < Edge; i++ {

		_, err = fmt.Scanf("%d", &e)
		_, err = fmt.Scanf("%d", &v)
		_, err = fmt.Scanf("%d", &dist)

		graph[e-1][v].first = int64(v)
		graph[e-1][v].second = int64(dist)

	}

	for i := 0; i < v; i++ {
		Dijkstra(i)
	}
	//Print bonito da tabela de distancias

	fmt.Printf("   ")
	for i := 0; i < v; i++ {
		fmt.Printf("%4d ", i+1)

	}
	fmt.Println("\n")

	fmt.Println("   ")
	for i := 0; i < v; i++ {
		fmt.Printf("____ ")

	}
	fmt.Println("\n")

	for i := 0; i < v; i++ {
		fmt.Printf("%d |", i+1)
		for j := 0; j < v; j++ {
			if Distance[i][j] == INF {
				fmt.Printf("%4d", -1)
			}
			fmt.Printf("%4d ", Distance[i][j])
		}
		fmt.Println("\n")
	}
	fmt.Printf("Formar caminho partindo e terminando em qual current: ")

	_, err = fmt.Scanf("%d", &Base)

	Base--
	var ll uint
	ll = uint(caixeiroViajante(Base, ll))
	// Chamamos a DP no current Initial com bitmask zerada

	//auto BestPath = caixeiroViajante(Base, 0LL);

	// Se o caminho deu maior que INF, esse caminho passa por arestas notadas como inexistente
	// Portanto não existe
	if ll >= INF {
		fmt.Println("There is no way back to the base\n")

	} else { // Mas caso exista, é retornado o par do UltraInf e caminho

		fmt.Printf("Tamanho do percurso do caminho mais curto = \n", ll)

		fmt.Println("Path: ")
		fmt.Printf("%d  ", (*BestPath))

		// for i := 0; i < BestPath.Len(); i++ {
		// }

		fmt.Println("\n")
	}

}
