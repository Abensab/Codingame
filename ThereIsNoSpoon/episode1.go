package main

import "fmt"
import "os"
import "bufio"
import "strconv"

func iToStr(i int) string {
	return strconv.Itoa(i)
}

func get_neighbours(matrix [][]int, vertex [2]int) (neighbours map[string]string) {

	neighbours = make(map[string]string)
	neighbours["right"] = "-1 -1"
	neighbours["down"] = "-1 -1"
	fmt.Fprintln(os.Stderr, "Vertex: "+iToStr(vertex[0])+" "+iToStr(vertex[1]))
	row := vertex[1]
	col := vertex[0]
	for j := col + 1; j < len(matrix[0]); j++ {
		if matrix[row][j] == 0 {
			neighbours["right"] = iToStr(j) + " " + iToStr(row)
			break
		}
	}
	for i := row + 1; i < len(matrix); i++ {
		if matrix[i][col] == 0 {
			neighbours["down"] = iToStr(col) + " " + iToStr(i)
			break
		}
	}
	return
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
	graph := make(map[[2]int]map[string]string)

	for i := 0; i < height; i++ {
		scanner.Scan()
		line := scanner.Text() // width characters, each either 0 or .
		for j := 0; j < width; j++ {
			fmt.Fprintln(os.Stderr, "I: "+iToStr(i)+" ,J: "+iToStr(j))
			if line[j] == '.' {
				matrix[i][j] = -1
			} else {
				graph[[2]int{j, i}] = make(map[string]string)
			}
		}
	}

	fmt.Fprintln(os.Stderr, matrix)
	for key, _ := range graph {
		graph[key] = get_neighbours(matrix, key)
	}
	fmt.Fprintln(os.Stderr, graph)
	for key, value := range graph {
		fmt.Println(iToStr(key[0]) + " " + iToStr(key[1]) + " " + value["right"] + " " + value["down"])
	}
}
