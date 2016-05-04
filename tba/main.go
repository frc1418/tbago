package main

import (
	"fmt"
	"github.com/carlcolglazier/tba"
	"github.com/codegangsta/cli"
	"log"
	"os"
	"strconv"
)

const (
	VERSION = "0.0.0"
)

func main() {
	app := cli.NewApp()
	tba, err := tba.Init("carlcolglazier", "tba-cli", VERSION)
	if err != nil {
		log.Fatal(err)
	}
	app.Name = "tba"
	app.Usage = "a command line interface for The Blue Alliance"
	app.Version = VERSION
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		{
			Name:        "team",
			Aliases:     []string{"t"},
			Usage:       "tba team ####",
			Description: "find general information about a team",
			Action: func(c *cli.Context) {
				if !c.Args().Present() {
					fmt.Println("Usage: " + c.Command.Usage)
					return
				}
				teamNumberString := c.Args().First()
				teamNumber, err := strconv.Atoi(c.Args().First())
				if err != nil {
					fmt.Println("Invalid team number")
					return
				}
				team, err := tba.GetTeam(teamNumber)
				if err != nil {
					fmt.Println("Could not get infromation for " + teamNumberString)
					return
				}
				fmt.Printf("Team %-d - %-40s\n"+
					"  \"%s\"\n"+
					"  Location:    %40s\n"+
					"  Website:     %40s\n"+
					"  Rookie Year: %40d\n"+
					"  Full Name:   %-40s\n",
					team.TeamNumber, team.Nickname,
					team.Motto,
					team.Location,
					team.Website,
					team.RookieYear,
					team.Name)
			},
		},
		{
			Name:        "match",
			Aliases:     []string{"m"},
			Description: "find results or information about a match",
			Action: func(c *cli.Context) {
				if !c.Args().Present() {
					fmt.Println("Usage: " + c.Command.Usage)
					return
				}
				matchKey := c.Args().First()
				match, err := tba.GetMatch(matchKey)
				if err != nil {
					fmt.Println("Could not get infromation for " + matchKey)
					return
				}
				fmt.Printf("%s - %s\n", match.Key, match.TimeString)
				fmt.Println("match")
			},
		},
	}

	app.Run(os.Args)
}
