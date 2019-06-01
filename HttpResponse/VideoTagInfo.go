package HttpResponse

type VideoTagInfoResponse struct {
	Code    int                      `json:"code"`
	Message string                   `json:"message"`
	Ttl     int                      `json:"ttl"`
	Data    VideoTagInfoResponseData `json:"data"`
}

type VideoTagInfoResponseData struct {
	TagDetail []VideoTagDetail `json:"tag_detail"`
}

type VideoTagDetail struct {
	TagId       int                 `json:"tag_id"`
	TagName     string              `json:"tag_name"`
	Cover       string              `json:"cover"`
	HeadCover   string              `json:"head_cover"`
	Content     string              `json:"content"`
	ShotContent string              `json:"shot_content"`
	Type        int                 `json:"type"`
	State       int                 `json:"state"`
	Ctime       int                 `json:"ctime"`
	Count       VideoTagDetailCount `json:"count"`
	IsAtten     int                 `json:"is_atten"`
	Likes       int                 `json:"likes"`
	Hates       int                 `json:"hates"`
	Attribute   int                 `json:"attribute"`
	Liked       int                 `json:"liked"`
	Hated       int                 `json:"hated"`
}

type VideoTagDetailCount struct {
	View  int `json:"view"`
	Use   int `json:"use"`
	Atten int `json:"atten"`
}
