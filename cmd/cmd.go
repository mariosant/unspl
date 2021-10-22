package cmd

import (
	"time"
	w "unspl/wallpaperUrl"

	"github.com/briandowns/spinner"
	"github.com/reujab/wallpaper"
	"github.com/urfave/cli/v2"
)

var randomCommand = &cli.Command{
	Name:  "random",
	Usage: "Set a random image as wallpaper",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "resolution",
			Value: "1920x1080",
			Usage: "resolution of wallaper",
		},
		&cli.StringSliceFlag{
			Name:  "keyword",
			Usage: "search keyword",
		},
	},
	Action: func(c *cli.Context) error {
		resolution := c.String("resolution")
		keywords := c.StringSlice("keyword")

		url := w.GetRandomWallpaperURL(resolution, keywords)

		err := wallpaper.SetFromURL(url)

		return err
	},
}

var dailyCommand = &cli.Command{
	Name:  "daily",
	Usage: "Set the daily pick as wallpaper",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "resolution",
			Value: "1920x1080",
			Usage: "resolution of wallaper",
		},
		&cli.StringSliceFlag{
			Name:  "keyword",
			Usage: "search keyword",
		},
	},
	Action: func(c *cli.Context) error {
		resolution := c.String("resolution")
		keywords := c.StringSlice("keyword")

		url := w.GetDailyWallpaperURL(resolution, keywords)

		err := wallpaper.SetFromURL(url)

		return err
	},
}

var weeklyCommand = &cli.Command{
	Name:  "weekly",
	Usage: "Set the weekly pick as wallpaper",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "resolution",
			Value: "1920x1080",
			Usage: "resolution of wallaper",
		},
		&cli.StringSliceFlag{
			Name:  "keyword",
			Usage: "search keyword",
		},
	},
	Action: func(c *cli.Context) error {
		resolution := c.String("resolution")
		keywords := c.StringSlice("keyword")

		url := w.GetWeeklyWallpaperURL(resolution, keywords)

		err := wallpaper.SetFromURL(url)

		return err
	},
}

func defaultActionHandler(c *cli.Context) error {
	resolution := c.String("resolution")
	keywords := c.StringSlice("keyword")

	url := w.GetRandomWallpaperURL(resolution, keywords)

	err := wallpaper.SetFromURL(url)

	return err
}

func beforeAction(c *cli.Context) error {
	s.Suffix = " Fetching a new image from Unsplash..."
	s.Start()

	return nil
}

func afterAction(c *cli.Context) error {
	if err := c.Err(); err == nil {
		s.FinalMSG = "Done!"
	}

	s.Stop()

	return nil
}

var s = spinner.New(spinner.CharSets[41], 100*time.Millisecond)

func App() *cli.App {
	app := &cli.App{
		Name:  "Unspl",
		Usage: "Switch to Unsplash wallpapers",

		Commands: []*cli.Command{
			randomCommand,
			dailyCommand,
			weeklyCommand,
		},

		Action: defaultActionHandler,
		Before: beforeAction,
		After:  afterAction,
	}

	return app
}
