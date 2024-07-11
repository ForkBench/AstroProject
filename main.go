package main

import (
	"embed"
	"log"

	"github.com/wailsapp/wails/v3/pkg/application"

	"changeme/astro/services"
)

var assets embed.FS

func main() {

	session := services.Session{}

	app := application.New(application.Options{
		Name:        "AstroProject",
		Description: "A demo of using raw HTML & CSS",
		Services: []application.Service{
			application.NewService(&session),
			application.NewService(&services.Competition{}),
			application.NewService(&services.Club{}),
			application.NewService(&services.Pool{}),
			application.NewService(&services.Player{}),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})

	app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		Title: "AstroProject",
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTranslucent,
			TitleBar:                application.MacTitleBarHiddenInset,
		},
		BackgroundColour: application.NewRGB(27, 38, 54),
		URL:              "/",
		Width:            1080,
		Height:           720,
		Centered:         false,
	})

	err := app.Run()

	if err != nil {
		log.Fatal(err)
	}
}
