package aaa9

import (
	"github.com/qiangxue/go-rest-api/internal/entity"
)

var (
	Name_rule1        = "rule1"
	Description_rule1 = "This is rule1"
)

func When_rule1(sta *entity.STA) bool {
	return sta.Property1 == "value1"
}

func Then_rule1(sta *entity.STA) *entity.Finding {
	// execute then for rule1

	return &entity.Finding{
		FindingType: entity.Success,
		Desc:        "executed then for aaa rule1",
		Status:      true,
	}
}
