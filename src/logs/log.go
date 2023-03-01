package logs

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	logPath  = "/usr/local"
	encoding = "gb18030"
)

type logFormatter struct{} // 日志自定义格式

type LogFileWriter struct {
	file     *os.File
	fileDate string //判断日期切换目录
	appName  string
	encoding string
}

func (s *logFormatter) Format(entry *log.Entry) ([]byte, error) {
	// Format 格式详情
	formatTime := entry.Time.Format("2006/01/02 15:04:05")
	var file string
	var len int
	if entry.Caller != nil {
		file = filepath.Base(entry.Caller.File)
		len = entry.Caller.Line
	}
	msg := fmt.Sprintf("%s [%s:%d] [%s] %s\n", formatTime, file, len, strings.ToUpper(entry.Level.String()), entry.Message)
	return []byte(msg), nil
}

func (p *LogFileWriter) Write(data []byte) (n int, err error) {
	if p == nil {
		return 0, errors.New("logFileWriter is nil")
	}
	if p.file == nil {
		return 0, errors.New("file not opened")
	}

	//判断是否需要切换日期
	fileDate := time.Now().Format("20060102")
	if p.fileDate != fileDate {
		p.file.Close()
		err = os.MkdirAll(fmt.Sprintf("%s/%s", logPath, p.appName), os.ModePerm)
		if err != nil {
			return 0, err
		}
		filename := fmt.Sprintf("%s/%s/%s-%s.log", logPath, p.appName, fileDate)

		p.file, err = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
		if err != nil {
			return 0, err
		}

	}

	n, e := p.file.Write(data)
	return n, e

}

func LogInit(appName string) {
	var logWriter LogFileWriter

	// InitLog 初始化日志
	fileDate := time.Now().Format("20060102")
	//创建目录
	err := os.MkdirAll(fmt.Sprintf("%s/%s", logPath, appName), os.ModePerm)
	if err != nil {
		log.Error(err)
	}

	filename := fmt.Sprintf("%s/%s/%s.log", logPath, appName, fileDate)
	// /usr/local/appname/data
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		log.Error(err)

	}

	logWriter.file = file
	logWriter.encoding = encoding
	logWriter.appName = appName

	log.SetOutput(&logWriter)
	log.SetReportCaller(true)
	log.SetFormatter(new(logFormatter))
}
