package task

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"

	"github.com/fghwett/leetcodedaily/util"
)

type Task struct {
	client    *http.Client
	titleSlug string
	result    []string
}

func New() *Task {
	return &Task{
		client: &http.Client{},
		result: []string{"==== LeeCode 每日一题 ===="},
	}
}

func (t *Task) Do() {
	if err := t.getQuestionOfToday(); err != nil {
		t.result = append(t.result, fmt.Sprintf("【获取标题】：失败 %s", err))
		return
	}

	if err := t.getQuestionData(); err != nil {
		t.result = append(t.result, fmt.Sprintf("【获取内容】：失败 %s", err))
		return
	}

}

func (t *Task) getQuestionOfToday() error {
	reqUrl := "https://leetcode-cn.com/graphql"
	reqData := "{\n    \"operationName\": \"questionOfToday\",\n    \"variables\": {},\n    \"query\": \"query questionOfToday { todayRecord {   question {     questionFrontendId     questionTitleSlug     __typename   }   lastSubmission {     id     __typename   }   date   userStatus   __typename }} \"\n}"

	req, err := http.NewRequest(http.MethodPost, reqUrl, bytes.NewReader([]byte(reqData)))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Connection", "keep-alive")

	resp, err := t.client.Do(req)

	response := &QuestionOfTodayResponse{}
	err = util.GetHTTPResponse(resp, reqUrl, err, response)
	if err != nil {
		return err
	}

	t.titleSlug = response.Data.TodayRecord[0].Question.QuestionTitleSlug

	return nil
}

func (t *Task) getQuestionData() error {
	reqUrl := "https://leetcode-cn.com/graphql"
	reqData := fmt.Sprintf("{\n    \"operationName\": \"questionData\",\n    \"variables\": {\n        \"titleSlug\": \"%s\"\n    },\n    \"query\": \"query questionData($titleSlug: String!) {  question(titleSlug: $titleSlug) {    questionId    questionFrontendId    boundTopicId    title    titleSlug    content    translatedTitle    translatedContent    isPaidOnly    difficulty                                            likes    dislikes    isLiked    similarQuestions    contributors {      username      profileUrl      avatarUrl      __typename    }    langToValidPlayground    topicTags {      name      slug      translatedName      __typename    }    companyTagStats    codeSnippets {      lang      langSlug      code      __typename    }    stats    hints    solution {      id      canSeeDetail      __typename    }    status    sampleTestCase    metaData    judgerAvailable    judgeType    mysqlSchemas    enableRunCode    envInfo    book {      id      bookName      pressName      source      shortDescription      fullDescription      bookImgUrl      pressImgUrl      productUrl      __typename    }    isSubscribed    isDailyQuestion    dailyRecordStatus    editorType    ugcQuestionId    style    __typename  }}\"\n}", t.titleSlug)

	req, err := http.NewRequest(http.MethodPost, reqUrl, bytes.NewReader([]byte(reqData)))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Connection", "keep-alive")

	resp, err := t.client.Do(req)

	response := &QuestionDataResponse{}
	err = util.GetHTTPResponse(resp, reqUrl, err, response)
	if err != nil {
		return err
	}

	question := response.Data.Question
	title := fmt.Sprintf("【获取标题】：%s.%s(%s)", question.QuestionFrontendId, question.TranslatedTitle, question.Difficulty)
	content := fmt.Sprintf("【获取内容】：https://leetcode-cn.com/problems/%s", t.titleSlug)

	t.result = append(t.result, title, content)

	return nil
}

func (t *Task) GetResult() string {
	return strings.Join(t.result, " \n\n ")
}
