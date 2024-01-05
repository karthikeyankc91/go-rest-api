package entity

import (
	"fmt"
	"time"
)

type Analysis struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	StaId        string                `json:"staId"`
	KnowledgeMap *map[string]Knowledge `json:"knowledgeMap"`
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

func (a *Analysis) Init(sta *STA) {
	a.ID = GenerateID()
	a.StaId = sta.StaId
	now := time.Now()
	a.CreatedAt = now
	a.UpdatedAt = now
	newMap := make(map[string]Knowledge)
	a.KnowledgeMap = &newMap
}

func (a *Analysis) AddKnowledge(knowledgeName string) (*Knowledge, error) {
	knowledge := (*(a.KnowledgeMap))[knowledgeName]

	//knowledge already added case
	if knowledge.KnowledgeName != "" {
		return &knowledge, nil
	}
	knowledge = Knowledge{}
	knowledge.KnowledgeName = knowledgeName
	newMap := make(map[string]RuleFinding)
	knowledge.RuleFindingsMap = &newMap
	(*(a.KnowledgeMap))[knowledgeName] = knowledge
	return &knowledge, nil
}

func (a *Analysis) GetKnowledge(knowledgeName string) (*Knowledge, error) {
	knowledge, ok := (*(a.KnowledgeMap))[knowledgeName]
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
