package ruleengine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testKnowledgeFolder = "../test/knowledge/"
)

func TestCreateKnowledge(t *testing.T) {
	allKnowledgeBase, knowledgeMap, err := CreateKnowledge(logger, testKnowledgeFolder)
	if err != nil {
		t.Errorf("Failed to create knowledge: %v", err)
	}

	assert.NotNil(t, allKnowledgeBase)

	assert.NotNil(t, knowledgeMap)
	assert.Equal(t, 2, len(*knowledgeMap))
	assert.Equal(t, "knowledge1", (*knowledgeMap)["knowledge1"].KnowledgeName)
	assert.Equal(t, "knowledge2", (*knowledgeMap)["knowledge2"].KnowledgeName)

	assert.Equal(t, 2, len(*((*knowledgeMap)["knowledge1"].RulesMap)))
	assert.Equal(t, "rule1", (*((*knowledgeMap)["knowledge1"].RulesMap))["rule1"].RuleName)
	assert.Equal(t, "rule2", (*((*knowledgeMap)["knowledge1"].RulesMap))["rule2"].RuleName)
	assert.Equal(t, "rule1", (*((*knowledgeMap)["knowledge2"].RulesMap))["rule1"].RuleName)
	assert.Equal(t, "rule2", (*((*knowledgeMap)["knowledge2"].RulesMap))["rule2"].RuleName)

	assert.Equal(t, "knowledge1", (*((*knowledgeMap)["knowledge1"].RulesMap))["rule1"].KnowledgeName)
	assert.Equal(t, "knowledge1", (*((*knowledgeMap)["knowledge1"].RulesMap))["rule2"].KnowledgeName)
	assert.Equal(t, "knowledge2", (*((*knowledgeMap)["knowledge2"].RulesMap))["rule1"].KnowledgeName)
	assert.Equal(t, "knowledge2", (*((*knowledgeMap)["knowledge2"].RulesMap))["rule2"].KnowledgeName)

}
