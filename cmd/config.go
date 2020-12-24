package main

type config struct {
	Port string `envconfig:"PORT" default:"8080"`
}
