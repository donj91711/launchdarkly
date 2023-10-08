package dao

type Score struct {
	StudentID string  `json:"studentId"`
	Exam      int     `json:"exam"`
	Score     float64 `json:"score"`
}

var ResultMap map[string]Score

func init() {
	ResultMap = make(map[string]Score)
}
