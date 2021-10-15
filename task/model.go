package task

type QuestionOfTodayResponse struct {
	Data struct {
		TodayRecord []struct {
			Question struct {
				QuestionFrontendId string `json:"questionFrontendId"`
				QuestionTitleSlug  string `json:"questionTitleSlug"`
				Typename           string `json:"__typename"`
			} `json:"question"`
			LastSubmission interface{} `json:"lastSubmission"`
			Date           string      `json:"date"`
			UserStatus     interface{} `json:"userStatus"`
			Typename       string      `json:"__typename"`
		} `json:"todayRecord"`
	} `json:"data"`
}

type QuestionDataResponse struct {
	Data struct {
		Question struct {
			QuestionFrontendId string `json:"questionFrontendId"`
			TranslatedTitle    string `json:"translatedTitle"`
			Difficulty         string `json:"difficulty"`
		} `json:"question"`
	} `json:"data"`
}
