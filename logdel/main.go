package main

import (
	"bufio"
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var optConfDir string
var optDryRun bool
var optDate string

var logFiles []*logFile

type logFile struct {
	LogFilePattern string
	RemainedDays   int
	LogFiles       []string
}

func (logfile *logFile) getLogFiles() {
	// 支持日志文件中日期格式：":", "-", "."和"_"
	var logFileSeps = []string{":", "-", ".", "_"}
	logfile.LogFiles = make([]string, 0)
	// wildcardMaxNum还需要加1是因为最后的日志文件中包含中一个"*"，来达到支持三层目录通配。
	wildcardMaxNum := 3

	for _, logFileSep := range logFileSeps {
		// 支持目录通配符，最大支持三层目录通配，如/tmp/*/*/*/*.log，第一层目录不能为"*"。
		wildcardNum := strings.Count(logfile.LogFilePattern, "/*/")
		if wildcardNum > wildcardMaxNum {
			log.Warnf("配置项中，支持目录通配符，最大支持三层目录通配，如/tmp/*/*/*/*.log，不支持多个*。")
			return
		}

		if strings.Index(logfile.LogFilePattern, "/*/") == 0 {
			log.Warnf("配置项中，第一层目录不能为\"*\"。")
			return
		}

		logFileList := strings.Split(logfile.LogFilePattern, "/")
		logFileList = logFileList[1:] // 去掉最前面的空字符串
		logFileFlag := logFileList[len(logFileList)-1]
		if !strings.Contains(logFileFlag, "*") {
			log.Warnf("配置项中，最后日志文件中没有通配符\"*\"，这个是不可以的，因为当前不能根据给出的目录删除何种日志文件。")
			return
		}
		var topLogFileDir string
		topLogFileDir += "/" + logFileList[0]
		i := 0
	LOOP1:
		dirs, _ := ioutil.ReadDir(topLogFileDir)
		for _, dir := range dirs {
			if dir.IsDir() {
				i++
				topLogFileDir += "/" + dir.Name()
				fmt.Printf("LOOP1: %v\n", topLogFileDir)
				if i != wildcardNum {
					goto LOOP1
				}
				goto LOOP2
			}
		}
	LOOP2:
		fmt.Printf("LOOP2: %v\n", topLogFileDir)
		dirs2, _ := ioutil.ReadDir(topLogFileDir)
		for _, dir := range dirs2 {
			if dir.IsDir() {
				logfile.LogFiles = append(logfile.LogFiles, dir.Name())
			}
		}
		fmt.Println(logFileSep)
	}
}

func (logfile *logFile) delLogFiles() {

}

// getConfFiles 根据--conf-dir参数给定的配置目录，获取到yaml配置文件列表。
func getConfFiles(confDir string) []string {
	confFiles := make([]string, 0)
	files, err := ioutil.ReadDir(confDir)
	if err != nil {
		log.Fatalf("打开配置文件目录\"%s\"失败。", confDir, err)
	}

	for _, file := range files {
		if !file.IsDir() {
			fileName := file.Name()
			// 根据文件后缀".conf"检查是否为配置文件。
			if strings.HasSuffix(fileName, ".conf") {
				confFiles = append(confFiles, optConfDir+"/"+file.Name())
			}
		}
	}
	return confFiles
}

// 过滤掉以#开头的注释。
func getConfItem(str string) {
	reg, _ := regexp.Compile("^(/.*(/.*)*):\\s+(\\d+)\\s*$")
	regStrings := reg.FindAllSubmatch([]byte(str), -1)

	for _, str := range regStrings {
		myLogFile := new(logFile)
		myLogFile.LogFilePattern = string(str[1])
		myLogFile.RemainedDays, _ = strconv.Atoi(string(str[3]))
		logFiles = append(logFiles, myLogFile)
	}
}

// parseConfFile 解析yaml配置文件
func parseConfFile(confFile string) {
	f, err := os.Open(confFile)
	if err != nil {
		log.Warnf("打开配置文件目录\"%s\"失败。", confFile)
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			getConfItem(str)
			return
		}
		if err != nil {
			log.Warnf("读取配置项\"%s\"失败。具体报错：%v\n", str, err)
		}
		getConfItem(str)
	}
}

func main() {
	// 获取参数。
	// 默认配置文件目录在/etc/logdel.d，配置文件以.conf结尾，不能放在这个目录下其他目录下。
	flag.StringVar(&optConfDir, "conf-dir", "/etc/logdel.d", "Set configuration directory.")

	// 运行时，默认跑在dry-run下，并且日志打印在终端，如果不在此模式下就就印在/var/log/logdel.log下，并且日志按100M分割大小，
	// 保留7个副本，不压缩，日志内容格式为json格式。
	flag.BoolVar(&optDryRun, "dry-run", true, "Run in dry-run mode.")

	// 指定日期，用于调试，并且日志打印在终端。
	flag.StringVar(&optDate, "date", "", "Set date (in format of YYYY-MM-DD) for debugging.")
	flag.Parse()

	// 获取配置目录下的配置文件。
	confFiles := getConfFiles(optConfDir)

	if len(confFiles) == 0 {
		log.Warnf("配置文件目录\"%s\"下无配置文件。", optConfDir)
		return
	} else {
		log.Infof("配置文件目录\"%s\"下配置文件列表：%#v。\n", optConfDir, confFiles)
	}

	// 解析配置文件，配置文件默认存放在/etc/logdel.d中，文件是yaml格式。
	for _, confFile := range confFiles {
		parseConfFile(confFile)
	}

	// 根据获取到的每个配置项，获取到要删除的文件列表。
	for _, item := range logFiles {
		item.getLogFiles()
		fmt.Println(item.LogFiles)
	}

	// 根据文件列表删除文件。
}
