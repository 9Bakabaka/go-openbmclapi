/**
 * OpenBmclAPI (Golang Edition)
 * Copyright (C) 2023 Kevin Z <zyxkad@gmail.com>
 * All rights reserved
 * 
 *  This program is free software: you can redistribute it and/or modify
 *  it under the terms of the GNU General Public License as published by
 *  the Free Software Foundation, either version 3 of the License, or
 *  (at your option) any later version.
 *
 *  This program is distributed in the hope that it will be useful,
 *  but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *  GNU General Public License for more details.
 *
 *  You should have received a copy of the GNU General Public License
 *  along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"sync/atomic"
	"time"

	ufile "github.com/KpnmServer/go-util/file"
)

var logdir string = "logs"
var logfile atomic.Pointer[os.File]

var logTimeFormat string = "15:04:05"

func logX(x string, args ...any) {
	sa := make([]string, len(args))
	for i, _ := range args {
		sa[i] = fmt.Sprint(args[i])
	}
	c := strings.Join(sa, " ")
	buf := bytes.NewBuffer(nil)
	buf.Grow(6 + len(logTimeFormat) + len(x) + len(c)) // (1 + len(x) + 2 + len(logTimeFormat) + 3)
	buf.WriteString("[")
	buf.WriteString(x)
	buf.WriteString("][")
	buf.WriteString(time.Now().Format(logTimeFormat))
	buf.WriteString("]: ")
	buf.WriteString(c)
	fmt.Fprintln(os.Stderr, buf.String())
	if fd := logfile.Load(); fd != nil {
		fd.Write(buf.Bytes())
		fd.Write([]byte{'\n'})
	}
}

func logXf(x string, format string, args ...any) {
	c := fmt.Sprintf(format, args...)
	buf := bytes.NewBuffer(nil)
	buf.Grow(6 + len(logTimeFormat) + len(x) + len(c)) // (1 + len(x) + 2 + len(logTimeFormat) + 3)
	buf.WriteString("[")
	buf.WriteString(x)
	buf.WriteString("][")
	buf.WriteString(time.Now().Format(logTimeFormat))
	buf.WriteString("]: ")
	buf.WriteString(c)
	fmt.Fprintln(os.Stderr, buf.String())
	if fd := logfile.Load(); fd != nil {
		fd.Write(buf.Bytes())
		fd.Write([]byte{'\n'})
	}
}

func logDebug(args ...any) {
	if config.Debug {
		logX("DBUG", args...)
	}
}

func logDebugf(format string, args ...any) {
	if config.Debug {
		logXf("DBUG", format, args...)
	}
}

func logInfo(args ...any) {
	logX("INFO", args...)
}

func logInfof(format string, args ...any) {
	logXf("INFO", format, args...)
}

func logWarn(args ...any) {
	logX("WARN", args...)
}

func logWarnf(format string, args ...any) {
	logXf("WARN", format, args...)
}

func logError(args ...any) {
	logX("ERRO", args...)
}

func logErrorf(format string, args ...any) {
	logXf("ERRO", format, args...)
}

func flushLogfile() {
	if ufile.IsNotExist(logdir) {
		ufile.CreateDir(logdir)
	}
	lfile, err := os.OpenFile(ufile.JoinPath(logdir, time.Now().Format("20060102-15.log")),
		os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		logError("Create new log file error:", err)
		return
	}

	logfile.Store(lfile)
}

func startFlushLogFile() {
	flushLogfile()
	logfile.Load().Write(([]byte)("================================================================\n"))
	go func() {
		tma := (time.Now().Unix()/(60*60) + 1) * (60 * 60)
		for {
			select {
			case <-time.After(time.Duration(tma-time.Now().Unix()) * time.Second):
				tma = (time.Now().Unix()/(60*60) + 1) * (60 * 60)
				flushLogfile()
			}
		}
	}()
}