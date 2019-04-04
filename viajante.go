package main

import (
	"container/heap"
	"container/list"
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

var Distance [][]int

//using ll = long long;

type pair struct {
	first  float32
	second int
}

var NumberVertex int               //V
var graph = make([][]pair, N)      //vector<pair<int,int>> grafo[N];
var currentedges = make([]pair, N) //vector<pair<int,int>> grafo[N];

func Dijkstra(Initial int) {

	// Definimos como infinita a Distance de todo current ligado ao 'incial'

	for i := 0; i < NumberVertex; i++ {
		Distance[Initial][i] = INF

	}
	Line := &IntHeap{Initial}
	heap.Init(Line)
	//heap.Push(Line, Initial)
	Distance[Initial][Initial] = 0
	for h.Len() > 0 { //While it is not empty
		var TopLine = (*Line)[0]
		heap.Pop(Line)
		for i := 0; i < len(graph[TopLine]); i++ {
			currentedges = graph[TopLine]
			if currentedges[i].first+Distance[Initial][TopLine] < Distance[Initial][currentedges[i].second] {
				Distance[Initial][v.second] = Distance[Initial][TopLine] + currentedges[i].first

				heap.push(Line, currentedges[i].second) // E botar na fila

			}
		}
	}

}
func caixeiroViajante(int current, ll bitmask) (int64, []int32) {
	if current == Base && bitmask == ((1<<NumberVertex)-1) {
		found := list.New()
		found.PushBack(Base)
		return 0, found
	} else if current == Base && (bitmask >> Base & 1) {
		var invalid []int
		return INF * INF, invalid
	}
	bitmask = bitmask + (1 << current)
	var UltraInf int64 = INF * INF
	var Path []int32

	for i := 0; i < NumberVertex; i++ {
		if i != current && (!((bitmask >> i) & 1) || i == Base) {

			var auxLL int64
			var result []int32 = caixeiroViajante(i, bitmask)
			if result.first+Distance[current][i] < UltraInf {
				UltraInf = resultado.first + Distance[current][i]
				Path = resultado.second
			}
		}
	}
	Path.push_back(vertice)

	return UltraInf, Path // Retornando par de vetor caminho + valor do caminho
}
func main() {
	var Edge int //E
	var Base int //verticeInicial;

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input number of vertex and edges")
	NumberVertex, _ := reader.ReadInt('\n')
	Edge, _ = reader.ReadInt('\n')

	fmt.Println("Digite o numero dr vertices e arestas, depois para cada aresta\n")
	fmt.Println("a aresta de origem, destino e o peso da aresta\n")

	for i := 0; i < Edge; i++ {

		var dist, e, v int
		e, _ = reader.ReadInt()
		v, _ = reader.ReadInt()
		dist, _ = reader.Readfloat32('\n')

		grafo[e-1][v].first = v
		grafo[e-1][v].second = dist

	}

	for i := 0; i < V; i++ {
		Dijkstra(i)
	}
	//Print bonito da tabela de distancias

	fmt.Printf("   ")
	for i := 0; i < V; i++ {
		fmt.Printf("%4d ", i+1)

	}
	fmt.Println("\n")

	printf("   ")
	for i := 0; i < V; i++ {
		fmt.Printf("____ ")

	}
	fmt.Println("\n")

	for i := 0; i < V; i++ {
		fmt.Printf("%d |", i+1)
		for j := 0; j < V; j++ {
			if Distance[i][j] == INF {
				fmt.Printf("%4d", -1)
			}
			fmt.Printf("%4d ", Distance[i][j])
		}
		fmt.Println("\n")
	}
	fmt.Printf("Formar caminho partindo e terminando em qual current: ")

	Base, _ = reader.ReadInt('\n')

	Base--
	var ll bitmask
	var BestPaht []int = caixeiroViajante(Base, ll)
	// Chamamos a DP no current Initial com bitmask zerada

	//auto menorCaminho = caixeiroViajante(Base, 0LL);

	// Se o caminho deu maior que INF, esse caminho passa por arestas notadas como inexistente
	// Portanto não existe
	if menorCaminho.first >= INF {
		fmt.Println("There is no way back to the base\n")

	} else { // Mas caso exista, é retornado o par do UltraInf e caminho

		fmt.Printf("Tamanho do percurso do caminho mais curto = \n", menorCaminho.first)

		vetor = menorCaminho.second
		reverse(vetor.begin(), vetor.end())

		fmt.Println("Caminho: ")

		for i := 0; i < len(vector); i++ {
			fmt.Printf("%d  ", vetor[i])
		}

		fmt.Println("\n")
	}

}
