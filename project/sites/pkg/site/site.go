package site

func NewSite(userID, title string, domains []string) Site {
	return Site{
		NewID(),
		userID,
		title,
		NewCode(),
		domains,
	}
}

type Site struct {
	ID      ID
	UserID  string
	Title   string
	Code    Code
	Domains []string
}
