package main

import (
	"fmt"
	"net/http"
	"strings"
	s "strings"
	"time"
)

/**
* For each post, checks to see if the post title includes
* the search term in some way, and if so, prints out the posts
* data.
* @param w http.ResponseWriter from handler, to write output to
* @param Data topData struct which contains all the posts received
* @param Term Search term given by the client.
**/
func output(w http.ResponseWriter, Data topData, Term string) {
	currentTime := time.Now().Local()
	postCount := 0
	fmt.Fprintf(w, "Front page of /r/"+subReddit+" as of: %s\n", currentTime.Format(time.RFC1123))
	fmt.Fprintf(w, "Search term: %s\n\n", Term)
	for _, v := range Data.Children {
		if s.Contains(strings.ToUpper(v.ChildData.Title), strings.ToUpper(Term)) {
			format(v.ChildData, w)
			postCount = postCount + 1
		}

	}
	if postCount == 0 {
		fmt.Fprintln(w, "Sorry, it appears your search yielded no results!\n Try again with a different term!")
	}
}

/**
 * Prints out the data for a single post.
 * @param Childdata struct containg the post data.
 * @param w http.ResponseWriter to print the data to.
 */
func format(Childdata data, w http.ResponseWriter) {
	fmt.Fprintln(w, Childdata.Title)
	fmt.Fprintf(w, "\tPost Score: %d\n", Childdata.Score)
	fmt.Fprintf(w, "\tPost Author: %s\n", Childdata.Author)
	fmt.Fprintf(w, "\tNumber of Comments: %d\n", Childdata.CommentNum)
	fmt.Fprintf(w, "\tDomain: %s\n", Childdata.Domain)
	fmt.Fprintf(w, "\tPermalink: %s\n", Childdata.Perma)
	fmt.Fprintf(w, "\tThumbnail: %s\n", Childdata.Thumb)
	fmt.Fprintf(w, "\tLink Flair: %s\n", Childdata.Flair)
	fmt.Fprintf(w, "\tNumber of Gildings: %d\n", Childdata.Gilded)
	fmt.Fprintf(w, "\tNumber of Crossposts: %d\n", Childdata.Crossposts)
	fmt.Fprintf(w, "\tNSFW?: %t\n", Childdata.NSFW)
	fmt.Fprintf(w, "\tStickied?: %t\n", Childdata.Stickied)
}
