package knowledge1

import (
	"github.com/qiangxue/go-rest-api/internal/entity"
)

var (
	Name_rule1        = "rule1"
	Description_rule1 = "This is rule1"
)

func When_rule1(sta *entity.STA) bool {
	return sta.StaId != ""
}

func Then_rule1(sta *entity.STA) *entity.Finding {
	// execute then for rule1

	return &entity.Finding{
		FindingType: entity.Success,
		Desc:        "executed then for knowledge1 rule1",
		Status:      true,
	}
}
