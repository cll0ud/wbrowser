//
// Author: Marcelo Gomes Jr <marcelo.gomes.junior@gmail.com>
// Created: mai/2021
//
// wbrowser
//

package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// error log
	log := NewErrorLog()
	defer log.Close()

	// read settings file
	src, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatalf("Error reading config file - %v", err)
	}

	var config Config
	err = json.Unmarshal(src, &config)
	if err != nil {
		log.Fatalf("Error parsing config file - %v", err)
	}

	var domain Domain
	var target string
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 && args[0] != "" {
		urlObj, err := url.Parse(args[0])
		if err != nil {
			log.Fatalf("Error parsing url - %v", err)
		}

		domain = Domain(urlObj.Host)
		target = args[0]

		if domain == "" {
			domain = Domain(urlObj.Path)
		}
	}

	if domain == "" || target == "" {
		log.Fatalf("Error - no url")
	}

	// checks if the domain is a redirector
	for _, val := range config.Redirects {
		if val == domain {
			response, err := http.Get(target)
			if err != nil {
				log.Fatalf("Error - url unreachable - %v", err)
			}
			target = response.Request.URL.String()
			domain = Domain(response.Request.URL.Host)
			break
		}
	}

	whichBrowser, ok := config.Domains[domain]
	// if no browser found, check if url contains one of the available domains
	if !ok || whichBrowser == "" {
		for k, v := range config.Domains {
			if strings.Contains(string(k), string(domain)) {
				whichBrowser = v
				break
			}
		}

		// still no browser, reverts to default
		if whichBrowser == "" {
			whichBrowser = config.Domains[DefaultBrowser]
		}
	}

	// run the chosen browser
	cmd := exec.Command(string(config.Browsers[whichBrowser]), args...)
	if err := cmd.Start(); err != nil {
		log.Fatalf("Error - could not open browser - %v", err)
	}
	os.Exit(0)
}
