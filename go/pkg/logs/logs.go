// Copyright 2021 Mojo-lang.org
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logs

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type ZapLogger = zap.SugaredLogger

var sugar *ZapLogger
var logger *zap.Logger
var atomicLevel zap.AtomicLevel

type PanicInfo struct {
	Stack, Panic string
}

func init() {
	sugar = newZap(NewConfig())

	//SetPanicHandler(func(a interface{}, stack []byte) {
	//	info := PanicInfo{}
	//	info.Panic = fmt.Sprint(a)
	//	info.Stack = string(stack)
	//	bytes, _ := json.Marshal(info)
	//	Error(string(bytes))
	//})
}

func Logger() *ZapLogger {
	return sugar
}

func New(cfg *Config) *ZapLogger {
	return newZap(cfg)
}

func SetLevel(level string) {
	l := strings.ToLower(level)
	if l == "info" {
		atomicLevel.SetLevel(zap.InfoLevel)
	} else if l == "debug" {
		atomicLevel.SetLevel(zap.DebugLevel)
	} else if l == "error" {
		atomicLevel.SetLevel(zap.ErrorLevel)
	} else if l == "warn" {
		atomicLevel.SetLevel(zap.WarnLevel)
	} else if l == "panic" {
		atomicLevel.SetLevel(zap.PanicLevel)
	} else if l == "fatal" {
		atomicLevel.SetLevel(zap.FatalLevel)
	} else {
		atomicLevel.SetLevel(zap.InfoLevel)
	}
}

func initConsoleCore(encode string) zapcore.Core {
	encodeConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	var formatEncoder zapcore.Encoder
	enc := strings.ToLower(encode)
	if enc == "json" {
		formatEncoder = zapcore.NewJSONEncoder(encodeConfig)
	} else {
		formatEncoder = zapcore.NewConsoleEncoder(encodeConfig)
	}
	consoleDebugging := zapcore.Lock(os.Stdout)
	return zapcore.NewCore(formatEncoder, consoleDebugging, atomicLevel)
}

func initFileCore(encode, filename string, maxSize, maxBackups, maxAge int) zapcore.Core {
	encodeConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	var formatEncoder zapcore.Encoder
	enc := strings.ToLower(encode)
	if enc == "json" {
		formatEncoder = zapcore.NewJSONEncoder(encodeConfig)
	} else {
		formatEncoder = zapcore.NewConsoleEncoder(encodeConfig)
	}

	f := handleFileName(filename)
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   f,
		MaxSize:    maxSize, // megabytes
		MaxBackups: maxBackups,
		MaxAge:     maxAge, // days
		Compress:   true,
	})

	return zapcore.NewCore(formatEncoder, w, atomicLevel)
}

func handleFileName(filename string) string {
	filename = path.Clean(filename)
	parts := make([]string, 0)
	var ret string
	paths := strings.Split(filename, string(os.PathSeparator))
	for _, v := range paths {
		val := handleTemplateFileName(v)
		if len(val) > 0 {
			parts = append(parts, val)
		}
	}

	if path.IsAbs(filename) {
		ret = string(os.PathSeparator) + path.Join(parts...)
	} else {
		ret = path.Join(parts...)
	}
	return ret
}

func handleTemplateFileName(template string) string {
	// foo1{hostname}foo2{port}foo3
	lefts := make([]int, 0)
	rights := make([]int, 0)

	size := len(template)
	for i := 0; i < size; i++ {
		if template[i] == '{' {
			lefts = append(lefts, i)
		} else if template[i] == '}' {
			rights = append(rights, i)
		}
	}

	leftSize := len(lefts)
	rightSize := len(rights)
	var minSize int
	if leftSize < rightSize {
		minSize = leftSize
	} else {
		minSize = rightSize
	}

	ret := template
	for i := minSize - 1; i >= 0; i-- {
		variableName := ret[lefts[i]+1 : rights[i]]
		v := os.Getenv(variableName)
		ret = ret[:lefts[i]] + v + ret[rights[i]+1:]
	}
	return ret
}

// level string, encode string, port int, pattern string, initFields map[string]interface{}
func newZap(cfg *Config) *ZapLogger {
	var opts []zap.Option
	opts = append(opts, zap.Development())
	opts = append(opts, zap.AddCaller())
	opts = append(opts, zap.AddCallerSkip(1))
	opts = append(opts, zap.AddStacktrace(zap.FatalLevel))

	atomicLevel = zap.NewAtomicLevel()
	SetLevel(cfg.Level)

	if len(cfg.LevelPattern) > 0 && cfg.LevelPort > 0 {
		http.HandleFunc(cfg.LevelPattern, atomicLevel.ServeHTTP)
		go func() {
			fmt.Printf("level serve on port:%d\nusage: [GET] curl http://localhost:%d%s\nusage: [PUT] curl -XPUT --data '{\"level\":\"debug\"}' http://localhost:%d%s\n", cfg.LevelPort, cfg.LevelPort, cfg.LevelPattern, cfg.LevelPort, cfg.LevelPattern)
			if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.LevelPort), nil); err != nil {
				panic(err)
			}
		}()
	}

	cores := make([]zapcore.Core, 0)
	output := strings.ToLower(cfg.Output)
	if strings.Contains(output, "console") {
		cores = append(cores, initConsoleCore(cfg.Encode))
	}
	if strings.Contains(output, "file") {
		cores = append(cores, initFileCore(cfg.File.Encode, cfg.File.Path, cfg.File.MaxSize, cfg.File.MaxBackups, cfg.File.MaxAge))
	}

	core := zapcore.NewTee(cores...)

	initFields := cfg.InitFields
	logger = zap.New(core)
	if len(initFields) > 0 {
		initFieldList := make([]zap.Field, 0)
		for k, v := range initFields {
			initFieldList = append(initFieldList, zap.Any(k, v))
		}
		logger = logger.With(initFieldList...)
	}

	logger = logger.WithOptions(opts...)
	return logger.Sugar()
}

// Debug logs a message at DebugLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Debug(args ...interface{}) {
	Logger().Debug(args...)
}

// Info logs a message at InfoLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Info(args ...interface{}) {
	Logger().Info(args...)
}

// Warn logs a message at WarnLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Warn(args ...interface{}) {
	Logger().Warn(args...)
}

// Error logs a message at ErrorLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Error(args ...interface{}) {
	Logger().Error(args...)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func Fatal(args ...interface{}) {
	Logger().Fatal(args...)
}

// Debugf uses fmt.Sprintf to log a templated message.
func Debugf(template string, args ...interface{}) {
	Logger().Debugf(template, args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func Infof(template string, args ...interface{}) {
	Logger().Infof(template, args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func Warnf(template string, args ...interface{}) {
	Logger().Warnf(template, args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func Errorf(template string, args ...interface{}) {
	Logger().Errorf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	Logger().Fatalf(template, args...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	Logger().Debugw(msg, keysAndValues...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	Logger().Infow(msg, keysAndValues...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	Logger().Warnw(msg, keysAndValues...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	Logger().Errorw(msg, keysAndValues...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	Logger().Fatalw(msg, keysAndValues...)
}
