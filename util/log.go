package util

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
)

// ConfigLog configure log file
func ConfigLog(appName string) {
	// set time format log
	customFormatter := new(log.TextFormatter)
	customFormatter.FullTimestamp = true
	customFormatter.TimestampFormat = "2006-01-02 15:04:05.000"
	log.SetFormatter(customFormatter)

	path := "/var/log/" + appName
	writer, err := rotatelogs.New(
		fmt.Sprintf("%s.%s.log", path, "%Y-%m-%d"),
		rotatelogs.WithMaxAge(-1),
		rotatelogs.WithRotationTime(time.Hour*24),
		rotatelogs.WithHandler(rotatelogs.HandlerFunc(func(e rotatelogs.Event) {
			if e.Type() != rotatelogs.FileRotatedEventType {
				return
			}
			compressToGz(e.(*rotatelogs.FileRotatedEvent).PreviousFile())
		})),
	)
	if err != nil {
		log.Fatalf("Failed to Initialize Log File %s", err)
	}
	log.SetOutput(io.MultiWriter(writer, os.Stdout))
	log.Info("Config Log ready!")
	return
}

func compressToGz(file string) {
	log.Printf("Start compressing file %s", file)
	f, _ := os.Open(file)
	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)
	_ = os.Remove(file)
	f, _ = os.Create(file + ".gz")
	w := gzip.NewWriter(f)
	_, _ = w.Write(content)
	_ = w.Close()
	log.Printf("Finish compressing file %s", file)
}
