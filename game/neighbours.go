package game

func getNeighbours(x int, y int, elements [][]bool) int {
    var count int
    count += cellExists(x-1,y-1,elements)
    count += cellExists(x-1,y,elements)
    count += cellExists(x-1,y+1,elements)
    count += cellExists(x,y-1,elements)
    count += cellExists(x,y+1,elements)
    count += cellExists(x+1,y-1,elements)
    count += cellExists(x+1,y,elements)
    count += cellExists(x+1,y+1,elements)
    return count
}

func cellExists(x int, y int, elements [][]bool) int {
    if x >= len(elements) || y >= len(elements[:]) || x < 0 || y < 0 {
        return 0
    } else if elements[x][y] == false {
        return 0
    } else {
        return 1
    }
}
