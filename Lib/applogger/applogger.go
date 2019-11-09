package applogger

import (
	"fmt"
	"time"
)

const (
	reset = "\u001b[0m"

	black   = "0"
	red     = "1"
	green   = "2"
	yellow  = "3"
	blue    = "4"
	magenta = "5"
	cyan    = "6"
	white   = "7"

	textColor       = "\u001b[3"
	backgroundColor = "\u001b[4"

	bright = ";1m"
	dark   = "m"

	msgData = "<> %s"

	bold      = "\u001b[1m"
	underline = "\u001b[4m"
	reversed  = "\u001b[7m"

	supportData = "\t %s - %s "

	fatalTemplate    = backgroundColor + red + bright + textColor + white + bright + bold + msgData + reset + supportData
	criticalTemplate = backgroundColor + yellow + bright + textColor + black + dark + bold + msgData + reset + supportData
	allertTemplate   = backgroundColor + white + bright + textColor + black + dark + bold + msgData + reset + supportData
	infoTemplate     = backgroundColor + green + bright + textColor + black + dark + bold + msgData + reset + supportData
)

func Info(msg string) {
	info := fmt.Sprintf(infoTemplate, "[INFO]", getTime(), msg)
	printData(info)
}
func Allert(msg string) {
	info := fmt.Sprintf(allertTemplate, "[ALLERT]", getTime(), msg)
	printData(info)
}
func Critical(msg string) {
	info := fmt.Sprintf(criticalTemplate, "[CRITICAL]", getTime(), msg)
	printData(info)
}
func Fatal(msg string) {
	info := fmt.Sprintf(fatalTemplate, "[FATAL]", getTime(), msg)
	printData(info)
}
func printData(msg string) {
	fmt.Println(msg)
}
func getTime() string {
	return time.Now().Format("[2006-01-02] 15:04:05.999999 ")
}
