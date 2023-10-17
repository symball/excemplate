package lib

import "errors"

type GeneratorItem struct {
	Template    string
	OutputSheet string
	Variables   map[string]string
}

func (a *Agent) GeneratorParse() ([]GeneratorItem, error) {

	if a.GeneratorsSheet == "" {
		return nil, errors.New("generator sheet not defined")
	}

	generatorSheet, err := a.Excel.GetRows(a.GeneratorsSheet)
	if err != nil {
		return nil, err
	}

	// Take a single pass through the first row to decide on variables to set
	variableMap := []string{}
	for _, heading := range generatorSheet[0] {
		variableMap = append(variableMap, heading)
	}

	generatorItemCount := len(generatorSheet) - 1
	// Set the list size manually to save on dynamic list modification
	a.Generators = make([]GeneratorItem, generatorItemCount)

	var variables map[string]string
	var generatorItem GeneratorItem

	for idx, row := range generatorSheet[1:] {

		generatorItem = GeneratorItem{}
		variables = make(map[string]string)

		// Expect index zero to contain Template
		for columnIndex, colCell := range row {
			switch columnIndex {
			case 0:
				if colCell == "" {
					return nil, errors.New("template name required for generator")
				}
				if _, templateError := a.GetTemplate(colCell); templateError != nil {
					return nil, templateError
				}
				generatorItem.Template = colCell
			case 1:
				if colCell == "" {
					return nil, errors.New("output sheet required for generator")
				}
				generatorItem.OutputSheet = colCell
			default:
				variables[variableMap[columnIndex]] = colCell
			}
		}

		generatorItem.Variables = variables
		a.Generators[idx] = generatorItem
	}
	return a.Generators, nil
}
