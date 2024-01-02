package rules

import (
	"github.com/qiangxue/go-rest-api/internal/entity"
)

func (r *RuleContext) When_rule1(sta *entity.STA) bool {
	return sta.Property1 == "value1"
}

func (r *RuleContext) Then_rule1(sta *entity.STA) *entity.Finding {
	finding := &entity.Finding{
		Name:   "rule1",
		Desc:   "executed rule1",
		Status: true,
	}

	return finding
}
