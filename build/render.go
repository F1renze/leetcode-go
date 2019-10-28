package build

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"strconv"
)

type RenderInfo struct {
	Details []*DetailInfo
	Total   int
}

type DetailInfo struct {
	Id         int
	Title      string
	Difficulty string
	Path       string
}

func NewRenderInfo(total int, details []*DetailInfo) *RenderInfo {
	return &RenderInfo{
		Details: details,
		Total:   total,
	}
}

func NewDetailInfo(id int, title, difficulty, path string) *DetailInfo {
	return &DetailInfo{
		Id:         id,
		Title:      title,
		Difficulty: difficulty,
		Path:       GetGitUrl(path),
	}
}

func RenderDetail(resp *GraphQlResp, solutionInfo *SolutionInfo) *DetailInfo {
	slug := resp.Data.Question.TitleSlug

	t := template.New(slug)
	data, _ := ioutil.ReadFile(_detailTemplate)

	var b bytes.Buffer

	t.Parse(string(data))
	t.Execute(&b, resp.Data.Question)

	id, _ := strconv.Atoi(resp.Data.Question.QuestionId)

	path := fmt.Sprintf("solutions/%v/%v.md", solutionInfo.FolderMap[id], slug)
	ioutil.WriteFile(path, b.Bytes(), 0755)

	fmt.Println("rendered: ", path)

	// path 指向具体 solution 文件
	return NewDetailInfo(id,
		resp.Data.Question.Title,
		resp.Data.Question.Difficulty,
		solutionInfo.SolutionMap[id])
}

func Render(meta *RenderInfo) {
	t := template.New("README")

	data, _ := ioutil.ReadFile(_readMeTemplate)

	var b bytes.Buffer

	t.Parse(string(data))
	t.Execute(&b, meta)
	ioutil.WriteFile("README.md", b.Bytes(), 0755)
}
