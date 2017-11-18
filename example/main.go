package main

import (
	tbago "../../tbago"
	"fmt"
	"log"
)

func main() {
    tba, err := tbago.New("jVuUB9i97FgP8vVaCTCWFw5bjBmKyjG80nAs4nQhbS2G2xBVmIKvvE0lXnCeuciV")
	if err != nil {
		log.Fatal(err)
	}

	short, err := tba.Team(1418).Simple().Get()
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(short.Nickname)
}
