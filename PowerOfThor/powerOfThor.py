def compare(point, position):
    x = 1 if point[0]-position[0]>0 else (0 if point[0]-position[0]==0 else -1)
    y = 1 if point[1]-position[1]>0 else (0 if point[1]-position[1]==0 else -1)
    return x, y

# lx: the X position of the light of power
# ly: the Y position of the light of power
# tx: Thor's starting X position
# ty: Thor's starting Y position
lx, ly, tx, ty = [int(i) for i in input().split()]

directions = {
	(0,-1):"N", 
	(1,-1):"NE", 
	(1,0):"E",
	(1,1):"SE",
	(0,1):"S",
	(-1,1):"SW",
	(-1,0):"W",
	(-1,-1):"NW"
}

while True:
    move = compare((lx,ly),(tx,ty))
    print(directions[move])
    tx+=move[0]
    ty+=move[1]
