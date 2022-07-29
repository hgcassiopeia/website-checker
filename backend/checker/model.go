package checker

type CheckerResponse struct {
	StatusCode   int             `json:"status_code"`
	FileName     string          `json:"file_name"`
	SitesStatus  []WebsiteStatus `json:"sites_status"`
	CounterUp    int             `json:"counter_up"`
	CounterDown  int             `json:"counter_down"`
	TotalCounter int             `json:"total_counter"`
	ProcessTime  int64           `json:"process_time"`
}

type WebsiteStatus struct {
	ID     int
	Link   string
	Status bool
}

type Link struct {
	Url string
}

type Counter struct {
	LinksUp   []string
	LinksDown []string
	Up        int
	Down      int
}
