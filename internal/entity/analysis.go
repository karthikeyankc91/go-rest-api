package entity

import "time"

type Analysis struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	StaId       string             `json:"staId"`
	FindingsMap map[string]Finding `json:"findingsMap"`
}

type Finding struct {
	RuleName     string        `json:"ruleName"`
	WhenStatus   bool          `json:"whenStatus"`
	RuleFindings []RuleFinding `json:"ruleFindings"`
}

type FindingType string

const (
	Unknown FindingType = "Unknown"
	Warning FindingType = "Warning"
	Failure FindingType = "Failure"
	Success FindingType = "Success"
)

type RuleFinding struct {
	FindingType FindingType `json:"findingType"`
	Desc        string      `json:"desc"`
	Status      bool        `json:"status"`
}

func CreateSuccessFinding(name string, desc string) *Finding {
	return &Finding{
		WhenStatus: true,
	}
}

func CreateFailureFinding(name string, desc string) *Finding {
	return &Finding{
		WhenStatus: false,
	}
}
