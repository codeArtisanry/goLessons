package main

import (
	"regexp"

	"github.com/lunny/log"
)

const (
	PATH_SEP = "/"

	LOG_TYPE_NGINX_500  = "nginx_500"
	LOG_TYPE_NGINX_404  = "nginx_404"
	LOG_TYPE_APACHE_500 = "apache_500"
	LOG_TYPE_APACHE_404 = "apache_404"
	LOG_TYPE_PHP_ERROR  = "php_error"
	LOG_TYPE_APP        = "app"

	LOG_TYPE_MEMBER_ACTIVITY        = "ma"
	LOG_TYPE_MEMBER_ACTIVITY_COUPON = "mac"
	LOG_TYPE_MEMBER_COUPON          = "mc"
)

var (
	//[05/13/15 20:26:09] [INFO] 127.0.0.1 - - [13/May/2015:20:26:09 +0800] "GET /javascript/ffan/common/json2.jsa HTTP/1.1" 404 8978 "-" "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:37.0) Gecko/20100101 Firefox/37.0"
	//nginxAccessLogReg  = regexp.MustCompile(`^([^ ]*) ([^ ]*) ([^ ]*) \[([^\]]*)\] "([^"]*)" ([^ ]*) ([^ ]*) "([^"]*)" "([^"]*)"$`)
	nginxAccessLogReg  = regexp.MustCompile(`^([^ ]*) ([^ ]*) ([^ ]*) \[([^\]]*)\] ([^ ]*) "([^"]*)" ([^ ]*) ([^ ]*) ([^ ]*) "([^"]*)" "([^"]*)" "([^"]*)" "([^"]*)"$`)
	apacheAccessLogReg = nginxAccessLogReg
	//[14-May-2015 08:29:10 Asia/Chongqing] PHP Notice:  Use of undefined constant a - assumed 'a' in /home/nicholaskh/wanda/ffan/web/index.php on line 2
	phpErrorReg = regexp.MustCompile(`^\[([^\]]*)\] ([^:]*): (.*)$`)
	parser      *Parser
)

type Parser struct {
	Regexps map[string]*regexp.Regexp
}

func NewParser() *Parser {
	parser := new(Parser)
	parser.Regexps = make(map[string]*regexp.Regexp)
	parser.Regexps[LOG_TYPE_NGINX_404] = nginxAccessLogReg
	parser.Regexps[LOG_TYPE_NGINX_500] = nginxAccessLogReg
	parser.Regexps[LOG_TYPE_APACHE_404] = apacheAccessLogReg
	parser.Regexps[LOG_TYPE_APACHE_500] = apacheAccessLogReg
	parser.Regexps[LOG_TYPE_PHP_ERROR] = phpErrorReg

	return parser
}

func (this *Parser) parse(txt, tp string) []string {
	re, exists := this.Regexps[tp]
	if !exists {
		log.Warn("regexp not found for %s", tp)
	}
	match := re.FindAllStringSubmatch(txt, -1)
	for _, r := range match {
		for _, rr := range r {
			log.Debug(rr)
		}
	}

	if len(match) < 1 || len(match[0]) < 2 {
		return []string{}
	}
	return match[0][1:]
}

func (this Parser) match(txt, tp string) bool {
	re, exists := this.Regexps[tp]
	if !exists {
		log.Warn("regexp not found for %s", tp)
	}
	return re.MatchString(txt)
}
