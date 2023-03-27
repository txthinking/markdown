package main

import "embed"

//go:embed static/*
var static embed.FS
