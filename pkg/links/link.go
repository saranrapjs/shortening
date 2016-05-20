package links

import "net/url"

type Link struct {
	Slug string `json:"slug",db:"slug"`
	URL  string `json:"url",db:"url"`
}

func (l *Link) ToRedirect() string {
	u, _ := url.Parse(l.URL)
	if !u.IsAbs() {
		u.Scheme = "http"
	}
	return u.String()
}
