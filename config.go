package main

import (
	"net/http"

	"gitlab.inria.fr/pmorillo/gog5k"
)

// Config specifies gog5k connect configuration
type Config struct {
	Username      string
	Password      string
	RestfullyFile string
}

// Client returns a *gog5k.Client to interact with Grid5000 API
func (c *Config) Client() interface{} {
	var tp gog5k.BasicAuthTransport
	if c.RestfullyFile != "" {
		tp = gog5k.BasicAuthTransport{
			RestfullyFile: c.RestfullyFile,
		}
		return gog5k.NewClient(tp.RestfullyClient())
	}
	if (c.Username != "") && (c.Password != "") {
		tp = gog5k.BasicAuthTransport{
			Username: c.Username,
			Password: c.Password,
		}
		return gog5k.NewClient(tp.Client())
	}

	return gog5k.NewClient(&http.Client{})
}
