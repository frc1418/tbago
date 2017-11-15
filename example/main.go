package main

import (
	tbago "../../tbago"
	"fmt"
	"log"
)

func main() {
	token := "jVuUB9i97FgP8vVaCTCWFw5bjBmKyjG80nAs4nQhbS2G2xBVmIKvvE0lXnCeuciV"

    tba, err := tbago.New(token)
	if err != nil {
		log.Fatal(err)
	}

	short, err := tba.Team(1418).Simple().Get()
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(short.Nickname)
}
