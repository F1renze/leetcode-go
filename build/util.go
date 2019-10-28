package build

import (
	"bufio"
	"fmt"
	"html"
	"log"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
)

func errorf(format string, a ...interface{}) error {
	return errors.Wrap(fmt.Errorf(format, a...), "")
}

/**
跟踪栈堆
*/
func errWrap(err error, a ...interface{}) error {
	var m string
	if len(a) > 0 {
		if format, ok := a[0].(string); ok {
			m = fmt.Sprintf(format, a[1:])
		}

	}
	return errors.Wrap(err, m)
}

/**
解析 html
*/
func GetHtmlContent(s, url string) string {
	body := html.UnescapeString(s)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}

	text := "refer: " + url + "\n\n" + doc.Text()

	return text
}

func SaveFile(data string) error {
	t := time.Now()
	date := fmt.Sprintf("%d%02d%02d", t.Year(), t.Month(), t.Day())
	name := fmt.Sprintf("build/file_%v.json", date)
	file, err := os.Create(name)
	if err != nil {
		return errorf(err.Error())
	}
	defer file.Close()
	w := bufio.NewWriter(file)

	_, err = w.WriteString(data)
	w.Flush()

	if err != nil {
		return errorf(err.Error())
	}
	return nil
}

func GetGitPrefix() string {
	// TODO read git
	//https://github.com/F1renze/leetcode-go/blob/master/solutions/q104/maximumdepth.go
	return "https://github.com/F1renze/leetcode-go/blob/dev/"
}
