package setting

import (
    "os"
    "os/exec"
    "path/filepath"
    "strings"
    "path"

    "gopkg.in/ini.v1"

    "github.com/zeuxisoo/go-tormag/config"
    "github.com/zeuxisoo/go-tormag/pkg/logger"
    "github.com/zeuxisoo/go-tormag/pkg/utils"
)

var (
    //
    AppPath         string

    //
    Address         string
    Port            string
    Mode            string
    StoragePath     string

    //
    AttachmentPath      string
    AttachmentMaxSize   int64

    //
    CrossOriginEnable           bool
    CrossOriginAllowOrigins     []string
    CrossOriginAllowMethods     []string
    CrossOriginAllowHeaders     []string
    CrossOriginExposeHeaders    []string
    CrossOriginAllowCredentials bool
    CrossOriginMaxAge           int

    //
    Config          *ini.File
    CustomConfig    string
)

func init() {
    var err error

	if AppPath, err = executablePath(); err != nil {
        logger.Fatalf("Cannot found the app path => %v\n", err)
    }

    AppPath = strings.Replace(AppPath, "\\", "/", -1)
}

func executablePath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
    }

	return filepath.Abs(file)
}

// AppDirectory return the application directory path
func AppDirectory() (string, error) {
    i := strings.LastIndex(AppPath, "/")

    if i == -1 {
        return AppPath, nil
    }

    return AppPath[:i], nil
}

// NewSetting will setup the application config constant
func NewSetting() {
    //
    appDirectory, err := AppDirectory()
	if err != nil {
		logger.Fatalf("Cannot found the app directory => %v", err)
	}

    //
    Config, err = ini.LoadSources(ini.LoadOptions{
        IgnoreInlineComment: true,
    }, config.MustAsset("config/app.ini"))
    if err != nil {
        logger.Fatalf("Cannot load the config/app.ini => %v\n", err)
    }

    if len(CustomConfig) != 0 && utils.IsFile(CustomConfig) == true {
        if err := Config.Append(CustomConfig); err != nil {
            logger.Fatalf("Cannot load the custom config file (%s) => %v\n", CustomConfig, err)
        }
    }

    //
    var section *ini.Section

    section     = Config.Section("server")
    Address     = section.Key("ADDRESS").MustString("0.0.0.0")
    Port        = section.Key("PORT").MustString("3000")
    Mode        = section.Key("MODE").MustString("dev")
    StoragePath = section.Key("STORAGE_PATH").MustString("storage")

    section           = Config.Section("attachment")
    AttachmentPath    = section.Key("PATH").MustString(path.Join(StoragePath, "attachments"))
    AttachmentMaxSize = section.Key("MAX_SIZE").MustInt64(4)

    section                     = Config.Section("cors")
    CrossOriginEnable           = section.Key("ENABLE").MustBool(false)
    CrossOriginAllowOrigins     = section.Key("ALLOW_ORIGINS").Strings(",")
    CrossOriginAllowMethods     = section.Key("ALLOW_METHODS").Strings(",")
    CrossOriginAllowHeaders     = section.Key("ALLOW_HEADERS").Strings(",")
    CrossOriginExposeHeaders    = section.Key("EXPOSE_HEADERS").Strings(",")
    CrossOriginAllowCredentials = section.Key("ALLOW_CREDENTIALS").MustBool(true)
    CrossOriginMaxAge           = section.Key("MAX_AGE").MustInt(12)

    if filepath.IsAbs(AttachmentPath) == false {
        AttachmentPath = path.Join(appDirectory, AttachmentPath)
    }
}
