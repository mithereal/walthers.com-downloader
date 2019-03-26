package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app      = kingpin.New("walthers", "Version: "+Version+" \n A command-line application to scrape sales from walthers.com.")
	images   = kingpin.Flag("images", "Fetch Images.").Short('i').Bool()
	username = kingpin.Arg("username", "Username.").Required().String()
	password = kingpin.Arg("password", "Password.").Required().String()
	jobtype  = kingpin.Arg("type", "Type.").Default("sales").String()
)
