package main

import (
	"fmt"
	"flag"
	"github.com/panoplyio/cwlogs/app"
)

var conf = &app.Config{}

func init() {
	flag.StringVar(&conf.Group, "group", "", "AWS CloudWatch group name")
	flag.StringVar(&conf.Stream, "stream", "", "AWS CloudWatch stream name")
	flag.StringVar(&conf.Region, "region", "us-east-1", "AWS CloudWatch region")
	flag.BoolVar(&conf.FromHead, "head", true, "AWS CloudWatch logs will be loaded from the head")
	flag.Parse()
}

func main() {
	downloader := app.NewLogsDownloader(conf)
	downloader.OnLoaded = func(total string) {
		fmt.Printf("Data loaded: %-100s\r", total)
	}

	err := downloader.Run()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Completed!")
}
