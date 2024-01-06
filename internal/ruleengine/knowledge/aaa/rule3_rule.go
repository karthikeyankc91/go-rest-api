package aaa

import (
	"github.com/qiangxue/go-rest-api/internal/entity"
)

var (
	Name_rule3        = "rule3"
	Description_rule3 = "This is rule3"
)

func When_rule3(sta *entity.STA) bool {
	return sta.Id != ""
}

func Then_rule3(sta *entity.STA) *entity.Finding {
	// execute then for rule1

	return &entity.Finding{
		FindingType: entity.Success,
		Desc:        "executed then for aaa rule3",
		Status:      true,
	}
}
