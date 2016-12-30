package config

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/bughou-go/xiaomei/utils"
)

const ISO8601 = `2006-01-02T15:04:05Z0700`

func Log(msg string) {
	println(time.Now().Format(ISO8601), msg)
}

func IsLocalEnv() bool {
	goenv := os.Getenv(`GOENV`)
	return goenv == `` || goenv == App.Env()
}

func Debug(name string) bool {
	matched, _ := regexp.MatchString(`\b`+name+`\b`, os.Getenv(`debug`))
	return matched
}

func Protect(fn func()) {
	defer func() {
		err := recover()
		if err != nil {
			errMsg := fmt.Sprintf("PANIC: %s\n%s", err, utils.Stack(4))
			App.Alarm(`Protect错误`, errMsg)
			log.Printf(errMsg)
		}
	}()
	fn()
}
