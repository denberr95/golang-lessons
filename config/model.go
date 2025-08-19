package config

import (
	"fmt"
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
	Base    *LoggingBaseConfig  `mapstructure:"base"`
	Loggers map[string]LogLevel `mapstructure:"loggers"`
	Text    *LoggingTextConfig  `mapstructure:"text"`
	JSON    *LoggingJSONConfig  `mapstructure:"json"`
}

type LoggingBaseConfig struct {
	ReportCaller bool `mapstructure:"reportCaller"`
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
	p := strconv.Itoa(cfg.Port)
	return cfg.Host + util.Colon + p
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

func (cfg *Config) PrintProperties() []string {
	result := []string{}
	result = append(result, cfg.propertiesWebServerBaseConfig()...)
	result = append(result, cfg.propertiesWebServerHttpConfig()...)
	result = append(result, cfg.propertiesWebServerLogConfig()...)
	result = append(result, cfg.propertiesLoggingBaseConfig()...)
	result = append(result, cfg.propertiesLoggingLoggers()...)
	result = append(result, cfg.propertiesLoggingTextConfig()...)
	result = append(result, cfg.propertiesLoggingJSONConfig()...)
	return result
}

func (cfg *Config) propertiesWebServerBaseConfig() []string {
	return []string{
		fmt.Sprintf("goapp.webserver.base.host=%s", cfg.GoApp.WebServer.Base.Host),
		fmt.Sprintf("goapp.webserver.base.port=%d", cfg.GoApp.WebServer.Base.Port),
		fmt.Sprintf("goapp.webserver.base.useHostname=%t", cfg.GoApp.WebServer.Base.UseHostname),
		fmt.Sprintf("goapp.webserver.base.basePath=%s", cfg.GoApp.WebServer.Base.BasePath),
	}
}

func (cfg *Config) propertiesWebServerHttpConfig() []string {
	return []string{
		fmt.Sprintf("goapp.webserver.http.readTimeout=%s", cfg.GoApp.WebServer.HTTP.ReadTimeout),
		fmt.Sprintf("goapp.webserver.http.writeTimeout=%s", cfg.GoApp.WebServer.HTTP.WriteTimeout),
		fmt.Sprintf("goapp.webserver.http.maxHeaderSizeMB=%d", cfg.GoApp.WebServer.HTTP.MaxHeaderSizeMB),
		fmt.Sprintf("goapp.webserver.http.maxMultipartMemoryMB=%d", cfg.GoApp.WebServer.HTTP.MaxMultipartMemoryMB),
		fmt.Sprintf("goapp.webserver.http.idleTimeout=%s", cfg.GoApp.WebServer.HTTP.IdleTimeout),
		fmt.Sprintf("goapp.webserver.http.gracefulShutdownTime=%s", cfg.GoApp.WebServer.HTTP.GracefulShutdownTime),
	}
}

func (cfg *Config) propertiesWebServerLogConfig() []string {
	return []string{
		fmt.Sprintf("goapp.webserver.log.enableAccessLog=%t", cfg.GoApp.WebServer.Log.EnableAccessLog),
		fmt.Sprintf("goapp.webserver.log.enableLogMiddleware=%t", cfg.GoApp.WebServer.Log.EnableLogMiddleware),
		fmt.Sprintf("goapp.webserver.log.enablePrintExposedRouter=%t", cfg.GoApp.WebServer.Log.EnablePrintExposedRouter),
	}
}

func (cfg *Config) propertiesLoggingBaseConfig() []string {
	return []string{
		fmt.Sprintf("goapp.logging.base.reportCaller=%t", cfg.GoApp.Logging.Base.ReportCaller),
	}
}

func (cfg *Config) propertiesLoggingLoggers() []string {
	props := []string{}
	for k, v := range cfg.GoApp.Logging.Loggers {
		props = append(props, fmt.Sprintf("goapp.logging.loggers.%s=%d", k, v))
	}
	return props
}

func (cfg *Config) propertiesLoggingTextConfig() []string {
	return []string{
		fmt.Sprintf("goapp.logging.text.forceQuote=%t", cfg.GoApp.Logging.Text.ForceQuote),
		fmt.Sprintf("goapp.logging.text.disableColors=%t", cfg.GoApp.Logging.Text.DisableColors),
		fmt.Sprintf("goapp.logging.text.fullTimestamp=%t", cfg.GoApp.Logging.Text.FullTimestamp),
	}
}

func (cfg *Config) propertiesLoggingJSONConfig() []string {
	return []string{
		fmt.Sprintf("goapp.logging.json.disableTimestamp=%t", cfg.GoApp.Logging.JSON.DisableTimestamp),
		fmt.Sprintf("goapp.logging.json.disableHTMLEscape=%t", cfg.GoApp.Logging.JSON.DisableHTMLEscape),
		fmt.Sprintf("goapp.logging.json.prettyPrint=%t", cfg.GoApp.Logging.JSON.PrettyPrint),
	}
}
