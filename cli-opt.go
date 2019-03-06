package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	images   = kingpin.Flag("images", "Fetch Images.").Short('i').Bool()
	username = kingpin.Arg("username", "Username.").Required().String()
	password = kingpin.Arg("password", "Password.").Required().String()
)
