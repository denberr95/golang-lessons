package config

import (
	"main/util"
	"strconv"
	"time"
)

type WebServerConfig struct {
	Base WebServerBaseConfig `mapstructure:"base"`
	HTTP WebServerHTTPConfig `mapstructure:"http"`
	Log  WebServerLogConfig  `mapstructure:"log"`
}

type WebServerBaseConfig struct {
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	UseHostname bool   `mapstructure:"useHostname"`
	BasePath    string `mapstructure:"basePath"`
}

type WebServerHTTPConfig struct {
	ReadTimeout          time.Duration `mapstructure:"readTimeout"`
	WriteTimeout         time.Duration `mapstructure:"writeTimeout"`
	MaxHeaderSizeMB      int           `mapstructure:"maxHeaderSizeMB"`
	MaxMultipartMemoryMB int           `mapstructure:"maxMultipartMemoryMB"`
	IdleTimeout          time.Duration `mapstructure:"idleTimeout"`
	GracefulShutdownTime time.Duration `mapstructure:"gracefulShutdownTime"`
}

type WebServerLogConfig struct {
	EnableAccessLog          bool `mapstructure:"enableAccessLog"`
	EnableLogMiddleware      bool `mapstructure:"enableLogMiddleware"`
	EnablePrintExposedRouter bool `mapstructure:"enablePrintExposedRouter"`
}

type LoggingConfig struct {
	Base *LoggingBaseConfig `mapstructure:"base"`
	Text *LoggingTextConfig `mapstructure:"text"`
	JSON *LoggingJSONConfig `mapstructure:"json"`
}

type LoggingBaseConfig struct {
	Level        LogLevel `mapstructure:"level"`
	ReportCaller bool     `mapstructure:"reportCaller"`
}

type LoggingTextConfig struct {
	ForceQuote    bool `mapstructure:"forceQuote"`
	DisableColors bool `mapstructure:"disableColors"`
	FullTimestamp bool `mapstructure:"fullTimestamp"`
}

type LoggingJSONConfig struct {
	DisableTimestamp  bool `mapstructure:"disableTimestamp"`
	DisableHTMLEscape bool `mapstructure:"disableHTMLEscape"`
	PrettyPrint       bool `mapstructure:"prettyPrint"`
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

type LogLevel int

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

func (l LogLevel) String() string {
	return logLevel[l]
}

func (l LogLevel) IsValid() bool {
	switch l {
	case PANIC, FATAL, ERROR, WARN, INFO, DEBUG, TRACE:
		return true
	default:
		return false
	}
}

func (cfg WebServerBaseConfig) GetFullAddress() string {
	portStr := strconv.Itoa(cfg.Port)
	return cfg.Host + util.Colon + portStr
}

func (cfg WebServerHTTPConfig) GetMaxHeaderSizeMB() int64 {
	return util.ShiftMB(cfg.MaxHeaderSizeMB)
}

func (cfg WebServerHTTPConfig) GetMaxHeaderBytes() int {
	return util.ConvertMegabitesToBytes(cfg.MaxHeaderSizeMB)
}

func (cfg WebServerHTTPConfig) GetMaxMultipartMemoryMB() int64 {
	return util.ShiftMB(cfg.MaxMultipartMemoryMB)
}

func (cfg *Config) Print() string {
	// TODO
	return ""
}
