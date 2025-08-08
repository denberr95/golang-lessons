package config

import (
	"main/util"
	"strconv"
	"time"
)

type WebServerConfig struct {
	Host                     string `mapstructure:"host"`
	Port                     int    `mapstructure:"port"`
	UseHostname              bool   `mapstructure:"useHostname"`
	BasePath                 string `mapstructure:"basePath"`
	ReadTimeout              int    `mapstructure:"readTimeout"`
	WriteTimeout             int    `mapstructure:"writeTimeout"`
	MaxHeaderSizeMB          int    `mapstructure:"maxHeaderSizeMB"`
	MaxMultipartMemoryMB     int    `mapstructure:"maxMultipartMemoryMB"`
	IdleTimeout              int    `mapstructure:"idleTimeout"`
	GracefulShutdownTime     int    `mapstructure:"gracefulShutdownTime"`
	EnableAccessLog          bool   `mapstructure:"enableAccessLog"`
	EnableLogMiddleware      bool   `mapstructure:"enableLogMiddleware"`
	EnablePrintExposedRouter bool   `mapstructure:"enablePrintExposedRouter"`
}

type LoggingConfig struct {
	Format            LogFormat `mapstructure:"format"`
	Level             LogLevel  `mapstructure:"level"`
	ReportCaller      bool      `mapstructure:"reportCaller"`
	DisableColors     bool      `mapstructure:"disableColors"`
	DisableTimestamp  bool      `mapstructure:"disableTimestamp"`
	DisableHTMLEscape bool      `mapstructure:"disableHTMLEscape"`
	PrettyPrint       bool      `mapstructure:"prettyPrint"`
	FullTimestamp     bool      `mapstructure:"fullTimestamp"`
	ForceQuote        bool      `mapstructure:"forceQuote"`
}

type Config struct {
	GoApp struct {
		WebServer WebServerConfig `mapstructure:"webserver"`
		Logging   LoggingConfig   `mapstructure:"logging"`
	} `mapstructure:"goapp"`
}

type ProgramFlags struct {
	ConfigFileName string
	FileType       string
	FilePath       string
}

type LogFormat int
type LogLevel int

var logFormat = map[LogFormat]string{
	FormatText: "text",
	FormatJSON: "json",
}

var logLevel = map[LogLevel]string{
	PANIC: "panic",
	FATAL: "fatal",
	ERROR: "error",
	WARN:  "warn",
	INFO:  "info",
	DEBUG: "debug",
	TRACE: "trace",
}

const (
	PANIC LogLevel = iota
	FATAL
	ERROR
	WARN
	INFO
	DEBUG
	TRACE
)

const (
	FormatText LogFormat = iota
	FormatJSON
)

func (f LogFormat) String() string {
	return logFormat[f]
}

func (l LogLevel) String() string {
	return logLevel[l]
}

func (f LogFormat) IsValid() bool {
	switch f {
	case FormatText, FormatJSON:
		return true
	default:
		return false
	}
}

func (l LogLevel) IsValid() bool {
	switch l {
	case PANIC, FATAL, ERROR, WARN, INFO, DEBUG, TRACE:
		return true
	default:
		return false
	}
}

func (cfg WebServerConfig) GetFullAddress() string {
	portStr := strconv.Itoa(cfg.Port)
	return cfg.Host + util.Colon + portStr
}

func (cfg WebServerConfig) GetMaxHeaderSizeMB() int64 {
	return util.ShiftMB(cfg.MaxHeaderSizeMB)
}

func (cfg WebServerConfig) GetMaxHeaderBytes() int {
	return util.ConvertMegabitesToBytes(cfg.MaxHeaderSizeMB)
}

func (cfg WebServerConfig) GetReadTimeout() time.Duration {
	return time.Duration(cfg.ReadTimeout) * time.Second
}

func (cfg WebServerConfig) GetWriteTimeout() time.Duration {
	return time.Duration(cfg.WriteTimeout) * time.Second
}

func (cfg WebServerConfig) GetIdleTimeout() time.Duration {
	return time.Duration(cfg.IdleTimeout) * time.Second
}
