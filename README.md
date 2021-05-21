# wbrowser

`wbrowser` (_"which-browser"_) is a simple tool that opens different browsers 
for specific domains.

I use a couple different browsers on a daily basis for multiple tasks, like 
watching videos, testing/debugging local
projects, reading emails, etc. So I decided to make a simple tool that knows 
which browser to open when I click a link
on Telegram, Twitter and pretty much any other place outside a browser.

Works on Windows and it'd probably work on Linux with a couple tweaks. I don't 
use Mac so I can't help with that.

### Usage
Make sure you have a `config.json` file on the same folder as the executable.
Then just run the program with the url you want to open as an argument:
```bash
wbrowser https://github.com/cll0ud/wbrowser
```

Please note that the config file will match subdomains if you don't have 
specific rules. `wbrowser` will always look for a rule for that specific domain
first, if a rule is not found then it'll try to look for a string match on all
available domains.

E.g:
```json5
{
  "domains": {
    "default": "edge",
    // this will match "youtube.com" and "www.youtube.com"
    "youtube.com": "chrome",
    // this will match *any* google.com subdomain, like:
    //  - www.google.com
    //  - mail.google.com (gmail)
    //  - play.google.com (playstore)
    //  - accounts.google.com
    //  etc
    "google.com":  "chrome",
    // but if you create a specific rule to a specific subdomain
    // then that rule will have priority over the "catch-all"
    "mail.google.com": "edge",
  }
}
```

### Build
You need `go 1.16` or a newer version.

Run:
```bash
make clean build
```
Then copy the contents of the dist folder to your desired location and edit the
`config.json` file to fit your needs.

### Using as default browser on Windows
You need to register `wbrowser` as one of the available default browsers
for Windows.

There's probably a better way to do that, but the following anwser on 
StackOverflow worked just fine for me: 
https://stackoverflow.com/questions/32671277/how-do-i-register-a-custom-application-as-a-web-browser-in-windows-8-1
