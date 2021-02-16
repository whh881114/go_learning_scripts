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

	// "strconv"
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

func (logfile *logFile) getLogFiles(logFile logFile) {

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

	if len(confFiles) == 0 {
		log.Warnf("配置文件目录\"%s\"下无配置文件。", confDir)
	}

	return confFiles
}

// 过滤掉以#开头的注释。
func filterExplanations(str string) {
	// var logFilePattern string
	// var remainedDays int

	// re := regexp.MustCompile("^/.*(/.*)+:\\ +[0-9]+")
	// fmt.Printf("%q\n", re.FindString(str))

	// ok, _ := regexp.Match("^/.*(/.*)*: +\\d+$", []byte(str))
	// ok, err := regexp.Match(`^/.*(/.*)*: +\\d+$`, []byte(str))
	// if ok {
	// 	fmt.Print(str)
	// 	fmt.Println("匹配到了")
	// 	fmt.Println()
	// } else {
	// 	fmt.Println(err)
	// }

	reg, _ := regexp.Compile("^(/.*(/.*)*):\\s+(\\d+)\\s*$")
	strings := reg.FindAllSubmatch([]byte(str), -1)

	for _, str := range strings {
		fmt.Print(string(str[1]))
		fmt.Print("----------->")
		fmt.Println(string(str[3]))
	}

	// reg := regexp.MustCompile("^/.*(/.*)*:\\s+\\d+\\s*$")
	// fmt.Printf("%q\n", reg.FindAllString(str, -1))

	// if ok {
	// 	log.Infof("解析到配置项：%s", str)
	// 	ret := strings.Split(str, ":")
	// 	logFilePattern = strings.TrimSpace(ret[0])
	// 	value, err := strconv.Atoi(strings.TrimSpace(ret[1]))
	// 	if err != nil {
	// 		log.Warnf("配置项\"%s\"的值不正确，应该要为整型。", str)
	// 	}
	// 	remainedDays = value
	// }
	//
	// myLogFile := logFile{
	// 	LogFilePattern: logFilePattern,
	// 	RemainedDays:   remainedDays,
	// }
	// logFiles = append(logFiles, &myLogFile)
}

// parseConfFile 解析yaml配置文件
// func parseConfFile(confFile string) (logFile, error) {
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
			return
		}
		if err != nil {
			log.Warnf("读取配置项\"%s\"失败。具体报错：%v\n", str, err)
		}
		filterExplanations(str)
	}

	// config, err := ioutil.ReadFile(confFile)
	// if err != nil {
	// 	log.Warnf("打开配置文件目录\"%s\"失败。", confFile)
	// }
	//
	// // fmt.Println(string(config))
	//
	// if err := yaml.Unmarshal(config, &settings); err != nil {
	// 	log.Warnf("Warning: %v", err)
	// }
	// fmt.Printf("%#v\n", settings)

	// fmt.Println(settings.LogFilePattern)
	// fmt.Println(settings.RemainedDays)

	// var conf yamlConfFile
	// yaml.Unmarshal(conf, &yamlConfFile{})
	//
	// if err != nil {
	// 	log.Fatalf("配置文件\"%s\"无法被yaml解析。", yamlConfFile)
	// }
	// fmt.Println(string(yamlConfFile))
	// fmt.Println(string(yamlConfFile))
	//
	// return logFile{
	// 	LogFilePattern: "/some/path/file_pattern",
	// 	RemainedDays:   0,
	// }, err
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

	confFiles := getConfFiles(optConfDir)
	fmt.Println(confFiles)

	// 解析配置文件，配置文件默认存放在/etc/logdel.d中，文件是yaml格式。
	for _, confFile := range confFiles {
		parseConfFile(confFile)
		// _, err := parseConfFile(confFile)
		// if err != nil {
		// 	continue
		// }

	}

	// 根据获取到的每个配置项，获取到要删除的文件列表。

	// 根据文件列表删除文件。
}
