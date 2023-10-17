package lib

import (
	"fmt"
	"testing"
)

// TestTemplateParse ensures Excel file can be opened
func TestSheetConfigParse(t *testing.T) {
	// Open a simple file and check matches expectations
	agent, err := NewAgent(Config{
		File:        "../examples/sheet-config.xlsx",
		ConfigSheet: "Sheet Config",
	})
	defer agent.CloseExcel()
	if err != nil {
		t.Error("Cannot open basic example")
	}
	sheetConfigs, err := agent.SheetConfigParse()
	if err != nil {
		t.Errorf("error processing sheet config %v", err)
	}

	numberOfSheetsParsed := len(sheetConfigs)
	if numberOfSheetsParsed != 4 {
		t.Errorf("expected 4 templates, got %v", numberOfSheetsParsed)
	}

	want := map[string]Sheet{
		"Template": {
			Name:     "Template",
			WriteRow: 11,
		},
		"Things To Generate": {
			Name:     "Things To Generate",
			WriteRow: 6,
		},
		"Sheet Config": {
			Name:     "Sheet Config",
			WriteRow: 3,
		},
		"Output": {
			Name:     "Output",
			WriteRow: 10,
		},
	}

	if fmt.Sprint(want) != fmt.Sprint(sheetConfigs) {
		t.Error("sheet state not parsed as expected")
	}
}

func TestInvalidSheetConfig(t *testing.T) {
	// Open a simple file and check matches expectations
	agent, err := NewAgent(Config{
		File:        "./test/invalid-sheet-config.xlsx",
		ConfigSheet: "Sheet Config",
	})
	defer agent.CloseExcel()
	_, err = agent.SheetConfigParse()
	if err == nil {
		t.Error("Expected error parsing sheet config")
		return
	}
	if err.Error() != "unable to convert starting row. given: invalid" {
		t.Errorf("expected error about converting starting row, got: %v", err)
	}
}
