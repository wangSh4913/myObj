package mylog

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	DEBUG uint8 = iota
	WARING
	INFO
	ERROR
	FATAL
)

type log struct {
	File *os.File
	Level uint8
	name string
}

func NewLogFile(fileName string, level uint8)*log{
	file := os.Stdout
	if fileName == "stdout" {
		file = os.Stdout
	}else{
		file, _ = os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
	}
	return &log{
		File:file,
		Level:level,
		name:fileName,
	}
}

func writeLog(l *log, level uint8, format string,a ...interface{}){
	strTime := time.Now().Format("2006-01-02 15:04:05.000")
	var strLevel string
	switch level {
	case DEBUG:
		strLevel = "DEBUG"
	case WARING:
		strLevel = "WARING"
	case INFO:
		strLevel = "INFO"
	case ERROR:
		strLevel = "ERROR"
	case FATAL:
		strLevel = "FATAL"
	default:
		strLevel = ""
	}
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		_, _ = fmt.Fprintf(l.File, "runtime.Caller error")
	}
	resFile := strings.Split(file,"/")[2]
	strLine := strconv.Itoa(line)
	format = strTime  + "[" + strLevel + " " + resFile + " " + strLine + "]" + format

	fileInfo, err := l.File.Stat()
	if err != nil {
		fmt.Println(err)
	}
	if fileInfo.Size() > 1024 * 1024 {
		err := l.File.Close()
		if err != nil {
			fmt.Println(err)
		}
		newName := strings.Split(l.name,".")[0] + time.Now().Format("20060102150405") + ".log"
		_ = os.Rename(l.name, newName)
		l.File, _ = os.OpenFile(l.name, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
	}
	_, _ = fmt.Fprintf(l.File, format, a...)
}

func (l *log)LogDebug(format string,msg ...interface{}){
	if l.Level <= DEBUG {
		writeLog(l, DEBUG, format, msg...)
	}
}
func (l *log)LogWarning(format string, msg ...interface{}){
	if l.Level <= WARING {
		writeLog(l, WARING, format, msg...)
	}
}
func (l *log)LogInfo(format string,msg ...interface{}){
	if l.Level <= INFO {
		writeLog(l, INFO, format, msg...)
	}
}
func (l *log)LogError(format string,msg ...interface{}){
	if l.Level <= ERROR {
		writeLog(l, ERROR, format, msg...)
	}
}
func (l *log)LogFatal(format string,msg ...interface{}){
	if l.Level <= FATAL {
		writeLog(l, FATAL, format, msg...)
	}
}

