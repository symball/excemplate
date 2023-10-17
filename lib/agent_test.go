package lib

import (
	"testing"
)

// TestAgentOpen ensures Excel file can be opened
func TestAgentOpen(t *testing.T) {
	// Open a simple file and check matches expectations
	agent, err := NewAgent(Config{
		File: "../examples/basic.xlsx",
	})
	defer agent.CloseExcel()
	if err != nil {
		t.Error("Cannot open basic example")
	}
	workSpaceSheets := agent.Excel.GetSheetList()
	if len(workSpaceSheets) != 2 {
		t.Error("Excel improperly detected basic example sheet count")
	}

	// Open another file and check agent changes properly
	err = agent.OpenExcel("../examples/sheet-config.xlsx")
	if err != nil {
		t.Error("Cannot open sheet-config example")
	}
	workSpaceSheets = agent.Excel.GetSheetList()
	if len(workSpaceSheets) != 4 {
		t.Error("Excel improperly detected sheet-config example sheet count")
	}

	// Open a non existent file and check returns error as expected
	err = agent.OpenExcel("./non-existent.xlsx")
	if err.Error() != "could not find input file" {
		t.Error("Expected error opening non-existent file")
	}
}
