package main

/**
 * Top level data found in all Reddit_API.
 * topData will be the data more specific to the subreddit,
 * and Kind will always be "Listing"
 **/
type redditAPI struct {
	Kind string  `json:"kind"`
	Data topData `json:"data"`
}

/**
 * topData is just a slice of child structs,
 * representing all the posts on the front page.
 **/
type topData struct {
	Children []child `json:"children"`
}

/**
 * Data for a specific post. Data included consist of
 * domain: where the data is hosted,
 * Score: Current score of the post
 * Author: Reddit User responsible for the post
 * Crossposts: how many times this post has been x-posted
 * NSFW: If the post is marked as 18 only
 * thumbnail: link to the thumbnail image of the post
 * Flair: CSS specific link flair of the post, if any
 * GIlded: How many times the post has been gilded, if any
 * Stickied: If the post is stickied to the top of the subreddit or not
 * Title: title of the post
 * Perm: Permalink to the post
 * CommentNum: Number of comments
 * Note that these are presented in the order they appear in
 * Reddit's API
 **/
type data struct {
	Domain     string `json:"domain"`
	Score      int    `json:"score"`
	Author     string `json:"author"`
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

/**
 * Single post. Has post specific data, as well
 * as a kind field, which should always evaluate to
 * "t3"
 **/
type child struct {
	Kind      string `json:"kind"`
	ChildData data   `json:"data"`
}
