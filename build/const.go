package build

import (
	"fmt"
	"github.com/spf13/viper"
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
	_solutionPath = "./solutions"
)

var (
	_username string
	_password string
)

func init() {
	viper.SetConfigFile("build/conf.yml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	_username = viper.GetString("username")
	_password = viper.GetString("password")
}
