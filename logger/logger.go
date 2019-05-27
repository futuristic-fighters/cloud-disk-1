package logger

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"syscall"
	"time"
)

type Logger interface {
	Write(...string)
}

type logger struct {
	out           io.Writer
	logFilePrefix string
	runtimePath   string
	logPath       string
	logExpire     string
	debug         bool
	Logger
}

func New(RuntimePath, LogPath, LogFilePrefix, LogExpire string, Debug bool) *logger {
	l := &logger{
		runtimePath:   RuntimePath,
		logPath:       LogPath,
		logFilePrefix: LogFilePrefix,
		logExpire:     LogExpire,
		debug:         Debug,
	}

	l.checkLoggerWriter()
	l.startLogsCleaner()
	return l
}

func (l *logger) Write(Strings ...string) {
	s := ""
	for _, v := range Strings {
		s += v + " "
	}

	if l.out == nil {
		fmt.Println("no set logger out")
		l.out = os.Stdout
	}

	_, _ = fmt.Fprintf(l.out, "[Log] %s %s\n",
		time.Now().Format("15:04:05.999999999"),
		s,
	)
}

//自定义日志文件 按天记录
func (l *logger) checkLoggerWriter() {
	l.checkLoggerDir()

	if l.debug == true {
		l.out = os.Stdout
		return
	}

	name := l.logFilePrefix + time.Now().Format("2006-01-02") + ".log"
	file := l.logPath + "/" + name

	_, err := os.Stat(file)
	if err != nil {
		if f, ok := l.out.(*os.File); ok && l.out != nil {
			_ = f.Close()
		}

		var err error
		l.out, err = os.Create(file)
		if err != nil {
			fmt.Println("create", file, "err :", err.Error())
			l.out = os.Stdout
		}
		return
	}

	if l.out != nil {
		return
	}

	var err1 error
	l.out, err1 = os.OpenFile(file, syscall.O_APPEND|syscall.O_RDWR, 0666)
	if err1 != nil {
		fmt.Println("OpenFile", file, "err :", err1.Error())
		l.out = os.Stdout
	}

	return
}

//初始化生成日志文件目录
// 配置runtime目录 .env RUNTIME_PATH
// 配置logs目录名 .env LOG_PATH_NAME
func (l *logger) checkLoggerDir() {
	{
		_, err := os.Stat(l.runtimePath)
		if err != nil {
			err := os.Mkdir(l.runtimePath, 0777)
			if err != nil {
				panic("mkdir " + l.runtimePath + " err:" + err.Error())
			}
		}
	}

	{
		_, err := os.Stat(l.logPath)
		if err != nil {
			err := os.Mkdir(l.logPath, 0777)
			if err != nil {
				panic("mkdir " + l.logPath + " err:" + err.Error())
			}
		}
	}
}

//定时清除日志
// 配置 .env LOG_EXPIRE 设置日志过期时间 以小时为单位
func (l *logger) startLogsCleaner() {
	day, _ := time.ParseDuration(l.logExpire + "h")
	t := time.NewTicker(day)

	go func() {
		for {
			select {
			case <-t.C:
				l.Write("start logs cleaner...")
				_, err := os.Stat(l.logPath)
				if err != nil {
					l.Write(err.Error())
					continue
				}

				files, err1 := ioutil.ReadDir(l.logPath)
				if err1 != nil {
					l.Write(err1.Error())
					continue
				}

				for _, f := range files {
					if !f.IsDir() {
						err := os.Remove(l.logPath + "/" + f.Name())
						if err != nil {
							l.Write(err1.Error())
							continue
						}
						l.Write("remove log file success:", l.logPath+"/"+f.Name())
					}
				}
			}
		}
	}()
}
