package ruleengine

import (
	"os"

	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"github.com/qiangxue/go-rest-api/pkg/log"
)

type Knowledge struct {
	KnowledgeName   string
	KnowlodgeBase   *ast.KnowledgeBase
	RulesMap        *RulesMap
	KnowledgeString string
}

type Knowledgemap map[string]*Knowledge

func CreateKnowledge(logger log.Logger, knowledgeFolder string) (*ast.KnowledgeBase, *Knowledgemap, error) {
	knowledgeMap := make(Knowledgemap)

	files, err := os.ReadDir(knowledgeFolder)
	if err != nil {
		return nil, nil, err
	}

	for _, file := range files {
		if !file.IsDir() {
			continue
		}
		knowledgeName := file.Name()
		rulesMap, err := CreateRules(logger, knowledgeFolder, knowledgeName)
		if err != nil {
			return nil, nil, err
		}
		knowledgeString := rulesMap.GetKnowledgeGrlString()

		knowledgeBase, err := createKnowledgeFromGrlString(logger, knowledgeString)
		if err != nil {
			return nil, nil, err
		}

		knowledge := &Knowledge{
			KnowledgeName:   knowledgeName,
			RulesMap:        rulesMap,
			KnowledgeString: knowledgeString,
			KnowlodgeBase:   knowledgeBase,
		}

		knowledgeMap[knowledgeName] = knowledge
	}

	// create all knowledge
	var allKnowledgeString string
	for knowledgeName := range knowledgeMap {
		allKnowledgeString += knowledgeMap[knowledgeName].KnowledgeString
	}
	allKnowledgeBase, err := createKnowledgeFromGrlString(logger, allKnowledgeString)
	if err != nil {
		return nil, nil, err
	}

	return allKnowledgeBase, &knowledgeMap, nil
}

func createKnowledgeFromGrlString(logger log.Logger, knowledgeString string) (*ast.KnowledgeBase, error) {
	// building knowledge
	lib := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(lib)
	err := ruleBuilder.BuildRuleFromResource("Test", "0.1.1", pkg.NewBytesResource([]byte(knowledgeString)))
	if err != nil {
		return nil, err
	}
	knowledgeBase, err := lib.NewKnowledgeBaseInstance("Test", "0.1.1")
	if err != nil {
		return nil, err
	}
	return knowledgeBase, nil
}
