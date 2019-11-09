package applogger

import (
	"fmt"
	"sync"
	"time"
)

const (
	black   = iota
	red     = iota
	green   = iota
	yellow  = iota
	blue    = iota
	magenta = iota
	cyan    = iota
	white   = iota

	reset = "\u001b[0m"

	logoTemplate = "<> %s"
	textTemplate = "\t %s - %s "

	bold      = "\u001b[1m"
	underline = "\u001b[4m"
	reversed  = "\u001b[7m"
)

type message struct {
	backgroundColor, textColor, textStyle string
}

func (m *message) printTemplate() string {
	return m.backgroundColor + m.textColor + m.textStyle
}
func (m *message) setTextStyle(style string) {
	m.textStyle = style
}
func (m *message) setTextColor(color int) {
	m.textColor = fmt.Sprintf("\u001b[3%dm", color)
}
func (m *message) setBeckgroundColor(color int) {
	m.backgroundColor = fmt.Sprintf("\u001b[4%dm", color)
}

type appLog struct {
	infoLogoTemplate, infoTextTemplate         message
	allertLogoTemplate, allertTextTemplate     message
	criticalLogoTemplate, criticalTextTemplate message
	fatalLogoTemplate, fatalTextTemplate       message
	mux                                        sync.Mutex
}

func (a *appLog) printInfo() string {
	s := a.infoLogoTemplate.printTemplate() + logoTemplate + reset + a.infoTextTemplate.printTemplate() + textTemplate + reset + "\n"
	return s
}
func (a *appLog) printAllert() string {
	s := a.allertLogoTemplate.printTemplate() + logoTemplate + reset + a.allertTextTemplate.printTemplate() + textTemplate + reset + "\n"
	return s
}
func (a *appLog) printCritical() string {
	s := a.criticalLogoTemplate.printTemplate() + logoTemplate + reset + a.criticalTextTemplate.printTemplate() + textTemplate + reset + "\n"
	return s
}
func (a *appLog) printFatal() string {
	s := a.fatalLogoTemplate.printTemplate() + logoTemplate + reset + a.fatalTextTemplate.printTemplate() + textTemplate + reset + "\n"
	return s
}

var appLogger appLog

func init() {
	appLogger.infoLogoTemplate.setTextColor(black)
	appLogger.infoLogoTemplate.setBeckgroundColor(green)
	appLogger.infoLogoTemplate.setTextStyle(bold)
	appLogger.infoTextTemplate.setTextColor(green)
	appLogger.infoTextTemplate.setBeckgroundColor(black)
	appLogger.infoTextTemplate.setTextStyle("")
}

// Info - log info message
func Info(msg string) {
	go printData(msg, 0)
}

// Allert - log info message
func Allert(msg string) {
	go printData(msg, 1)
}

// Critical - log info message
func Critical(msg string) {
	go printData(msg, 2)
}

// Fatal - log info message
func Fatal(msg string) {
	go printData(msg, 3)
}

func printData(msg string, msgType int) {
	defer appLogger.mux.Unlock()
	appLogger.mux.Lock()
	switch msgType {
	case 0:
		fmt.Printf(appLogger.printInfo(), "[INFO]", getTime(), msg)
	case 1:
		fmt.Printf(appLogger.printInfo(), "[ALLERT]", getTime(), msg)
	case 2:
		fmt.Printf(appLogger.printInfo(), "[CRITICAL]", getTime(), msg)
	case 3:
		fmt.Printf(appLogger.printInfo(), "[FATAL]", getTime(), msg)
	}
}
func getTime() string {
	return time.Now().Format("[2006-01-22] 15:04:05.999999 ")
}
