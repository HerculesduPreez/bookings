package models

// Templatedata holds data sent from handlers to template
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string //cross site request fraud token
	Flash     string
	Warning   string
	Error     string
}
