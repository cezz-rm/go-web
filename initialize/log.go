package initialize

import (
	"path"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
)

func InitLogger(logPath, logFile string) {
	baseLogPath := path.Join(logPath, logFile)
	writer, err := rotatelogs.New(
		baseLogPath+".%Y%m%d%H%M",
		//rotatelogs.WithLinkName(path),             // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)
	if err != nil {
		log.Errorf("config local file system logger error. %v", errors.WithStack(err))
	}

	lfHook := lfshook.NewHook(
		lfshook.WriterMap{
			log.DebugLevel: writer, // 为不同级别设置不同的输出目的
			log.InfoLevel:  writer,
			log.WarnLevel:  writer,
			log.ErrorLevel: writer,
			log.FatalLevel: writer,
			log.PanicLevel: writer,
		},
		&log.JSONFormatter{},
	)
	log.SetReportCaller(true)
	log.AddHook(lfHook)
}
