package lib

import (
	"fmt"
	"testing"
)

// TestTemplateParse ensures Excel file can be opened
func TestAgentState(t *testing.T) {
	// Open a simple file and check matches expectations
	agent, err := NewAgent(Config{
		File: "../examples/basic.xlsx",
	})
	defer agent.CloseExcel()
	if err != nil {
		t.Error("Cannot open excel file")
	}

	agent.SetTemplate(Template{
		Identifier: "test",
		Content:    nil,
	})
	newTemplate, err := agent.GetTemplate("test")
	if err != nil {
		t.Error("Error retrieving dynamically created template")
	}
	if newTemplate.Identifier != "test" {
		t.Error("Error retrieving saved template")
	}

	agent.AddItem(GeneratorItem{
		Template:    "test",
		OutputSheet: "test",
		Variables: map[string]string{
			"test": "test",
		},
	})

	want := []GeneratorItem{
		{
			Template:    "test",
			OutputSheet: "test",
			Variables: map[string]string{
				"test": "test",
			},
		},
	}

	if fmt.Sprint(want) != fmt.Sprint(agent.Generators) {
		t.Error("items not saved as expected")
	}
}
