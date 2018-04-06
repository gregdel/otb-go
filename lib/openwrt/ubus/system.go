package ubus

// Board represents the system board informations
type Board struct {
	Kernel    string  `json:"kernel"`
	Hostname  string  `json:"hostname"`
	System    string  `json:"system"`
	Model     string  `json:"model"`
	BoardName string  `json:"board_name"`
	Release   Release `json:"release"`
}

// Release represents the system release informations
type Release struct {
	Distribution string `json:"distribution"`
	Version      string `json:"version"`
	Revision     string `json:"revision"`
	CodeName     string `json:"codename"`
	Target       string `json:"target"`
	Description  string `json:"description"`
}

// SystemBoard returns the system board info
func SystemBoard() (*Board, error) {
	board := Board{}

	if err := call(&board, "system", "board"); err != nil {
		return nil, err
	}

	return &board, nil
}
