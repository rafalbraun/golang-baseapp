package models

// Message holds a message which can be rendered as responses on HTML pages
type Message struct {
	Type    string // success, warning, error, etc.
	Content string
}
