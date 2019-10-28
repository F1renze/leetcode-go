package build

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"os/exec"
	"strings"
)

const (
	_csrfKey    = "csrftoken"
	_sessionKey = "LEETCODE_SESSION"

	_problemsUrl = "https://leetcode.com/api/problems/Algorithms/"
	_loginUrl    = "https://leetcode.com/accounts/login/"
	_referer     = "https://leetcode.com/"
	_grapgql     = "https://leetcode.com/graphql"

	_userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.120 Safari/537.36"

	_detailTemplate = "build/detail_template.md"
	_readMeTemplate = "build/template.md"
	_solutionPath   = "solutions"
)

var (
	_username string
	_password string
	_branch string
)

func init() {
	viper.SetConfigFile("build/conf.yml")

	var (
		err error
		buf bytes.Buffer
	)
	defer func() {
		if err != nil {
			panic(err)
		}
	}()

	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	_username = viper.GetString("username")
	_password = viper.GetString("password")


	c1 := exec.Command("git", "branch")
	c2 := exec.Command("grep", "-i", `\*`)

	c2.Stdin, err = c1.StdoutPipe()
	c2.Stdout = &buf
	err = c2.Start()
	err = c1.Run()
	err = c2.Wait()

	data := string(buf.Bytes())
	// output example: * dev
	_branch = strings.TrimSpace(strings.Split(data, "*")[1])
	fmt.Println("current branch: ", _branch)
}
