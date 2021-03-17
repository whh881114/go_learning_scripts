package main

// 程序名：logdel
// 用途：删除带有日期文件，特别是java应用产生的日志，例如tomcat应用日志。

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"time"
)

var optConfDir string
var optDryRun bool

type Configurations struct {
	Global Global  `yaml:"global"`
	Items  []Items `yaml:"items"`
}

type Global struct {
	RemainedDays int      `yaml:"remained_days"`
	DateFormats  []string `yaml:"date_formats"`
	Suffixes     []string `yaml:"suffixes"`
}

type Items struct {
	Name         string   `yaml:"name"`
	Paths        []string `yaml:"paths"`
	RemainedDays int      `yaml:"remained_days"`
	DateFormats  []string `yaml:"date_formats"`
	Suffixes     []string `yaml:"suffixes"`
}

func getConfFiles(confDir string) []string {
	confFiles := make([]string, 0)
	files, err := ioutil.ReadDir(confDir)
	if err != nil {
		log.Fatalf("打开配置文件目录\"%s\"失败。", confDir, err)
	}

	for _, file := range files {
		if !file.IsDir() {
			fileName := file.Name()
			// 配置文件后缀名为`.yaml`或`.yml`。
			if strings.HasSuffix(strings.ToLower(fileName), ".yaml") || strings.HasSuffix(strings.ToLower(fileName), ".yml") {
				confFiles = append(confFiles, optConfDir+"/"+file.Name())
			}
		}
	}
	return confFiles
}

func decodeConfFile(confFile string) (*Configurations, error) {
	conf := new(Configurations)

	f, err := os.Open(confFile)
	if err != nil {
		log.Errorf("打开配置文件目录\"%s\"失败。\n", confFile)
		return conf, err
	}
	defer f.Close()

	if err := yaml.NewDecoder(f).Decode(conf); err != nil {
		return conf, err
	}

	return conf, nil
}

func parseConf(confFile string, confData *Configurations) []Items {
	instances := make([]Items, 0)

	// 当全局参数没有配置时，使用默认值。
	defaultRemainedDays := 7
	defaultDateFormats := []string{"YYYY-MM-DD", "YYYYMMDD", "YYYY_MM_DD"}
	defaultSuffixes := []string{".log", ".txt"}

	// 当全局参数配置时，覆盖默认值。
	if confData.Global.RemainedDays != 0 {
		defaultRemainedDays = confData.Global.RemainedDays
	}

	if len(confData.Global.DateFormats) != 0 {
		defaultDateFormats = confData.Global.DateFormats
	}

	if len(confData.Global.Suffixes) != 0 {
		defaultSuffixes = confData.Global.Suffixes
	}

	// 当配置文件中，没有填写items配置时，其中name和path均为默认值，则过滤此配置。
	for _, item := range confData.Items {
		if len(item.Name) == 0 {
			log.Warnf("配置文件\"%s\"配置项items子项目\"%#v\"中\"name\"参数为空，忽略此配置项。", confFile, item)
			continue
		}

		if len(item.Paths) == 0 {
			log.Warnf("配置文件\"%s\"配置项items子项目\"%#v\"中\"paths\"参数为空，忽略此配置项。", confFile, item)
			continue
		}

		if item.RemainedDays == 0 {
			item.RemainedDays = defaultRemainedDays
		}

		if len(item.DateFormats) == 0 {
			item.DateFormats = defaultDateFormats
		}

		if len(item.Suffixes) == 0 {
			item.Suffixes = defaultSuffixes
		}

		instances = append(instances, item)
	}
	return instances
}

func delLogFiles(logItems []Items, optDryRun bool) {
	for _, item := range logItems {
		// 第一步，生成正则表达式。
		regexpString := ""
		tmpRegexpString := ""
		for _, logSuffix := range item.Suffixes {
			for _, dateFormat := range item.DateFormats {
				tmpRegexpString += strings.ToLower(dateFormat) + strings.ToLower(logSuffix) + "$|"
				tmpRegexpString += strings.ToLower(logSuffix) + strings.ToLower(dateFormat) + "$|"

				regexpString = tmpRegexpString[:len(tmpRegexpString)-1]
				regexpString = strings.ReplaceAll(regexpString, "yyyy", "\\d{4}")
				regexpString = strings.ReplaceAll(regexpString, "mm", "\\d{1,2}")
				regexpString = strings.ReplaceAll(regexpString, "dd", "\\d{1,2}")
			}
		}

		log.Infof("开始处理配置文件中items配置项：%#v。\n", item)

		// 第二步，读取目录，进行匹配文件。
		for _, path := range item.Paths {
			files, err := ioutil.ReadDir(path)
			if err != nil {
				log.Errorf("读取\"%s\"目录出错。错误为：%#v\n", path, err)
				continue
			} else {
				log.Infof("读取\"%s\"目录成功。\n", path)
			}

			log.Infof("当前匹配模式：%#v\n", regexpString)
			for _, file := range files {
				if !file.IsDir() {
					fileName := file.Name()
					validString := regexp.MustCompile(regexpString)
					if validString.MatchString(fileName) {
						log.Infof("文件名格式匹配成功：%s。\n", fileName)
						currentTime := time.Now()
						remainedTime := currentTime.AddDate(0, 0, -item.RemainedDays)
						if file.ModTime().After(remainedTime) {
							log.Warnf("\"%s\"文件在保留时间内，其文件最后修改时间：%v\n", fileName, file.ModTime())
						} else {
							log.Infof("%s文件不在保留时间内，文件最后修改时间：%v\n", fileName, file.ModTime())
							logFile := path + string(os.PathSeparator) + fileName
							if err := os.Remove(logFile); err != nil {
								log.Infof("文件删除失败：%s，原因：%#v\n", logFile, err)
							} else {
								log.Infof("文件删除成功：%s\n", logFile)
							}
						}
					}
				}
			}
		}
	}
}

func main() {
	// 获取参数。
	// 默认配置文件目录在/etc/logdel.d，配置文件以`.yaml`或`.yml`结尾，不能放在这个目录下其他目录下。
	flag.StringVar(&optConfDir, "conf-dir", "/etc/logdel.d", "Set configuration directory.")

	// 运行时，默认跑在dry-run下，并且日志打印在终端。
	flag.BoolVar(&optDryRun, "dry-run", true, "Run in dry-run mode.")

	flag.Parse()

	// 获取配置目录下的配置文件。
	confFiles := getConfFiles(optConfDir)

	if len(confFiles) == 0 {
		log.Errorf("配置文件目录\"%s\"下无配置文件。", optConfDir)
		return
	} else {
		log.Infof("配置文件目录\"%s\"下配置文件列表：%v。\n", optConfDir, confFiles)
	}

	for _, confFile := range confFiles {
		if confData, err := decodeConfFile(confFile); err != nil {
			log.Errorf("解析配置文件\"%s\"失败，跳过此配置文件。\n", confFile)
		} else {
			log.Infof("解析配置文件\"%s\"成功。\n", confFile)
			logItems := parseConf(confFile, confData)

			if len(logItems) == 0 {
				log.Warnf("配置文件\"%s\"中无items配置项。\n", confFile)
			} else {
				delLogFiles(logItems, optDryRun)
			}
		}
	}
}
