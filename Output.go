package main

import (
	"fmt"
	"net/http"
	"strings"
	s "strings"
)

func output(w http.ResponseWriter, Data topData, Term string) {
	for _, v := range Data.Children {
		if s.Contains(strings.ToUpper(v.ChildData.Title), strings.ToUpper(Term)) {
			format(v.ChildData, w)
		}

	}
}

func format(Childdata data, w http.ResponseWriter) {
	fmt.Fprintln(w, Childdata.Title)
	fmt.Fprintf(w, "\tDomain: %s\n", Childdata.Domain)
	fmt.Fprintf(w, "\tPost Score: %d\n", Childdata.Score)
	fmt.Fprintf(w, "\tNumber of Comments: %d\n", Childdata.CommentNum)
	fmt.Fprintf(w, "\tThumbnail: %s\n", Childdata.Thumb)
	fmt.Fprintf(w, "\tLink Flair: %s\n", Childdata.Flair)
	fmt.Fprintf(w, "\tNumber of Gildings: %d\n", Childdata.Gilded)
	fmt.Fprintf(w, "\tNumber of Crossposts: %d\n", Childdata.Crossposts)
	fmt.Fprintf(w, "\tNSFW?: %t\n", Childdata.NSFW)
	fmt.Fprintf(w, "\tStickied?: %t\n", Childdata.Stickied)
}
