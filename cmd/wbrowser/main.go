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
	"path/filepath"
	"strings"
)

func main() {
	// when running as default browser for Windows, the working directory
	// becomes the place where the program was called, which means we can't
	// know where config or error file is, so we get the absolute path to
	// wbrowser OR current wd if we're debugging
	debug := flag.Bool("debug", false, "debug mode")
	flag.Parse()
	var dir string
	var err error

	if *debug {
		dir, err = os.Getwd()
	} else {
		dir, err = filepath.Abs(filepath.Dir(os.Args[0]))
	}

	// error log
	logger := NewErrorLog(dir)
	defer logger.Close()

	// read settings file
	src, err := ioutil.ReadFile(dir + "/config.json")
	if err != nil {
		logger.Fatalf("Error reading config file - %v", err)
	}

	var config Config
	err = json.Unmarshal(src, &config)
	if err != nil {
		logger.Fatalf("Error parsing config file - %v", err)
	}

	var domain Domain
	var target string
	args := flag.Args()
	if len(args) > 0 && args[0] != "" {
		urlObj, err := url.Parse(args[0])
		if err != nil {
			logger.Fatalf("Error parsing url - %v", err)
		}

		domain = Domain(urlObj.Host)
		target = args[0]

		if domain == "" {
			domain = Domain(urlObj.Path)
		}
	}

	// no domain or url, just open the default browser
	if domain == "" || target == "" {
		run(logger, string(config.Browsers[config.Domains[DefaultBrowser]]), args...)
		return
	}

	// checks if the domain is a redirector so we can fetch the real url
	// and decide which browser to run
	for _, val := range config.Redirects {
		if val == domain {
			response, err := http.Get(target)
			if err != nil {
				logger.Fatalf("Error - url unreachable - %v", err)
			}
			target = response.Request.URL.String()
			domain = Domain(response.Request.URL.Host)
			break
		}
	}

	whichBrowser, ok := config.Domains[domain]
	// if no browser found, check if url contains one of the available domains
	// this will match subdomains, so if you need a specific subdomain to match
	// a different browser you should put that subdomain on the config file
	//
	// e.g.: if you have youtube.com on your config file and the url you're
	// trying to open is "www.youtube.com" the following code will match that
	// if you have "google.com" and the url to open is "mail.google.com" (gmail)
	// or "play.google.com" (playstore) they will also open with the same browser
	// as "google.com" UNLESS you have an specific rule for these subdomains
	if !ok || whichBrowser == "" {
		for k, v := range config.Domains {
			if strings.Contains(string(domain), string(k)) {
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
	run(logger, string(config.Browsers[whichBrowser]), args...)
}

func run(logger *ErrorLog, browser string, args ...string) {
	if err := exec.Command(browser, args...).Start(); err != nil {
		logger.Fatalf("Error - could not open browser - %v", err)
	}
}
