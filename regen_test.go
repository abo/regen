package regen

import (
	"reflect"
	"regexp"
	"strings"
	"testing"
)

func TestNginx(t *testing.T) {
	nginx := `209.160.24.63 - - [15/May/2015:18:22:16] "GET /oldlink?itemId=EST-6&JSESSIONID=SD0SL6FF7ADFF4953 HTTP 1.1" 200 1748 "http://www.buttercupgames.com/oldlink?itemId=EST-6" "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.5 (KHTML, like Gecko) Chrome/19.0.1084.46 Safari/536.5" 731
209.160.24.63 - - [15/May/2015:18:22:17] "GET /product.screen?productId=BS-AG-G09&JSESSIONID=SD0SL6FF7ADFF4953 HTTP 1.1" 200 2550 "http://www.buttercupgames.com/product.screen?productId=BS-AG-G09" "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.5 (KHTML, like Gecko) Chrome/19.0.1084.46 Safari/536.5" 422
209.160.24.63 - - [15/May/2015:18:22:19] "POST /category.screen?categoryId=STRATEGY&JSESSIONID=SD0SL6FF7ADFF4953 HTTP 1.1" 200 407 "http://www.buttercupgames.com/cart.do?action=remove&itemId=EST-7&productId=PZ-SG-G05" "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.5 (KHTML, like Gecko) Chrome/19.0.1084.46 Safari/536.5" 211
209.160.24.63 - - [15/May/2015:18:22:20] "GET /product.screen?productId=FS-SG-G03&JSESSIONID=SD0SL6FF7ADFF4953 HTTP 1.1" 200 2047 "http://www.buttercupgames.com/category.screen?categoryId=STRATEGY" "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.5 (KHTML, like Gecko) Chrome/19.0.1084.46 Safari/536.5" 487`
	lines := strings.Split(nginx, "\n")

	learnAndExtract(t, lines, "209.160.24.63", []string{"209.160.24.63", "209.160.24.63", "209.160.24.63", "209.160.24.63"})
	learnAndExtract(t, lines, "15/May/2015:18:22:16", []string{"15/May/2015:18:22:16", "15/May/2015:18:22:17", "15/May/2015:18:22:19", "15/May/2015:18:22:20"})
	learnAndExtract(t, lines, "[15/May/2015:18:22:16]", []string{"[15/May/2015:18:22:16]", "[15/May/2015:18:22:17]", "[15/May/2015:18:22:19]", "[15/May/2015:18:22:20]"})
	learnAndExtract(t, lines, "18:22:16", []string{"18:22:16", "18:22:17", "18:22:19", "18:22:20"})
	learnAndExtract(t, lines, "GET", []string{"GET", "GET", "POST", "GET"})

}

func TestMailsv(t *testing.T) {
	mailsv := `Thu May 15 2015 00:15:05 mailsv1 sshd[2716]: Failed password for invalid user postgres from 86.212.199.60 port 4093 ssh2
Thu May 15 2015 00:15:05 mailsv1 sshd[2596]: Failed password for invalid user whois from 86.212.199.60 port 3311 ssh2
Thu May 15 2015 00:15:05 mailsv1 sshd[24947]: pam_unix(sshd:session): session opened for user djohnson by (uid=0)
Thu May 15 2015 00:15:05 mailsv1 sshd[3006]: Failed password for invalid user info from 86.212.199.60 port 4078 ssh2
Thu May 15 2015 00:15:05 mailsv1 sshd[5298]: Failed password for invalid user postgres from 86.212.199.60 port 1265 ssh2
Thu May 15 2015 00:15:05 mailsv1 sshd[5196]: Failed password for invalid user irc from 86.212.199.60 port 1454 ssh2
Thu May 15 2015 00:15:05 mailsv1 sshd[4472]: Failed password for invalid user vpxuser from 86.212.199.60 port 4203 ssh2
Thu May 15 2015 00:15:05 mailsv1 sshd[63551]: pam_unix(sshd:session): session opened for user djohnson by (uid=0)
Thu May 15 2015 00:15:05 mailsv1 sshd[5237]: Failed password for surly from 86.212.199.60 port 3734 ssh2
Thu May 15 2015 00:15:05 mailsv1 sshd[5737]: Failed password for invalid user mysql from 175.44.1.172 port 4073 ssh2
Thu May 15 2015 00:15:05 mailsv1 sshd[4508]: Failed password for invalid user services from 175.44.1.172 port 3288 ssh2
Thu May 15 2015 00:15:05 mailsv1 sshd[1254]: Failed password for invalid user testing from 175.44.1.172 port 1361 ssh2
Thu May 15 2015 00:15:05 mailsv1 sshd[46748]: Received disconnect from 10.3.10.46 11: disconnected by user
Thu May 15 2015 00:15:05 mailsv1 sshd[5730]: Failed password for invalid user admin from 175.44.1.172 port 4512 ssh2
Thu May 15 2015 00:15:05 mailsv1 sshd[3202]: Failed password for invalid user noone from 175.44.1.172 port 2394 ssh2
Thu May 15 2015 00:15:05 mailsv1 sshd[5555]: Failed password for invalid user noone from 175.44.1.172 port 2326 ssh2
Thu May 15 2015 00:15:05 mailsv1 sshd[1258]: Failed password for invalid user web002 from 175.44.1.172 port 4851 ssh2
Thu May 15 2015 00:15:05 mailsv1 sshd[12190]: pam_unix(sshd:session): session opened for user djohnson by (uid=0)`
	lines := strings.Split(mailsv, "\n")
	learnAndExtract(t, lines, "2716", []string{"2716",
		"2596", "24947", "3006", "5298", "5196", "4472", "63551", "5237", "5737", "4508", "1254", "46748", "5730", "3202", "5555", "1258", "12190"})
}

func learnAndExtract(t *testing.T, lines []string, field string, expect []string) {
	p, err := Generate(lines[0], field)
	if err != nil {
		t.Fatal("extract failed", err)
	}

	re := regexp.MustCompile(p)
	fields := make([]string, len(lines))
	for i, v := range lines {
		matches := re.FindStringSubmatch(v)
		if len(matches) >= 2 {
			fields[i] = matches[1]
		}
	}

	if !reflect.DeepEqual(fields, expect) {
		t.Fatal("expect: ", expect, " actual: ", fields)
	}
}
