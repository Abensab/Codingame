import sys
import math

def get_neighbours(matrix, vertex):
    neighbours = {"right": "-1 -1", "down":"-1 -1"}
    print("Vertex: ",str(vertex[0]), str(vertex[1]), file=sys.stderr)
    row=vertex[0]
    col=vertex[1]
    for j in range(col+1, len(matrix[0])): #Right Neighbours
        if matrix[row][j]==0:
            neighbours["right"]=str(j)+" "+str(row)
            break
        
    for i in range(row+1,len(matrix)):
        if matrix[i][col]==0:
            neighbours["down"]=str(col)+" "+str(i)
            break
        
    return neighbours
    

if __name__=='__main__':

	width = int(input())  # the number of cells on the X axis
	height = int(input())  # the number of cells on the Y axis

	graph = {}
	matrix = [0]*height
	for i in range(height):
		matrix[i] = [0]*width

	for i in range(height):
		line = input()  # width characters, each either 0 or .
		for j in range(len(line)):
			print("I:"+str(i)+" J: "+str(j), file=sys.stderr)
			if line[j]=='.':
				matrix[i][j] = -1
			else:
				graph[(j,i)]={"right": "-1 -1", "down":"-1 -1"}

	print("Matrix after setting:", matrix, file=sys.stderr)
	print("Graph before setting:",graph, file=sys.stderr)

	for k in graph:
			graph[k] = get_neighbours(matrix,(k[1], k[0]))
		
	print("Graph after setting:",graph, file=sys.stderr)

	for k, v in graph.items():
		print(str(k[0])+" "+str(k[1])+" " + v["right"] +" "+ v["down"])
		
