package build

import "fmt"

/**
存放 response 结构体
*/

type ProbResp struct {
	UserName        string       `json:"user_name"`
	NumSolved       int          `json:"num_solved"`
	NumTotal        int          `json:"num_total"`
	AcEasy          int          `json:"ac_easy"`
	AcMedium        int          `json:"ac_medium"`
	AcHard          int          `json:"ac_hard"`
	StatStatusPairs []StatusPair `json:"stat_status_pairs"`
	FrequencyHigh   int          `json:"frequency_high"`
	FrequencyMid    int          `json:"frequency_mid"`
	CategorySlug    string       `json:"category_slug"`
}

type StatusPair struct {
	Stat Stat `json:"stat"`
	// 题目是否通过
	Status     string     `json:"status"`
	Difficulty Difficulty `json:"difficulty"`
	PaidOnly   bool       `json:"paid_only"`
	IsFavor    bool       `json:"is_favor"`
	Frequency  int        `json:"frequency"`
	Progress   int        `json:"progress"`
}

type Difficulty struct {
	Level int `json:"level"`
}

type Stat struct {
	QuestionId        int    `json:"question_id"`
	QuestionTitle     string `json:"question__title"`
	QuestionTitleSlug string `json:"question__title_slug"`
}

type GraphQlResp struct {
	Data GraphQlRespData `json:"data"`
}

type GraphQlRespData struct {
	Question GraphQlQuestion `json:"question"`
}

type GraphQlQuestion struct {
	QuestionId string `json:"questionId"`
	Title      string `json:"title"`
	TitleSlug  string `json:"titleSlug"`
	Difficulty string `json:"difficulty"`
	Content    string `json:"content"`
	TopicTags  []Tag  `json:"topicTags"`
}

type Tag struct {
	Name string `json:"name"`
}

type GraphQlParam struct {
	OperationName string          `json:"operationName"`
	Query         string          `json:"query"`
	Variables     GraphQlParamVar `json:"variables"`
}

type GraphQlParamVar struct {
	TitleSlug string `json:"titleSlug"`
}

func NewGraphqlParma(slug string) GraphQlParam {
	return GraphQlParam{
		OperationName: "questionData",
		Query:         "query questionData($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    questionId\n    questionFrontendId\n    boundTopicId\n    title\n    titleSlug\n    content\n    translatedTitle\n    translatedContent\n    isPaidOnly\n    difficulty\n    likes\n    dislikes\n    isLiked\n    similarQuestions\n    contributors {\n      username\n      profileUrl\n      avatarUrl\n      __typename\n    }\n    langToValidPlayground\n    topicTags {\n      name\n      slug\n      translatedName\n      __typename\n    }\n    companyTagStats\n    codeSnippets {\n      lang\n      langSlug\n      code\n      __typename\n    }\n    stats\n    hints\n    solution {\n      id\n      canSeeDetail\n      __typename\n    }\n    status\n    sampleTestCase\n    metaData\n    judgerAvailable\n    judgeType\n    mysqlSchemas\n    enableRunCode\n    enableTestMode\n    envInfo\n    libraryUrl\n    __typename\n  }\n}\n",
		Variables: GraphQlParamVar{
			TitleSlug: slug,
		},
	}
}

func GetDetailUrl(slug string) string {
	return fmt.Sprintf("https://leetcode.com/problems/%v", slug)
}
