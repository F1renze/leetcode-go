package build

import "github.com/go-resty/resty/v2"

type Context struct {
	// 存放关键 cookie
	Kvs             map[string]string
	Client          *resty.Client
	CSRF            string
	LeetCodeSession string
	Referer         string
}

func (c *Context) Set(key, val string) {
	switch {
	case key == _csrfKey:
		c.CSRF = val
	case key == _sessionKey:
		c.LeetCodeSession = val
	}

	c.Kvs[key] = val
}

func (c *Context) Get(key string) (string, bool) {
	val, ok := c.Kvs[key]
	return val, ok
}

func NewContext() *Context {
	client := resty.New()
	//client.SetDebug(true)
	client.SetRedirectPolicy(resty.NoRedirectPolicy())
	return &Context{
		Client:  client,
		Kvs:     make(map[string]string),
		Referer: _referer,
	}
}
