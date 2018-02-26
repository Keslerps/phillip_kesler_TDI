package main

type redditAPI struct {
	Kind string  `json:"kind"`
	Data topData `json:"data"`
}
type topData struct {
	Children []child `json:"children"`
}
type data struct {
	Domain     string `json:"domain"`
	Score      int    `json:"score"`
	Crossposts int    `json:"num_crossposts"`
	NSFW       bool   `json:"over_18"`
	Thumb      string `json:"thumbnail"`
	Flair      string `json:"link_flair_css_class"`
	Gilded     int    `json:"gilded"`
	Stickied   bool   `json:"stickied"`
	Title      string `json:"title"`
	Perma      string `json:"permalink"`
	CommentNum int    `json:"num_comments"`
}

type child struct {
	Kind      string `json:"kind"`
	ChildData data   `json:"data"`
}
