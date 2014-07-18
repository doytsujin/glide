package cmd

import (
	"github.com/Masterminds/cookoo"
	"fmt"
)

var Quiet bool = false
const (
	Blue    = "0;34"
	Red     = "0;31"
	BoldRed = "1;31"
	Yellow  = "0;33"
	Cyan    = "0;36"
	Pink    = "1;35"
)

func Color(code, msg string) string {
	return fmt.Sprintf("\033[%sm%s\033[m", code, msg)
}


func BeQuiet(c cookoo.Context, p *cookoo.Params) (interface{}, cookoo.Interrupt) {
	qstr := p.Get("quiet", "false").(string)
	Quiet = qstr == "true"
	return Quiet, nil
}

func Info(msg string, args ...interface{}) {
	if Quiet { return }
	fmt.Print(Color(Yellow, "[INFO] "))
	Msg(msg, args...)
}
func Debug(msg string, args ...interface{}) {
	if Quiet { return }
	fmt.Print("[DEBUG] ")
	Msg(msg, args...)
}

func Warn(msg string, args ...interface{}) {
	fmt.Print(Color(Red, "[WARN] "))
	Msg(msg, args...)
}

func Error(msg string, args ...interface{}) {
	fmt.Print(Color(BoldRed, "[ERROR] "))
	Msg(msg, args...)
}

func Msg(msg string, args...interface{}) {
	if len(args) == 0 {
		fmt.Print(msg)
		return
	}
	fmt.Printf(msg, args...)
}
