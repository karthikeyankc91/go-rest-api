package rules

import (
	"github.com/qiangxue/go-rest-api/internal/entity"
)

var (
	Name_rule1        = "rule1"
	Description_rule1 = "This is rule1"
)

func (r *RuleContext) When_rule1(sta *entity.STA) bool {
	return sta.Property1 == "value1" && sta.Property2 == "value2"
}

func (r *RuleContext) Then_rule1(sta *entity.STA) *entity.RuleFinding {
	// execute then for rule1

	return &entity.RuleFinding{
		FindingType: entity.Success,
		Desc:        "executed then for rule1",
		Status:      true,
	}
}
