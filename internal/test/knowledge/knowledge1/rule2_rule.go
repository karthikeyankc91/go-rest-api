package knowledge1

import (
	"github.com/qiangxue/go-rest-api/internal/entity"
)

var (
	Name_rule2        = "rule2"
	Description_rule2 = "This is rule1"
)

func When_rule2(sta *entity.STA) bool {
	return sta.Property2 == "value2"
}

func Then_rule2(sta *entity.STA) *entity.Finding {
	// execute then for rule1

	return &entity.Finding{
		FindingType: entity.Success,
		Desc:        "executed then for knowledge1 rule2",
		Status:      true,
	}
}
