package common

import (
	"fmt"

	"k8s.io/klog/v2"
)

const (
	color_red     = uint8(iota + 91)
	color_green   //	绿
	color_yellow  //	黄
	color_blue    // 	蓝
	color_magenta //	洋红
	color_blue2   // 	浅蓝
)

func Red(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_red, s)
}

func Green(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_green, s)
}

func Yellow(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_yellow, s)
}

func Blue(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_blue, s)
}

func Blue2(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_blue2, s)
}

// 洋红
func Magenta(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_magenta, s)
}

// klog 封装
func KLogErrorf(format string, args ...interface{}) {
	klog.Errorf(Red(fmt.Sprintf(format, args...)))
}

func KLogInfof(format string, args ...interface{}) {
	klog.Infof(Green(fmt.Sprintf(format, args...)))
}

func KLogWarningf(format string, args ...interface{}) {
	klog.Warningf(Yellow(fmt.Sprintf(format, args...)))
}

func KLogError(args ...interface{}) {
	klog.Error(Red(fmt.Sprint(args...)))
}

func KLogInfo(format string, args ...interface{}) {
	klog.Info(Green(fmt.Sprint(args...)))
}

func KLogWarning(format string, args ...interface{}) {
	klog.Warning(Yellow(fmt.Sprint(args...)))
}
