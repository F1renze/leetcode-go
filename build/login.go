package build

func Login(ctx *Context, url string) error {
	err := GetCSRF(ctx, url)

	if err != nil {
		return err
	}

	bodyMap := map[string]string{
		"login":               _username,
		"password":            _password,
		"csrfmiddlewaretoken": ctx.CSRF,
		"next":                "problems",
	}

	// ignore error: "Post problems: auto redirect is disabled"
	resp, _ := ctx.Client.R().
		SetFormData(bodyMap).
		//SetHeader().
		SetHeader("referer", "https://leetcode.com/accounts/login/").
		SetHeader("x-csrftoken", ctx.CSRF).
		SetHeader("user-agent", _userAgent).
		Post(url)

	for _, c := range resp.Cookies() {
		switch {
		case c.Name == _sessionKey:
			ctx.Set(_sessionKey, c.Value)
		// 更新 csrf token
		case c.Name == _csrfKey:
			ctx.Set(_csrfKey, c.Value)
		}
	}

	if _, ok := ctx.Get(_sessionKey); ok == false {
		return errorf("login failed, unable found session id in resp header")
	}

	if resp.StatusCode() != 302 {
		return errorf("login failed, status code %v", resp.StatusCode())
	}

	return nil
}

func GetCSRF(ctx *Context, url string) error {
	resp, err := ctx.Client.R().
		SetHeader("user-agent", _userAgent).
		Get(url)

	if err != nil {
		return errWrap(err)
	}

	for _, c := range resp.Cookies() {
		if c.Name == _csrfKey {
			ctx.Set(_csrfKey, c.Value)
			return nil
		}
	}

	return errorf("unable found csrf token in resp header")
}
