package ruleengine

import (
	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/internal/ruleengine/knowledge/aaa"
	"github.com/qiangxue/go-rest-api/internal/ruleengine/knowledge/bbb"
)

type RuleContext struct {
}

func (b *RuleContext) ExecWhen(knowledgeName string, ruleName string, ruleWhenResult bool, analysis *entity.Analysis) bool {
	knowledge, _ := analysis.AddKnowledge(knowledgeName)
	knowledge.AddRuleFinding(ruleName, ruleWhenResult)
	return ruleWhenResult
}

func (b *RuleContext) ExecThen(knowledgeName string, ruleName string, finding *entity.Finding, analysis *entity.Analysis) {
	knowledge, _ := analysis.GetKnowledge(knowledgeName)
	ruleFinding, _ := knowledge.GetRuleFinding(ruleName)
	findings := ruleFinding.Findings
	*findings = append(*findings, *finding)
}

func (r *RuleContext) When_aaa_rule1(sta *entity.STA) bool {
	return aaa.When_rule1(sta)
}

func (r *RuleContext) Then_aaa_rule1(sta *entity.STA) *entity.Finding {
	return aaa.Then_rule1(sta)
}

func (r *RuleContext) When_aaa_rule2(sta *entity.STA) bool {
	return aaa.When_rule2(sta)
}

func (r *RuleContext) Then_aaa_rule2(sta *entity.STA) *entity.Finding {
	return aaa.Then_rule2(sta)
}

func (r *RuleContext) When_bbb_rule1(sta *entity.STA) bool {
	return bbb.When_rule1(sta)
}

func (r *RuleContext) Then_bbb_rule1(sta *entity.STA) *entity.Finding {
	return bbb.Then_rule1(sta)
}

func (r *RuleContext) When_bbb_rule2(sta *entity.STA) bool {
	return bbb.When_rule2(sta)
}

func (r *RuleContext) Then_bbb_rule2(sta *entity.STA) *entity.Finding {
	return bbb.Then_rule2(sta)
}
