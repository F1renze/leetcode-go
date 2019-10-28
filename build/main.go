package build

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

/**
返回题号对应的文件夹
*/
func IterSolutions() (map[int]string, error) {
	files, err := ioutil.ReadDir("./solutions")
	if err != nil {
		return nil, errWrap(err)
	}

	idMap := make(map[int]string)

	for _, f := range files {
		if strings.HasPrefix(f.Name(), ".") {
			continue
		}

		name := strings.Split(f.Name(), "q")[1]
		id, err := strconv.Atoi(name)
		if err != nil {
			return nil, errWrap(err)
		}

		idMap[id] = f.Name()
	}

	return idMap, nil
}

func Run() error {

	ctx := NewContext()

	err := Login(ctx, _loginUrl)

	if err != nil {
		return err
	}

	idSlugMap, err := GetAllProblems(ctx, _problemsUrl)
	if err != nil {
		return err
	}

	idMaps, err := IterSolutions()
	if err != nil {
		return err
	}
	var (
		ids  []int
		meta []*DetailMeta
	)
	for k := range idMaps {
		ids = append(ids, k)
	}
	slugs := QuerySlugs(idSlugMap, ids)

	for _, s := range slugs {
		resp, err := GetDetailBySlug(ctx, s)
		if err != nil {
			fmt.Println(err)
		}
		r := RenderDetail(resp, idMaps)
		meta = append(meta, r)
	}

	Render(NewMeta(len(meta), meta))

	return nil
}
