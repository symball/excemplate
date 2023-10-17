package lib

import (
	"fmt"
	"testing"
)

// TestTemplateParse ensures Excel file can be opened
func TestItemsParse(t *testing.T) {
	// Open a simple file and check matches expectations
	agent, err := NewAgent(Config{
		File:            "../examples/basic.xlsx",
		TemplateSheet:   "Template",
		GeneratorsSheet: "Things To Generate",
	})
	defer agent.CloseExcel()
	if err != nil {
		t.Error("Cannot open excel file")
	}
	_, _ = agent.TemplateParse()

	itemConfigs, err := agent.GeneratorParse()
	if err != nil {
		t.Errorf("error processing generator config %v", err)
	}

	numberOfItemsParsed := len(itemConfigs)
	if numberOfItemsParsed != 5 {
		t.Errorf("expected 5 templates, got %v", numberOfItemsParsed)
	}

	want := []GeneratorItem{
		{
			Template:    "Simple",
			OutputSheet: "Output",
			Variables:   map[string]string{},
		},
		{
			Template:    "Simple",
			OutputSheet: "Output",
			Variables:   map[string]string{},
		},
		{
			Template:    "Substitution",
			OutputSheet: "Output",
			Variables: map[string]string{
				"TempValue": "Hello",
			},
		},
		{
			Template:    "Substitution",
			OutputSheet: "Output",
			Variables: map[string]string{
				"TempValue": "World",
			},
		},
		{
			Template:    "Simple",
			OutputSheet: "Output",
			Variables:   map[string]string{},
		},
	}

	if fmt.Sprint(want) != fmt.Sprint(itemConfigs) {
		t.Error("items not parsed as expected")
	}
}

func TestItemsNonExistentTemplate(t *testing.T) {
	// Open a simple file and check matches expectations
	agent, err := NewAgent(Config{
		File:            "./test/invalid-generator-non-existent-template.xlsx",
		GeneratorsSheet: "Things To Generate",
	})
	defer agent.CloseExcel()
	if err != nil {
		t.Error("Cannot open excel file")
	}
	_, err = agent.GeneratorParse()
	if err == nil {
		t.Error("expected non-existent template error")
		return
	}
	if err.Error() != "template not found: invalid" {
		t.Errorf("unexpected error processing generator %v", err)
	}
}
func TestItemsNoTemplate(t *testing.T) {
	// Open a simple file and check matches expectations
	agent, err := NewAgent(Config{
		File:            "./test/invalid-generator-no-template.xlsx",
		GeneratorsSheet: "Things To Generate",
	})
	defer agent.CloseExcel()
	if err != nil {
		t.Error("Cannot open excel file")
	}
	_, err = agent.GeneratorParse()
	if err == nil {
		t.Error("expected template finder error")
		return
	}
	if err.Error() != "template name required for generator" {
		t.Errorf("unexpected error processing generator %v", err)
	}
}

func TestItemsNoOutput(t *testing.T) {
	// Open a simple file and check matches expectations
	agent, err := NewAgent(Config{
		File:            "./test/invalid-generator-no-output-sheet.xlsx",
		GeneratorsSheet: "Things To Generate",
	})
	defer agent.CloseExcel()
	if err != nil {
		t.Error("Cannot open excel file")
	}
	_, err = agent.GeneratorParse()
	if err == nil {
		t.Error("expected output sheet error")
		return
	}
	if err.Error() != "template not found: Simple" {
		t.Errorf("unexpected error processing generator %v", err)
	}
}
