package main

import (
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func main() {
	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 512,
		Title:  "UnGallery Desktop Client",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Run()
}
