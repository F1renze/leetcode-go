package build

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"strconv"
)

type RenderMeta struct {
	Details []*DetailMeta
	Total   int
}

type DetailMeta struct {
	Id         int
	Title      string
	Difficulty string
	Path       string
}

func NewMeta(total int, details []*DetailMeta) *RenderMeta {
	return &RenderMeta{
		Details: details,
		Total:   total,
	}
}

func NewDetailMeta(id int, title, difficulty, path string) *DetailMeta {
	return &DetailMeta{
		Id:         id,
		Title:      title,
		Difficulty: difficulty,
		Path:       GetGitPrefix() + path,
	}
}

func RenderDetail(resp *GraphQlResp, idMap map[int]string) *DetailMeta {
	slug := resp.Data.Question.TitleSlug

	t := template.New(slug)
	data, _ := ioutil.ReadFile(_detailTemplate)

	var b bytes.Buffer

	t.Parse(string(data))
	t.Execute(&b, resp.Data.Question)

	id, _ := strconv.Atoi(resp.Data.Question.QuestionId)

	path := fmt.Sprintf("solutions/%v/%v.md", idMap[id], slug)
	ioutil.WriteFile(path, b.Bytes(), 0755)

	fmt.Println("rendered: ", path)

	return NewDetailMeta(id,
		resp.Data.Question.Title,
		resp.Data.Question.Difficulty, path)
}

func Render(meta *RenderMeta) {
	t := template.New("README")

	data, _ := ioutil.ReadFile(_readMeTemplate)

	var b bytes.Buffer

	t.Parse(string(data))
	t.Execute(&b, meta)
	ioutil.WriteFile("README.md", b.Bytes(), 0755)
}
