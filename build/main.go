package build

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type SolutionInfo struct {
	FolderMap   map[int]string
	SolutionMap map[int]string
}

func (s *SolutionInfo) SetFolder(id int, folder string) {
	s.FolderMap[id] = folder
}

func (s *SolutionInfo) GetFolder(id int) (string, bool) {
	folder, ok := s.FolderMap[id]
	return folder, ok
}

func (s *SolutionInfo) SetSolution(id int, file string) {
	s.SolutionMap[id] = file
}

func (s *SolutionInfo) GetSolution(id int) (string, bool) {
	file, ok := s.SolutionMap[id]
	return file, ok
}

func NewSolutionInfo() *SolutionInfo {
	return &SolutionInfo{
		FolderMap:   make(map[int]string),
		SolutionMap: make(map[int]string),
	}
}

/**
返回题号对应的文件夹
*/
func IterSolutions() (*SolutionInfo, error) {

	files, err := ioutil.ReadDir(_solutionPath)
	if err != nil {
		return nil, ErrWrap(err)
	}

	info := NewSolutionInfo()

	for _, f := range files {
		if strings.HasPrefix(f.Name(), ".") {
			continue
		}

		name := strings.Split(f.Name(), "q")[1]
		id, err := strconv.Atoi(name)
		if err != nil {
			return nil, ErrWrap(err)
		}
		sfFiles, err := ioutil.ReadDir(fmt.Sprintf(
			_solutionPath+"/%v",
			f.Name()))

		if err != nil {
			return nil, Errorf(err.Error())
		}

		for _, sf := range sfFiles {
			if strings.HasSuffix(sf.Name(), ".go") &&
				!strings.HasSuffix(sf.Name(), "_test.go") {

				info.SetFolder(id, f.Name())
				info.SetSolution(id, fmt.Sprintf(
					_solutionPath+"/%v/%v", f.Name(), sf.Name()))
			}
		}
	}

	return info, nil
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

	solutionInfo, err := IterSolutions()
	if err != nil {
		return err
	}
	var (
		ids  []int
		meta []*DetailInfo
	)

	for id := range solutionInfo.FolderMap {
		ids = append(ids, id)
	}

	slugs := QuerySlugs(idSlugMap, ids)

	for _, s := range slugs {
		resp, err := GetDetailBySlug(ctx, s)
		if err != nil {
			fmt.Println(err)
		}
		r := RenderDetail(resp, solutionInfo)
		meta = append(meta, r)
	}

	Render(NewRenderInfo(len(meta), meta))

	return nil
}
