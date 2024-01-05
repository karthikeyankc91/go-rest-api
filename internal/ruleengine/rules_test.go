package ruleengine

import (
	"path"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/qiangxue/go-rest-api/pkg/log"
	"github.com/stretchr/testify/assert"
)

var (
	_, f, _, _      = runtime.Caller(0)
	testRulesFolder = path.Join(filepath.Dir(f), "../test/knowledge/")
	logger          = log.New()
)

func TestCreateRules(t *testing.T) {
	rulesMap, err := CreateRules(logger, testRulesFolder, "knowledge1")
	if err != nil {
		t.Errorf("Failed to create rules: %v", err)
	}

	assert.NotNil(t, rulesMap)
	assert.Equal(t, 2, len(*rulesMap))

	assert.Equal(t, "rule1", (*rulesMap)["rule1"].RuleName)
	assert.Equal(t, "rule2", (*rulesMap)["rule2"].RuleName)

	assert.Equal(t, "knowledge1", (*rulesMap)["rule1"].KnowledgeName)
	assert.Equal(t, "knowledge1", (*rulesMap)["rule2"].KnowledgeName)
}

func TestRuleToString(t *testing.T) {
	rule := &Rule{
		RuleName: "rule1",
		Desc:     "Rule description",
		RuleStr:  "Rule string",
	}

	_, err := rule.toString()
	assert.Nil(t, err)
}
