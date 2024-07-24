package main

import (
	"astroproject/astro/services"
	"embed"
	"log"

	"github.com/wailsapp/wails/v3/pkg/application"
)

var assets embed.FS

func main() {

	session := services.Session{}
	competitionManager := services.CompetitionManager{UserSession: &session}
	stageManager := services.StageManager{UserSession: &session}
	personManager := services.PersonManager{UserSession: &session}

	app := application.New(application.Options{
		Name:        "AstroProject",
		Description: "A demo of using raw HTML & CSS",
		Services: []application.Service{
			application.NewService(&session),
			application.NewService(&competitionManager),
			application.NewService(&stageManager),
			application.NewService(&personManager),
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
