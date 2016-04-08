package main

//import "./cmd"

import (
	"../../tba"
	"fmt"
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
	tba, err := tba.Init("CarlColglazier", "tba-cli", VERSION)
	if err != nil {
		log.Fatal(err)
	}
	app.Name = "tba"
	app.Usage = "a command line interface for The Blue Alliance"
	app.Commands = []cli.Command{
		{
			Name:    "team",
			Aliases: []string{"t"},
			Usage:   "find general information about a team",
			Action: func(c *cli.Context) {
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
	}

	app.Run(os.Args)
}
