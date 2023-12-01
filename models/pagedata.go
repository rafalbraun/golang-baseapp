package models

// PageData holds the default data needed for HTML pages to render
type PageData struct {
	Title           string
	Messages        []Message
	IsAuthenticated bool
	IsAdmin         bool
	CacheParameter  string
	Trans           func(s string) string
	LoggedIn        *User
	Search          string
	Pagination      Pagination
	Language        string
	BaseURL         string
	Roles           []SystemRole
	Users           []User
}
