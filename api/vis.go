package api

type Node struct {
	ID     string `json:"id"`
	Label  string `json:"label"`
	Type   string `json:"type"`
	Status string `json:"status"`
	IP     string `json:"ip"`
}

type Edge struct {
	From string `json:"from"`
	To   string `json:"to"`
}
