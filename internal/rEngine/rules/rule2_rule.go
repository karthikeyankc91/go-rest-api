package rules

import (
	"github.com/qiangxue/go-rest-api/internal/entity"
)

var (
	Name_rule2        = "rule2"
	Description_rule2 = "This is rule1"
)

func (r *RuleContext) When_rule2(sta *entity.STA) bool {
	return sta.Property2 == "value2"
}

func (r *RuleContext) Then_rule2(sta *entity.STA) *entity.RuleFinding {
	// execute then for rule1

	return &entity.RuleFinding{
		FindingType: entity.Success,
		Desc:        "executed then for rule2",
		Status:      true,
	}
}
