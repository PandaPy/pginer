package component

import (
	"github.com/fatih/color"
)

// TextFormatter 结构体，支持链式调用
type TextFormatter struct {
	text  string
	style *color.Color
}

// Text 函数初始化 TextFormatter
func Text(text string) *TextFormatter {
	return &TextFormatter{
		text:  text,
		style: color.New(),
	}
}

// Color 方法设置文本颜色
func (tf *TextFormatter) Color(attr color.Attribute) *TextFormatter {
	tf.style.Add(attr)
	return tf
}

// Bold 方法设置加粗样式
func (tf *TextFormatter) Bold() *TextFormatter {
	tf.style.Add(color.Bold)
	return tf
}

// String 方法返回格式化后的字符串
func (tf *TextFormatter) String() string {
	return tf.style.Sprint(tf.text)
}

func StyleText(text string, colorAttr color.Attribute, bold bool) string {
	formatter := Text(text).Color(colorAttr)
	if bold {
		formatter.Bold()
	}
	return formatter.String()
}
