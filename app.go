package main

import "github.com/urfave/cli/v2"

func CreateApp() *cli.App {
	return &cli.App{
		Name:  "convert",
		Usage: "converts any webp to png",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "out",
				Value: "./",
				Usage: "tells me where to write the pngs, default value is current directory",
			},
			&cli.BoolFlag{
				Name:    "delete",
				Value:   false,
				Aliases: []string{"D"},
				Usage:   "tells me to delete the webp file",
			},
		},
		Action: func(context *cli.Context) error {
			out := context.String("out")
			in := []string{"./"}
			if context.Args().Len() > 0 {
				in = context.Args().Slice()
			}
			ConvertImage(in, out, context.Bool("delete"))
			return nil
		},
		Commands: []*cli.Command{
			{
				Name: "populate",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:  "n",
						Usage: "specify the count of the populated",
						Value: 10,
					},
				},
				Action: func(context *cli.Context) error {
					Populate(context.Int("n"))
					return nil
				},
			},
		},
	}
}
