package lib

import (
	"errors"
)

type Template struct {
	Identifier string
	Content    [][]string
}

func (a *Agent) TemplateParse() (map[string]Template, error) {

	if a.TemplateSheet == "" {
		return nil, errors.New("template sheet not defined")
	}

	TemplateSheet, err := a.Excel.GetRows(a.TemplateSheet)
	if err != nil {
		return nil, err
	}

	var currentTemplate Template
	var currentRowIndex int32
	var currentRow []string
	var templateRows int32
	templateRowCountMap := map[string]int32{}
	var rowCount int32

	currentTemplate = Template{}
	currentRow = make([]string, 0)
	currentRowIndex = 0
	templateRows = 1000
	rowCount = 0

	// Do one pass of the templates to understand the number of rows to use for each template
	currentTemplate.Identifier = ""
	var templateFirst bool
	templateFirst = true
	for _, row := range TemplateSheet[1:] {
		colCell := row[0]
		if colCell != "" {
			if !templateFirst {
				templateRowCountMap[currentTemplate.Identifier] = rowCount
			}
			templateFirst = false
			rowCount = 0
			currentTemplate.Identifier = colCell
		}
		rowCount += 1
	}
	templateRowCountMap[currentTemplate.Identifier] = rowCount
	rowCount = 0
	currentTemplate.Identifier = ""

	//rowParsing:
	for _, row := range TemplateSheet[1:] {
		if templateRows == currentRowIndex+1 {
			currentTemplate.Content[currentRowIndex] = currentRow
			a.SetTemplate(currentTemplate)
			currentRowIndex = 0
			templateRows = 0
			currentRow = make([]string, 0)
		} else {
			if currentTemplate.Identifier != "" {
				currentTemplate.Content[currentRowIndex] = currentRow
				currentRow = make([]string, 0)
				currentRowIndex += 1
			}
		}

		for columnIndex, colCell := range row {

			switch columnIndex {
			case 0:
				// Start of a new template
				if colCell != "" {
					templateRows = templateRowCountMap[colCell]
					currentTemplate = Template{}
					currentTemplate.Identifier = colCell
					currentRowIndex = 0
					currentTemplate.Content = make([][]string, templateRows)
				}
			default:
				// Expect a template to be defined
				if currentTemplate.Identifier == "" {
					return nil, errors.New("no template identified but parsing values")
				}
				currentRow = append(currentRow, colCell)
			}
		}
	}
	// Write the final template now where are out of the loop
	if !templateFirst {
		currentTemplate.Content[currentRowIndex] = currentRow
		a.SetTemplate(currentTemplate)
	}
	return a.Templates, nil
}
