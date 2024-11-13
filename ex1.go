package main

import (
    "fmt"
    "math/rand"
    "os"
    "strconv"
    "time"
)

var numb int

func main() {
    var aSurvives int = 0 // Count of Player A's victories
    var bSurvives int = 0 // Count of Player B's victories
    var playerA bool
    var playerB bool
    var coin bool
    var bGoesFirst int = 0 // Count of times Player B goes first
    var aGoesFirst int = 0 // Count of times Player A goes first

    rand.Seed(time.Now().UnixNano())
    num1, err := strconv.Atoi(os.Args[1])
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    num2, err := strconv.Atoi(os.Args[2])
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    num3, err := strconv.Atoi(os.Args[3])
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    numb = num1

    file, err := os.Create("duel.txt")
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer file.Close()

    // The duel begins
    // Loop through the simulations
    for i := 0; i < num1; i++ {
        // Continue the duel until one player wins
        for {
            if coin {
                // Player A's turn
                playerA = rand.Intn(101) > num2
                if playerA {
                    // Player A wins the duel
                    aSurvives++
                    break
                }
                // Player A missed, switch to Player B's turn
                coin = false
            } else {
                // Player B's turn
                playerB = rand.Intn(101) > num3
                if playerB {
                    // Player B wins the duel
                    bSurvives++
                    break
                }
                // Player B missed, switch to Player A's turn
                coin = true
            }
        }
    }

    // Write results to file
    _, err = fmt.Fprintf(file, "Number of duels: %v \n", num1)
    _, err = fmt.Fprintf(file, "Player A survives: %v \n", aSurvives)
    _, err = fmt.Fprintf(file, "Player B survives: %v \n", bSurvives)
    _, err = fmt.Fprintf(file, "A survives in %v%% of the duels\n", aSurvives*100/num1)
    _, err = fmt.Fprintf(file, "B survives in %v%% of the duels\n", bSurvives*100/num1)
    _, err = fmt.Fprintf(file, "B was the first: %v \n", bGoesFirst)

    if err != nil {
        fmt.Println("Error writing to file:", err)
        return
    }

    fmt.Printf("Simulated %v duels\n", num1)
}