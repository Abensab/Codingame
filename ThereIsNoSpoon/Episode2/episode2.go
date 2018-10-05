package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Graph struct {
	nodes []*Node
	edges map[Node][]*Node
}

type Node struct {
	size int
	x    int
	y    int
}

func (n *Node) String() string {
	return fmt.Sprintf("[%v, %v] : %v", n.x, n.y, n.size)
}

func (g *Graph) AddNode(n *Node) {
	g.nodes = append(g.nodes, n)
}

func (g *Graph) AddEdge(n1, n2 *Node) {
	if g.edges == nil {
		g.edges = make(map[Node][]*Node)
	}
	g.edges[*n1] = append(g.edges[*n1], n2)
	g.edges[*n2] = append(g.edges[*n2], n1)
}

func (g *Graph) getNode(coordinates [2]int) *Node {
	for _, node := range g.nodes {
		if node.x == coordinates[0] && node.y == coordinates[1] {
			return node
		}
	}
	return nil
}

func (g *Graph) String() string {
	s := ""
	for i := 0; i < len(g.nodes); i++ {
		s += g.nodes[i].String() + " -> "
		near := g.edges[*g.nodes[i]]
		for j := 0; j < len(near); j++ {
			s += near[j].String() + " "
		}
		s += "\n"
	}
	return s
}

func iToStr(i int) string {
	return strconv.Itoa(i)
}

func sToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func (g *Graph) set_neighbours(matrix [][]int) {

	for _, node := range g.nodes {
		row := node.y
		col := node.x
		for j := col + 1; j < len(matrix[0]); j++ {
			if matrix[row][j] > 0 {
				g.AddEdge(node, g.getNode([2]int{row, j}))
				break
			}
		}
		for i := row + 1; i < len(matrix); i++ {
			if matrix[i][col] > 0 {
				g.AddEdge(node, g.getNode([2]int{col, i}))
				break
			}
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)
	// width: the number of cells on the X axis
	var width int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &width)
	// height: the number of cells on the Y axis
	var height int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &height)

	fmt.Fprintln(os.Stderr, "Width: "+iToStr(width)+" ,Height: "+iToStr(height))

	matrix := make([][]int, 0)
	for i := 0; i < height; i++ {
		tmp := make([]int, 0)
		for j := 0; j < width; j++ {
			tmp = append(tmp, 0)
		}
		matrix = append(matrix, tmp)
	}
	var graph Graph

	for i := 0; i < height; i++ {
		scanner.Scan()
		line := scanner.Text() // width characters, each either 0 or .
		for j := 0; j < width; j++ {
			fmt.Fprintln(os.Stderr, "I: "+iToStr(i)+" ,J: "+iToStr(j))
			if string(line[j]) == "." {
				matrix[i][j] = -1
			} else {
				size, _ := sToInt(string(line[j]))
				n := Node{size, j, i}
				graph.AddNode(&n)
				matrix[i][j] = size
			}
		}
	}
	graph.set_neighbours(matrix)
	fmt.Fprintln(os.Stderr, "Matrix: ", matrix)
	fmt.Fprintln(os.Stderr, graph.String())
	//fmt.Println(iToStr(key[0]) + " " + iToStr(key[1]) + " " + value["right"] + " " + value["down"])

}
