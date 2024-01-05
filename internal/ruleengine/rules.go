package ruleengine

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path"
	"strings"

	"github.com/qiangxue/go-rest-api/pkg/log"
)

type Rule struct {
	RuleName      string
	Desc          string
	KnowledgeName string
	RuleStr       string
	// When     func(*entity.STA) bool
	// Then     func(*entity.STA) *entity.RuleFinding
}

type RulesMap map[string]*Rule

func CreateRules(logger log.Logger, knowledgeFolder string, knowledgeName string) (*RulesMap, error) {
	rulesMap := make(RulesMap)

	files, err := os.ReadDir(path.Join(knowledgeFolder, knowledgeName))
	if err != nil {
		return nil, err
	}
	if len(files) == 0 {
		return nil, fmt.Errorf("no rules found in %s", knowledgeName)
	}

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), "_rule.go") {
			continue
		}
		ruleName := strings.TrimSuffix(file.Name(), "_rule.go")
		rule := &Rule{
			RuleName:      ruleName,
			Desc:          "desc",
			KnowledgeName: knowledgeName,
		}
		rule.RuleStr, err = rule.toString()
		if err != nil {
			return nil, err
		}
		rulesMap[ruleName] = rule
	}
	return &rulesMap, nil
}

func (rulesMap *RulesMap) GetKnowledgeGrlString() string {
	var rulesStr string
	for ruleName := range *rulesMap {
		rulesStr += (*rulesMap)[ruleName].RuleStr
	}
	return rulesStr
}

func (rulesMap *RulesMap) GetKnowledgeGrlStringWithFilter(ruleNames []string) (string, error) {
	var rulesStr string
	for _, ruleName := range ruleNames {
		if rule, ok := (*rulesMap)[ruleName]; ok {
			rulesStr += rule.RuleStr
		} else {
			return "", fmt.Errorf("rule %s not found", ruleName)
		}
	}
	return rulesStr, nil
}

func (rule *Rule) toString() (string, error) {
	tmpl, err := template.New("rule").Parse(`
    rule {{.KnowledgeName}}_{{.RuleName}} "{{.Desc}}" salience 10 {
        when
            ctx.ExecWhen("{{.KnowledgeName}}", "{{.RuleName}}", ctx.When_{{.KnowledgeName}}_{{.RuleName}}(sta), analysis)
        then
			ctx.ExecThen("{{.KnowledgeName}}", "{{.RuleName}}", ctx.Then_{{.KnowledgeName}}_{{.RuleName}}(sta), analysis);
            Retract("{{.KnowledgeName}}_{{.RuleName}}");
    }
`)
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, rule); err != nil {
		return "", err
	}

	return tpl.String(), nil
}
