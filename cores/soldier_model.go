package cores

type Solider struct {
	Rank       string `json:"rank"`
	Wife       int    `json:"wife"`
	Salary     int    `json:"salary"`
	Home       bool   `json:"home"`
	Car        bool   `json:"car"`
	Corruption bool   `json:"corruption"`
}
