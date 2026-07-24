

import (
	"fmt"
)

func Log(s string) func() any {
	return func() any {
		fmt.Println(s)
		return nil
	}
}

func Warn(s string) func() any {
	return func() any {
		fmt.Println("[WARN]", s)
		return nil
	}
}

func Error(s string) func() any {
	return func() any {
		fmt.Println("[ERROR]", s)
		return nil
	}
}

func Info(s string) func() any {
	return func() any {
		fmt.Println("[INFO]", s)
		return nil
	}
}

func Debug(s string) func() any {
	return Log(s)
}

func Time(s string) func() any {
	return Log(s)
}

func TimeLog(s string) func() any {
	return Log(s)
}

func TimeEnd(s string) func() any {
	return Log(s)
}

func Clear() {
}

func Group(s string) func() any {
	return Log(s)
}

func GroupCollapsed(s string) func() any {
	return Log(s)
}

func GroupEnd() {
	Clear()
}
