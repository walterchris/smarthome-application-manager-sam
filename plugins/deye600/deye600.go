package deye600

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/walterchris/smarthome-application-manager-sam/pkg/communication"
	"github.com/walterchris/smarthome-application-manager-sam/pkg/loader"
	"github.com/walterchris/smarthome-application-manager-sam/plugins"
)

const name = "Deye600"

var (
	rChannel *communication.Channels
	log      *logrus.Logger
)

var values = map[string]string{
	"webdata_now_p":   "0",
	"webdata_today_e": "0",
	"webdata_total_e": "0",
}

type config struct {
	IP net.IP
}

type Deye600 struct{}

func init() {
	loader.LoadFunctions = append(loader.LoadFunctions, Load)
}

func Load(logger *logrus.Logger, channels communication.Channels) (plugins.Plugin, error) {
	log = logger
	log.Tracef("%s:\tLoad\n", name)

	rChannel = &channels

	return Deye600{}, nil
}

func (deye600 Deye600) Run() error {
	log.Tracef("%s:\tRun\n", name)

	for {

		// 192.168.0.192/status.html
		url := "http://192.168.0.192/status.html"
		fmt.Printf("HTML code of %s ...\n", url)
		resp, err := http.Get(url)
		// handle the error if there is one
		if err != nil {
			// Deye is offline - thus webdata_now_p is 0
			values["webdata_now_p"] = "0"
			sendMQTTMessages()

			time.Sleep(30 * time.Second)
			continue
		}

		// reads html as a slice of bytes
		html, err := io.ReadAll(resp.Body)
		if err != nil {
			resp.Body.Close()
			time.Sleep(60 * time.Second)
			continue
		}

		for keyword := range values {
			fmt.Println(keyword)
			if strings.Contains(string(html), keyword) {

				r := regexp.MustCompile(fmt.Sprintf(`var\ %s\ =\ \"(?P<now_p>[0-9]+|[-+]?([0-9]*\.[0-9]+|[0-9]+))\";`, keyword))
				matches := r.FindStringSubmatch(string(html))
				if len(matches) > 1 {
					values[keyword] = matches[1]
				}
			}
		}

		resp.Body.Close()
		time.Sleep(5 * time.Second)
	}

	return nil
}

func (deye600 Deye600) Name() string {
	return name
}

func StringToLines(s string) []string {
	var lines []string

	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return lines
}

func sendMQTTMessages() {
	// Send out all values
	for keyword := range values {
		msg := communication.Mqtt{
			Msg:   values[keyword],
			Topic: fmt.Sprintf("deye600/%s", keyword),
		}
		rChannel.Messages <- msg
	}
}
