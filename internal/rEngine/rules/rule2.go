package rules

import (
	"github.com/qiangxue/go-rest-api/internal/entity"
)

func (r *RuleContext) When_rule2(sta *entity.STA) bool {
	return sta.Property2 == "value1"
}

func (r *RuleContext) Then_rule2(sta *entity.STA) *entity.Finding {
	finding := &entity.Finding{
		Name:   "rule2",
		Desc:   "executed rule2",
		Status: false,
	}

	return finding
}
