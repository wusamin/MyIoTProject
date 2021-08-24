package structs

type NobyResponse struct {
	CommandId    string     `json:"commandId"`
	CommandName  string     `json:"commandName"`
	Text         string     `json:"text"`
	Type         string     `json:"type"`
	Mood         float64    `json:"mood"`
	Negaposi     float64    `json:"negaposi"`
	NegaposiList []Negaposi `json:"negaposiList"`
	Emotion      Emotion    `json:"emotion"`
	Word         Word       `json:"word"`
	EmotionList  []Emotion  `json:"emotionList"`
	WordList     []Word     `json:"wordList"`
	Art          string     `json:"art"`
	Org          string     `json:"org"`
	Psn          string     `json:"psn"`
	Loc          string     `json:"loc"`
	Dat          string     `json:"dat"`
	Tim          string     `json:"tim"`
}

type Negaposi struct {
	Word  string  `json:"word"`
	Score float64 `json:"score"`
}

type Emotion struct {
	Word        string  `json:"word"`
	LikeDislike float64 `json:"likeDislike"`
	JoySad      float64 `json:"joySad"`
	AngerFear   float64 `json:"angerFear"`
}

type Word struct {
	Feature string `json:"feature"`
	Start   string `json:"start"`
	Surface string `json:"surface"`
}
