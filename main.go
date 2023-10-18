package main

import (
	"fmt"
	"game-of-life/game"
	"os"
	"os/exec"
	"time"
)

func main() {
    clear()
    var res int
    for {
        fmt.Print("res: ")
        _, e := fmt.Scanf("%d",&res)
        if e == nil {
            break
        }
    }
    var grid = game.GenGame(res,res)
    var errorReading int

    for {
        clear()
        grid.Print()
        fmt.Println()
        fmt.Println("(empty): step")
        fmt.Println("1: insert")
        fmt.Println("2: remove")
        fmt.Println("3: rand")
        fmt.Println("4: full")
        fmt.Println("5: empty")
        fmt.Println("6: auto")
        fmt.Println("7: exit")
        fmt.Println()

        if errorReading > 0 {
            fmt.Printf("not an option (%dx)\n", errorReading)
        }

        fmt.Print(">> ")

        var input string
        fmt.Scanln(&input)

        switch input {
        case "":
            clear()
            grid.Step()

        case "1":
            grid.Insert()

        case "2":
            grid.Remove()

        case "3":
            clear()
            grid.Rand()

        case "4":
            clear()
            grid.Full()

        case "5":
            clear()
            grid.Empty()

        case "6":

            var delay int
            fmt.Print("Delay(ms): ")
            fmt.Scanf("%d", &delay)
            for {
                clear()
                grid.Step()
                grid.Print()
                time.Sleep(time.Duration(delay) * time.Millisecond)
            }

        case "7":
            goto exit

        default:
            errorReading+=1
            continue
        }
        errorReading = 0
    }
    exit:
}

func clear() {
    c := exec.Command("clear")
    c.Stdout = os.Stdout
    c.Run()
}
