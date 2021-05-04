//
// Author: Marcelo Gomes Jr <marcelo.gomes.junior@gmail.com>
// Created: mai/2021
//
// wbrowser
//

package main

type Config struct {
	Browsers  map[Browser]BrowserPath `json:"browsers"`
	Redirects []Domain                `json:"redirects"`
	Domains   map[Domain]Browser      `json:"domains"`
}

type Browser string

type BrowserPath string

type Domain string

const DefaultBrowser = "default"
