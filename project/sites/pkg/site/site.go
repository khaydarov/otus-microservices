package site

func NewSite(title string, domains []string) Site {
	return Site{
		NewID(),
		title,
		NewCode(),
		domains,
	}
}

type Site struct {
	ID      ID
	title   string
	code    Code
	domains []string
}
