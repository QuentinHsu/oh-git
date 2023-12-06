package logger

import (
	"github.com/fatih/color"
)

// Logger 定义了一个 logger 结构体
type Logger struct{}

// Info 打印一条信息日志，使用默认的颜色样式
func (l *Logger) Info(message string, config ...color.Attribute) {
	printColoredMessage(color.FgWhite, message, config...)
}

// Success 打印一条成功日志，使用绿色的颜色样式
func (l *Logger) Success(message string, config ...color.Attribute) {
	printColoredMessage(color.FgGreen, message, config...)
}

// Warn 打印一条警告日志，使用黄色的颜色样式
func (l *Logger) Warn(message string, config ...color.Attribute) {
	printColoredMessage(color.FgYellow, message, config...)
}

// Error 打印一条错误日志，使用红色的颜色样式
func (l *Logger) Error(message string, config ...color.Attribute) {
	printColoredMessage(color.FgRed, message, config...)
}

// Label 打印一条 Label 日志，使用绿色的颜色样式
func (l *Logger) Label(message string, config ...color.Attribute) {
	printColoredMessage(color.FgGreen, message, config...)
}

// Value 打印一条 Value 日志，使用绿色的颜色样式
func (l *Logger) Value(message string, config ...color.Attribute) {
	printColoredMessage(color.FgYellow, message, config...)
}

// printColoredMessage 打印带有颜色样式的信息
func printColoredMessage(colorAttribute color.Attribute, message string, config ...color.Attribute) {
	c := color.New(colorAttribute)
	c.Add(config...)
	c.Print(message)
}
