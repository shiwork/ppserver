package ppserver

type Movie struct {
	Url string `json:"url"`
	Title string `json:"title"`
	Subtitle string `json:"subtitle,omitempty"`
	Chapter	int `json:"chapter,omitempty"`
	Category string `json:"category,omitempty"`
	RecordedAt string `json:"recorded_at,omitempty"`
}

type Channel struct {
	Movies []Movie `json:"movies"`
}
