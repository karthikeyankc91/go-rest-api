package entity

import (
	"encoding/json"
	"fmt"
	"time"
)

type AnalysisData struct {
	ID        string    `json:"id"`
	Staid     string    `json:"staid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Knowledgemap *map[string]Knowledge `json:"knowledgeMap"`
}

type Analysis struct {
	ID        string    `json:"id"`
	Staid     string    `json:"staid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Knowledgemap json.RawMessage `json:"knowledgeMap"`
}

type Knowledge struct {
	KnowledgeName   string                  `json:"knowledgeName"`
	RuleFindingsMap *map[string]RuleFinding `json:"ruleFindingsMap"`
}

type RuleFinding struct {
	RuleName   string     `json:"ruleName"`
	WhenStatus bool       `json:"whenStatus"`
	Findings   *[]Finding `json:"findings"`
}

type Finding struct {
	FindingType FindingType `json:"findingType"`
	Desc        string      `json:"desc"`
	Status      bool        `json:"status"`
}

type FindingType string

const (
	Success FindingType = "Success"
	Failure FindingType = "Failure"
	Warning FindingType = "Warning"
	Unknown FindingType = "Unknown"
)

func (a *AnalysisData) Init(sta *STA) {
	a.ID = GenerateID()
	a.Staid = sta.Id
	now := time.Now()
	a.CreatedAt = now
	a.UpdatedAt = now
	newMap := make(map[string]Knowledge)
	a.Knowledgemap = &newMap
}

func (a *AnalysisData) AddKnowledge(knowledgeName string) (*Knowledge, error) {
	knowledge := (*(a.Knowledgemap))[knowledgeName]

	//knowledge already added case
	if knowledge.KnowledgeName != "" {
		return &knowledge, nil
	}
	knowledge = Knowledge{}
	knowledge.KnowledgeName = knowledgeName
	newMap := make(map[string]RuleFinding)
	knowledge.RuleFindingsMap = &newMap
	(*(a.Knowledgemap))[knowledgeName] = knowledge
	return &knowledge, nil
}

func (a *AnalysisData) GetKnowledge(knowledgeName string) (*Knowledge, error) {
	knowledge, ok := (*(a.Knowledgemap))[knowledgeName]
	if !ok {
		return nil, fmt.Errorf("knowledge %s not found", knowledgeName)
	}
	return &knowledge, nil
}

func (k *Knowledge) AddRuleFinding(ruleName string, whenStatus bool) (*RuleFinding, error) {
	ruleFinding := (*(k.RuleFindingsMap))[ruleName]
	if ruleFinding.RuleName != "" {
		return nil, fmt.Errorf("rule finding %s already exists", ruleName)
	}
	ruleFinding.RuleName = ruleName
	ruleFinding.WhenStatus = whenStatus
	newArray := make([]Finding, 0)
	ruleFinding.Findings = &newArray

	(*(k.RuleFindingsMap))[ruleName] = ruleFinding
	return &ruleFinding, nil
}

func (k *Knowledge) GetRuleFinding(ruleName string) (*RuleFinding, error) {
	finding, ok := (*k.RuleFindingsMap)[ruleName]
	if !ok {
		return nil, fmt.Errorf("rule finding %s not found", ruleName)
	}
	return &finding, nil
}

func (f *RuleFinding) AddFinding(finding *Finding) {
	// f.Findings = append(f.Findings, *finding)
}

func (analysis Analysis) UnmarshalAnalysis() (AnalysisData, error) {
	byteArr, err := analysis.Knowledgemap.MarshalJSON()
	if err != nil {
		return AnalysisData{}, err
	}

	m := make(map[string]Knowledge)
	err = json.Unmarshal(byteArr, &m)
	if err != nil {
		return AnalysisData{}, err
	}

	return AnalysisData{
		ID:           analysis.ID,
		Staid:        analysis.Staid,
		CreatedAt:    analysis.CreatedAt,
		UpdatedAt:    analysis.UpdatedAt,
		Knowledgemap: &m,
	}, nil
}

func (analysisData AnalysisData) MarshalAnalysis() (Analysis, error) {
	byteArray, err := json.Marshal(analysisData.Knowledgemap)
	if err != nil {
		return Analysis{}, err
	}
	return Analysis{
		ID:           analysisData.ID,
		Staid:        analysisData.Staid,
		CreatedAt:    analysisData.CreatedAt,
		UpdatedAt:    analysisData.UpdatedAt,
		Knowledgemap: json.RawMessage(byteArray),
	}, nil
}
