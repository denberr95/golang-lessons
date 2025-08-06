package config

type WebServerConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type LoggingConfig struct {
	Format              LogFormat `mapstructure:"format"`
	Level               LogLevel  `mapstructure:"level"`
	ReportCaller        bool      `mapstructure:"reportCaller"`
	DisableColors       bool      `mapstructure:"disableColors"`
	DisableTimestamp    bool      `mapstructure:"disableTimestamp"`
	DisableHTMLEscaping bool      `mapstructure:"disableHTMLEscaping"`
	PrettyPrint         bool      `mapstructure:"prettyPrint"`
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
