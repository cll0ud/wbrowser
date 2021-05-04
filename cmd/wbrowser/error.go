//
// Author: Marcelo Gomes Jr <marcelo.gomes.junior@gmail.com>
// Created: mai/2021
//
// wbrowser
//

package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"wbrowser/pkg/dialog"
)

type ErrorLog struct {
	l    *log.Logger
	file *os.File
}

func NewErrorLog(dir string) *ErrorLog {
	errFile, err := os.OpenFile(dir+"/error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	writer := io.MultiWriter(os.Stdout, errFile)
	return &ErrorLog{
		log.New(writer, "", log.LstdFlags),
		errFile,
	}
}

func (e *ErrorLog) Fatalf(format string, v ...interface{}) {
	dialog.Show(fmt.Sprintf(format, v))
	e.l.Fatalf(format, v)
}

func (e *ErrorLog) Close() {
	if err := e.file.Close(); err != nil {
		panic(err)
	}
}
