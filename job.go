package salt

// JobsResponse ...
type JobsResponse struct {
	Jobs []map[string]Job `json:"return"`
}

// JobResponse ...
type JobResponse struct {
	Job []Job `json:"info"`
}

// ExecutionResponse ...
type ExecutionResponse struct {
	Job []Job `json:"return"`
}

// Job ...
type Job struct {
	ID         string            `json:"jid"`
	Function   string            `json:"Function"`
	Target     string            `json:"Target"`
	User       string            `json:"User"`
	StartTime  string            `json:"StartTime"`
	TargetType string            `json:"Target-Type"`
	Arguments  []string          `json:"Arguments"`
	Minions    []string          `json:"Minions"`
	Result     map[string]Result `json:"Result"`
}

// Result ...
type Result struct {
	Return struct {
		PID     int    `json:"pid"`
		Retcode int    `json:"retcode"`
		Stdout  string `json:"stdout"`
		Stderr  string `json:"stderr"`
	} `json:"return"`
}
