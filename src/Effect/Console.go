package Effect_Console

import (
	"fmt"
)

func Log(s string, _ interface{}) interface{} {
	fmt.Println(s)
	return nil
}

func Warn(s string, _ interface{}) interface{} {
	fmt.Println("[WARN]", s)
	return nil
}

func Error(s string, _ interface{}) interface{} {
	fmt.Println("[ERROR]", s)
	return nil
}

func Info(s string, _ interface{}) interface{} {
	fmt.Println("[INFO]", s)
	return nil
}

func Debug(s string, _ interface{}) interface{} {
	return Log(s, nil)
}

func Time(s string, _ interface{}) interface{} {
	return Log(s, nil)
}

func TimeLog(s string, _ interface{}) interface{} {
	return Log(s, nil)
}

func TimeEnd(s string, _ interface{}) interface{} {
	return Log(s, nil)
}

func Clear(_ interface{}) interface{} {
	return nil
}

func Group(s string, _ interface{}) interface{} {
	return Log(s, nil)
}

func GroupCollapsed(s string, _ interface{}) interface{} {
	return Log(s, nil)
}

func GroupEnd(_ interface{}) interface{} {
	return nil
}
