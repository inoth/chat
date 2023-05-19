package zap

import (
	"chat/toybox/components/logger"
	"sync"

	"go.uber.org/zap"
)

var (
	componentName = "logger"
	logOnce       sync.Once
)

// err_log: log/err.log
// warn_log: log/warn.log
// info_log: log/info.log
// maxsize: 100
// maxage: 15
// max_backup: 30
// compress: true
// json: true
type ZapLogComponent struct {
	zap *zap.Logger

	ErrLog    string `json:"err_log" yaml:"err_log" toml:"err_log"`
	WarnLog   string `json:"warn_log" yaml:"warn_log" toml:"warn_log"`
	InfoLog   string `json:"info_log" yaml:"info_log" toml:"info_log"`
	DebugLog  string `json:"debug_log" yaml:"debug_log" toml:"debug_log"`
	MaxSize   int    `json:"max_size" yaml:"max_size" toml:"max_size"`
	MaxAge    int    `json:"max_age" yaml:"max_age" toml:"max_age"`
	MaxBackup int    `json:"max_backup" yaml:"max_backup" toml:"max_backup"`
	Compress  bool   `json:"compress" yaml:"compress" toml:"compress"`
	Json      bool   `json:"json" yaml:"json" toml:"json"`
}

func New() logger.ToyBoxLogger {
	return nil
}

func (zc *ZapLogComponent) Name() string   { return componentName }
func (zc *ZapLogComponent) String() string { return "" }
func (zc *ZapLogComponent) Close() error   { return nil }

func (zc *ZapLogComponent) Init() (err error) {
	logOnce.Do(func() {
		logger.Log = zc
	})
	return
}

func (zc *ZapLogComponent) Debug(args ...interface{})
func (zc *ZapLogComponent) Debugf(msg string, args ...interface{})

func (zc *ZapLogComponent) Info(args ...interface{})
func (zc *ZapLogComponent) Infof(msg string, args ...interface{})

func (zc *ZapLogComponent) Warn(args ...interface{})
func (zc *ZapLogComponent) Warnf(msg string, args ...interface{})

func (zc *ZapLogComponent) Err(args ...interface{})
func (zc *ZapLogComponent) Errf(msg string, args ...interface{})
