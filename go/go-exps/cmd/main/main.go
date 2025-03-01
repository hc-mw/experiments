package main

import (
	ApacheLogParser "apacheParser/pkg/parser"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

type Integer int

func bar() Integer {
	return Integer(5)
}
func main() {
	m := map[string]int{}

	fmt.Println(len(m))
}

func countTo(max int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < max; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func logDemo() {
	//l1 := `172.18.0.4 - - [12/Mar/2023:09:56:43 +0000] "GET /api/user HTTP/1.1" 200 234`
	l2 := `172.18.0.2 - - [08/May/2024:17:51:17 +0000] "GET / HTTP/1.1" 200 8888 "-" "curl/8.5.0" 159`
	// p1 := `(?P<ip>\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}) (?P<logname>\S+) (?P<remoteuser>\S+) \[(?P<time>.*?)\] "(?P<method>\w+) (?P<url>.*?) HTTP\/(?P<version>\d\.\d)" (?P<status>\d{3}) (?P<size>\S+)`
	// p2 := `(?P<ip>\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}) (?P<logname>\S+) (?P<remoteuser>\S+) \[(?P<time>.*?)\] "(?P<method>\w+) (?P<url>.*?) HTTP\/(?P<version>\d\.\d)" (?P<status>\d{3}) (?P<size>\S+) "(?P<referer>.*?)" "(?P<useragent>.*?)"`
	//p3 := `(?P<remote_hostname>\S+) (?P<remote_logname>\S+) (?P<remote_user>\S+) \[(?P<request_time>.*?)\] "(?P<method>\w+) (?P<url>.*?) HTTP\/(?P<version>\d\.\d)" (?P<status>\d+) (?P<response_size_clf>\d+|-) \"(?P<Referer>.*?)\" \"(?P<User-Agent>.*?)\"`
	//     (?P<remote_hostname>\S+) (?P<remote_logname>\S+) (?P<remote_user>\S+) \[(?P<request_time>.*?)\] "(?P<method>\w+) (?P<url>.*?) HTTP\/(?P<version>\d\.\d)" (?P<status>\d+) (?P<response_size_clf>\d+|-) "(?P<Referer>.*?)i" "(?P<User-Agent>.*?)i"
	// foo(p1, l1)
	// foo(p2, l2)
	a := "%h %l %u %t \\\"%r\\\" %>s %b \\\"%{Referer}i\\\" \\\"%{User-Agent}i\\\" %D"
	regex := ApacheLogParser.ParseLogFormat(a)
	fmt.Println(regex)
	sanitizedRegex := ApacheLogParser.SanitizeRegex(regex)
	fmt.Println(sanitizedRegex)
	sanitizedCustom := strings.ReplaceAll(sanitizedRegex, `\`, `\\`)
	fmt.Println(sanitizedCustom)

	fmt.Println(fmt.Sprintf(`body matches "%s"`, sanitizedCustom))

	foo(regex, l2)
	//str := strings.ReplaceAll("\"%r\" %d %r", "%r", `(?P<method>\w+) (?P<url>.*?) HTTP\/(?P<version>\d\.\d)`)
	//fmt.Println(str)
}

func foo(pattern, logEntry string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	r := regexp.MustCompile(pattern)
	match := r.FindStringSubmatch(logEntry)

	result := make(map[string]string)
	for i, name := range r.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}

	json, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(json))
}
