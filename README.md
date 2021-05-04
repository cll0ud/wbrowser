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
You need `go 1.16` or a newer version.

Run:
```bash
make clean build
```
Then copy the contents of the dist folder to your desired location and edit the
`config.json` file to fit your needs.

### Using as default browser on Windows
There's probably a better way to do that, but the following anwser on 
StackOverflow worked just fine for me: 
https://stackoverflow.com/questions/32671277/how-do-i-register-a-custom-application-as-a-web-browser-in-windows-8-1
