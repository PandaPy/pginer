package logger

import (
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

// InitLogger 初始化日志输出格式，控制台输出带颜色，固定间距
func InitLogger() {
	// 定义日志文件路径
	logDir := "log"
	logFile := filepath.Join(logDir, "output.log")

	// 确保日志文件夹存在
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err := os.Mkdir(logDir, 0755)
		if err != nil {
			panic(err)
		}
	}

	encoderConfig := zapcore.EncoderConfig{
		LevelKey:     "level",
		TimeKey:      "time",
		CallerKey:    "caller",
		MessageKey:   "message",
		EncodeLevel:  zapcore.CapitalColorLevelEncoder,                   // 使用自动颜色编码器
		EncodeTime:   zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"), // 格式化日期
		EncodeCaller: zapcore.ShortCallerEncoder,
	}

	// 控制台和文件编码器
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
	fileEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	// 文件输出目标
	fileWriter, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	// 控制台和文件输出
	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel), // 控制台输出（带颜色）
		zapcore.NewCore(fileEncoder, zapcore.AddSync(fileWriter), zapcore.InfoLevel),    // 文件输出（无颜色）
	)

	// 创建 Logger 实例
	Logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	zap.ReplaceGlobals(Logger)
}
