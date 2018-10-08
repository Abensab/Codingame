package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Graph struct {
	nodes map[[2]int]*Node
	edges map[Node][]*Node
}

type Node struct {
	size int
	x    int
	y    int
}

func (n *Node) getCoordinates() [2]int {
	return [2]int{n.x, n.y}
}

func (n *Node) String() string {
	return fmt.Sprintf("{Size: %v, [%v, %v] } ", n.size, n.x, n.y)
}

func (g *Graph) AddNode(n *Node) {
	if g.nodes == nil {
		g.nodes = make(map[[2]int]*Node)
	}
	g.nodes[[2]int{n.x, n.y}] = n
}

func (g *Graph) RemoveNode(n *Node) {
	//near := g.edges[*g.nodes[n.getCoordinates()]]
	fmt.Fprintln(os.Stderr, "Deleting node:", n)
	if _, ok := g.nodes[n.getCoordinates()]; ok {
		delete(g.nodes, n.getCoordinates())
	}
	/*for _, e := range near {
		g.RemoveEdgeFromNode(e)
	}*/
}

func (g *Graph) RemoveEdgeFromNode(e *Node) {
	for i, neig := range g.edges[*e] {
		if neig == e {
			g.edges[*e][i] = g.edges[*e][len(g.edges[*e])-1]
			g.edges[*e] = g.edges[*e][:len(g.edges[*e])-2]
		}
	}
	delete(g.edges, *e)
}

func (g *Graph) AddEdge(n1, n2 *Node) {
	if g.edges == nil {
		g.edges = make(map[Node][]*Node)
	}
	g.edges[*n1] = append(g.edges[*n1], n2)
	g.edges[*n2] = append(g.edges[*n2], n1)
}

func (g *Graph) getNode(coordinates [2]int) *Node {
	if n, ok := g.nodes[coordinates]; ok {
		return n
	}
	return nil
}

func (g *Graph) getNodesBySize(size int) []*Node {
	var nodes []*Node
	for _, node := range g.nodes {
		if node.size == size {
			nodes = append(nodes, node)
		}
	}
	return nodes
}

func (g *Graph) decrementSize(node *Node, i int) {
	g.getNode([2]int{node.x, node.y}).size -= i
}

func (g *Graph) String() string {
	s := ""
	for k, _ := range g.nodes {
		s += g.nodes[k].String() + " -> "
		near := g.edges[*g.nodes[k]]
		for _, e := range near {
			s += e.String() + " "
		}
		s += "\n"
	}
	return s
}

func (g *Graph) getActiveNodes() []*Node {
	var activeNodes []*Node
	for _, node := range g.nodes {
		if node.size > 0 {
			activeNodes = append(activeNodes, node)
		}
	}
	return activeNodes

}

func (g *Graph) getActiveNeighbours(n *Node) []*Node {
	var neighbours []*Node
	for _, node := range g.edges[*g.nodes[n.getCoordinates()]] {
		if node.size > 0 {
			neighbours = append(neighbours, node)
		}
	}
	return neighbours
}

func (g *Graph) makeConnection(node *Node, size int, near []*Node) {
	for _, neighbour := range near {
		fmt.Fprintln(os.Stderr, "Node:", node, "Neighbour:", neighbour)
		delete := false
		if node.size > 0 && neighbour.size > 0 {
			fmt.Fprintln(os.Stderr, iToStr(node.x), iToStr(node.y), iToStr(neighbour.x), iToStr(neighbour.y), size)
			fmt.Println(iToStr(node.x), iToStr(node.y), iToStr(neighbour.x), iToStr(neighbour.y), size)
			g.decrementSize(node, size)
			g.decrementSize(neighbour, size)
			if node.size == 0 {
				//g.RemoveNode(node)
				delete = true
			}
			if neighbour.size == 0 {
				//g.RemoveNode(neighbour)
				delete = true
			}
		}
		if delete {
			solveGraph(g)
			break
		}
	}
}

/*RULES:
8->4 neighbours: 2 for each
6->3 neighbours: 2 for each
4->2 neighbours: 2 for each
3->2 neighbours: 2 to max and 1 to other
2->2 neighbours: 1 for each
2->1 neighbour: 2 to one neighbour
1->1 neighbour: 1 to neighbour
*/

func solveGraph(g *Graph) {
	fmt.Fprintln(os.Stderr, "Grafo:\n", g)
	for _, node := range g.getActiveNodes() {
		near := g.getActiveNeighbours(node)
		fmt.Fprintln(os.Stderr, node)
		fmt.Fprintln(os.Stderr, near)
		if len(near)*2 == node.size {
			g.makeConnection(node, 2, near)
		} else if node.size == 1 && len(near) == 1 {
			g.makeConnection(node, 1, near)
		} else if node.size == 2 && len(near) == 2 {
			g.makeConnection(node, 1, near[:len(near)-2])
			g.makeConnection(node, 1, near[len(near)-2:])

		} else if node.size == 3 && len(near) == 2 {
			near1 := near[0]
			near2 := near[1]
			if near1.size+near2.size == 3 {
				if near1.size == 1 {
					g.makeConnection(node, 1, near[:len(near)-1])
					g.makeConnection(node, 2, near[len(near)-2:])
				} else {
					g.makeConnection(node, 2, near[:len(near)-1])
					g.makeConnection(node, 1, near[len(near)-2:])
				}
			}
		} else if len(near) == node.size {
			g.makeConnection(node, 1, near)
		}
	}

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
				g.AddEdge(node, g.getNode([2]int{j, row}))
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
	solveGraph(&graph)

}
