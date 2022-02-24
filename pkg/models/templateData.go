package models

// TemplateData holds data sent from handlers to template
type TemplateData struct {
	StringMap    map[string]string
	IntMap       map[int]int
	FloatMap     map[float32]float32
	Data         map[string]interface{}
	CSRFToken    string
	FlashMessage string
	Warning      string
	Error        string
}
