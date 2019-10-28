package build

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetAllProblems(ctx *Context, url string) (map[int]string, error) {
	resp, err := ctx.Client.R().
		SetHeader("content-type", "application/json").
		SetHeader("user-agent", _userAgent).
		Get(url)

	if err != nil {
		return nil, ErrWrap(err)
	}

	data := fmt.Sprintf("%v", resp)

	// for debug
	//if err = SaveFile(data); err != nil {
	//	return nil, err
	//}

	probResp := new(ProbResp)
	err = json.Unmarshal([]byte(data), probResp)
	if err != nil {
		return nil, ErrWrap(err)
	}

	idSlugMap := make(map[int]string)

	for _, sp := range probResp.StatStatusPairs {
		idSlugMap[sp.Stat.QuestionId] = sp.Stat.QuestionTitleSlug
	}

	return idSlugMap, nil
}

func QuerySlugs(idSlugMap map[int]string, ids []int) (result []string) {
	for _, id := range ids {
		if slug, ok := idSlugMap[id]; ok {
			result = append(result, slug)
		}
	}
	return
}

func GetDetailBySlug(ctx *Context, slug string) (*GraphQlResp, error) {
	detailUrl := GetDetailUrl(slug)

	resp, err := ctx.Client.R().
		SetCookie(&http.Cookie{Name: _csrfKey, Value: ctx.CSRF}).
		SetHeader("content-type", "application/json").
		SetHeader("referer", detailUrl).
		SetHeader("user-agent", _userAgent).
		SetHeader("x-csrftoken", ctx.CSRF).
		SetBody(NewGraphqlParma(slug)).
		Post(_grapgql)

	if err != nil {
		return nil, ErrWrap(err)
	}

	if resp.StatusCode() != 200 {
		return nil, Errorf("query failed, status code: %v", resp.StatusCode())
	}

	data := fmt.Sprintf("%v", resp)
	respObj := new(GraphQlResp)
	err = json.Unmarshal([]byte(data), respObj)
	if err != nil {
		return nil, Errorf(err.Error())
	}
	// 解析 html
	respObj.Data.Question.Content = GetHtmlContent(respObj.Data.Question.Content, detailUrl)

	return respObj, nil
}
