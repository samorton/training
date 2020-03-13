package main

import "log"
import "github.com/labstack/gommon/color"

func color() {

	log.Print("starting")

	color.Println(color.Black("black"))
	color.Println(color.Red("red"))
	color.Println(color.Green("green"))
	color.Println(color.Yellow("yellow"))
	color.Println(color.Blue("blue"))
	color.Println(color.Magenta("magenta"))
	color.Println(color.Cyan("cyan"))
	color.Println(color.White("white"))
	color.Println(color.Grey("grey"))

}
