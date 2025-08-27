package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/bogem/id3v2/v2"
	"github.com/mitchellh/go-homedir"
)

type App struct {
	cli *CLI
}

func New(cli *CLI) *App {
	return &App{cli: cli}
}

func (a *App) Run(ctx context.Context) error {
	path, err := homedir.Expand(a.cli.File)
	if err != nil {
		return err
	}
	
	tag, err := id3v2.Open(path, id3v2.Options{Parse: true})
	if err != nil {
		log.Fatal(err)
	}
	defer tag.Close()

	for _, frame := range tag.GetFrames("CHAP") {
		cf := frame.(id3v2.ChapterFrame)
		st := cf.StartTime
		st = st.Round(time.Second)
		h := st / time.Hour
		st -= h * time.Hour
		m := st / time.Minute
		st -= m * time.Minute
		s := st / time.Second

		fmt.Printf("%02d:%02d:%02d %s\n", h, m, s, cf.Title.Text)
	}

	return nil
}