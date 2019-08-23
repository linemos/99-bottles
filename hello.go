package main

import (
    "fmt"
    "bottles/wall"
    "strings"
)

const numberOfBottles = 99

func sing() {
    var i = numberOfBottles
    for ; i >= 0; i-- {
        fmt.Print(verse(i))
    }
}

func verse(bottleNumber int) string {
    subject := fmt.Sprintf("%d bottles", bottleNumber)
    midSentence := "Take one down and pass it around"
    lastSentence := fmt.Sprintf("%d bottles of beer", bottleNumber - 1)

    if bottleNumber == 1 {
        subject = "bottle"
        lastSentence = "no more bottles"
    }

    if bottleNumber == 2 {
        lastSentence = fmt.Sprintf("%d bottle of beer", bottleNumber - 1)
    }

    if bottleNumber == 0 {
        subject = "No more bottles"
        midSentence = "Go to the store and buy some more"
        lastSentence = "99 bottles"

    }


    return fmt.Sprintf("%d %s of beer on the wall, %s of beer.\n%s, %s on the wall.\n\n",
        bottleNumber,
        subject,
        strings.ToLower(subject),
        midSentence,
        lastSentence)
}

func main() {
    sing()
    fmt.Print(wall.Hello())
}

