package lib

import (
	"fmt"
	"testing"
)

// TestTemplateParse ensures Excel file can be opened
func TestTemplateParse(t *testing.T) {
	// Open a simple file and check matches expectations
	agent, err := NewAgent(Config{
		File:          "../examples/basic.xlsx",
		TemplateSheet: "Template",
	})
	defer agent.CloseExcel()
	templates, err := agent.TemplateParse()
	if err != nil {
		t.Errorf("error processing basic template %v", err)
	}
	numberOfTemplatesParsed := len(templates)
	if numberOfTemplatesParsed != 2 {
		t.Errorf("expected 2 templates, got %v", numberOfTemplatesParsed)
	}
	//log.Println(fmt.Sprint(templates))
	want := map[string]Template{
		"Simple": {
			Identifier: "Simple",
			Content: [][]string{
				{"A", "1", "E"},
				{"B", "2", "F"},
				{"C", "3", "G"},
				{"D", "4", "H"},
			},
		},
		"Substitution": {
			Identifier: "Substitution",
			Content: [][]string{
				{"A", "{{ .TempValue }}"},
				{"B", "{{ .TempValue }}"},
				{"C", "{{ .TempValue }}"},
				{"D", "{{ .TempValue }}"},
			},
		},
	}

	if fmt.Sprint(want) != fmt.Sprint(templates) {
		t.Error("basic template not parsed as expected")
	}
}

func TestInvalidTemplateSheet(t *testing.T) {
	// Open a simple file and check matches expectations
	agent, err := NewAgent(Config{
		File:          "../examples/basic.xlsx",
		TemplateSheet: "Invalid",
	})
	defer agent.CloseExcel()
	_, err = agent.TemplateParse()
	if err.Error() != "sheet Invalid does not exist" {
		t.Errorf("expected sheet not defined error, got: %v", err)
	}

	// Test fixing template sheet name works
	agent.TemplateSheet = "Template"
	_, err = agent.TemplateParse()
	if err != nil {
		t.Errorf("unexpected error parsing template, got: %v", err)
	}
}

func TestNoTemplateName(t *testing.T) {
	// Open a simple file and check matches expectations
	agent, err := NewAgent(Config{
		File:          "../examples/empty-template.xlsx",
		TemplateSheet: "Template",
	})
	defer agent.CloseExcel()
	_, err = agent.TemplateParse()
	if err == nil {
		t.Error("Expected error parsing templates")
		return
	}
	if err.Error() != "no template identified but parsing values" {
		t.Errorf("expected template not identified error, got: %v", err)
	}
}
