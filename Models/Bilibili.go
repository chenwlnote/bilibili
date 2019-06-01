package Models

type Bilibili struct {
	Id          int    `json:"id"`
	Aid         int    `json:"aid"`
	Title       string `json:"title"`
	Url         string `json:"url"`
	Duration    int    `json:"duration"`
	View        int    `json:"view"`
	Danmaku     int    `json:"danmaku"`
	Reply       int    `json:"reply"`
	Favorite    int    `json:"favorite"`
	Coin        int    `json:"coin"`
	Share       int    `json:"share"`
	Like        int    `json:"like"`
	NowRank     int    `json:"now_rank"`
	HisRank     int    `json:"his_rank"`
	Keywords    string `json:"keywords"`
	ActionTag   string `json:"action_tag"`
	EmotionTag  string `json:"emotion_tag"`
	SceneTag    string `json:"scene_tag"`
	StarTag     string `json:"star_tag"`
	DialogTag   string `json:"dialog_tag"`
	UpdateCount int    `json:"update_count"`
	UpdatedAt   string `json:"updated_at"`
	CreatedAt   string `json:"created_at"`
}
