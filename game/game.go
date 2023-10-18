package game

import (
	"fmt"
	"math/rand"
)

type grid struct {
    elements [][]bool
}

func GenGame(x int, y int) *grid {
    var randElements [][]bool = make([][]bool, x)

    for i := range randElements {
        randElements[i] = make([]bool, y)
    }

    return &grid{elements: randElements}
}

func (g *grid) Step() {
    var newElements [][]bool = make([][]bool, len(g.elements))

    for i := range newElements {
        newElements[i] = make([]bool, len(g.elements[i]))
    }

    for i := range g.elements {
        for j := range g.elements[i] {
            neighbours := getNeighbours(i,j,g.elements[:][:])
            switch {
            case g.elements[i][j] && neighbours < 2:
                newElements[i][j] = false

            case g.elements[i][j] && neighbours > 3:
                newElements[i][j] = false

            case g.elements[i][j] && (neighbours == 2 || neighbours == 3):
                newElements[i][j] = g.elements[i][j]
                
            case !g.elements[i][j] && neighbours == 3:
                newElements[i][j] = true
            }
        }
    }

    g.elements = newElements
}

func (g *grid) Insert() {
    var x,y int
    for {
        fmt.Print("X Y: ")
        n,e := fmt.Scanf("%d %d",&x, &y)

        if e != nil {
            continue
        }

        if x >= len(g.elements) || y >= len(g.elements) || x < 0 || y < 0{
            continue
        }

        if n == 2 {
            break
        }

    }
    g.elements[x][y] = true
}

func (g *grid) Remove() {
    var x,y int
    for {
        fmt.Print("X Y: ")
        n,e := fmt.Scanf("%d %d",&x, &y)

        if e != nil {
            continue
        }

        if x >= len(g.elements) || y >= len(g.elements) || x < 0 || y < 0{
            continue
        }

        if n == 2 {
            break
        }

    }
    g.elements[x][y] = false
}

func (g *grid) Rand() {
    for i := range g.elements {
        for j := range g.elements[i] {
            randInt := rand.Intn(2)
            if randInt == 0 {
                g.elements[i][j] = false
            } else {
                g.elements[i][j] = true
            }
        }
    }
}

func (g *grid) Full() {
    for i := range g.elements {
        for j := range g.elements[i] {
            g.elements[i][j] = true
        }
    }

}


func (g *grid) Empty() {
    for i := range g.elements {
        for j := range g.elements[i] {
            g.elements[i][j] = false
        }
    }

}

func (g *grid) Print() {
    for i := 0; i < len(g.elements); i++  {
        for j := 0; j < len(g.elements[i]); j++  {
            if g.elements[i][j] == true {
                fmt.Printf("%v", "@" )
            } else {
                fmt.Printf("%v", "." )
            }
        }
        fmt.Printf("\n")
    }
}
