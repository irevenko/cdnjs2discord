package types

// StatsResponse - types for the https://api.cdnjs.com/stats
type StatsResponse struct {
	LibrariesNumber int `json:"libraries"`
}

// StatsResponse - types for the https://api.cdnjs.com/stats
type SpecificLibResponse struct {
	Name       string `json:"name"`
	LatestLink string `json:"latest"`
	Authors    []struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"authors"`
	AutoUpdate []struct {
		Source string `json:"source"`
		Target string `json:"target"`
	} `json:"autoupdate"`
	Description string   `json:"description"`
	FileName    string   `json:"filename"`
	HomePage    string   `json:"homepage"`
	KeyWords    []string `json:"keywords"`
	License     string   `json:"license"`
	Repository  struct {
		Type string `json:"type"`
		URL  string `json:"url"`
	} `json:"repository"`
	Version  string   `json:"version"`
	Author   string   `json:"author"`
	Versions []string `json:"versions"`
}
