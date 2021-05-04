// Edited, Windows-only version of:
// https://github.com/tawesoft/go/tree/master/dialog

package dialog

import (
	"fmt"

	"golang.org/x/sys/windows"
)

func Show(message string, args ...interface{}) {
	windowsAlert("wbrowser", message, args...)
}

func windowsAlert(title string, message string, args ...interface{}) {
	var msg string

	if len(args) > 0 {
		msg = fmt.Sprintf(message, args...)
	} else {
		msg = message
	}

	var wtitle = toWideChar(title)
	var wmessage = toWideChar(msg)
	var flags uint32 = windows.MB_OK |
		windows.MB_ICONEXCLAMATION |
		windows.MB_SETFOREGROUND |
		windows.MB_TOPMOST

	_, err := windows.MessageBox(0, &wmessage[0], &wtitle[0], flags)
	if err != nil {
		panic(err)
	}

	wtitle = nil
	wmessage = nil
}

const cpUtf8 uint32 = 65001

func toWideChar(input string) []uint16 {
	var buf []uint16
	var required, numchars int32
	var err error

	if len(input) == 0 {
		return make([]uint16, 1)
	}

	required, err = windows.MultiByteToWideChar(
		cpUtf8,
		0,
		&([]byte(input))[0],
		int32(len(input)),
		nil,
		0,
	)
	if err != nil {
		return make([]uint16, 1)
	}

	buf = make([]uint16, required+1)
	numchars, err = windows.MultiByteToWideChar(
		cpUtf8,
		0,
		&([]byte(input))[0],
		int32(len(input)),
		&(buf)[0],
		required,
	)

	if err != nil || len(buf) != int(numchars)+1 {
		buf = nil
	}
	return buf
}
