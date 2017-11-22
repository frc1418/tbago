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

	awards, err := tba.Team(1418).Event("2017vahay").Awards().Get()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(awards)

	matches, err := tba.Event("2017vahay").Matches().Team(1418).Simple().Get()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(matches)

	short, err := tba.Team(1418).Simple().Get()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(short.Nickname)

	vids, err := tba.Team(254).Media().Tag("chairmans_video").Year(2004).Get()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(vids)
	*/
}
