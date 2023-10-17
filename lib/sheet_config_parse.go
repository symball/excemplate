package lib

import (
	"errors"
	"log"
	"slices"
	"strconv"
	"unicode"
)

type Sheet struct {
	Name     string
	WriteRow int
}

func (a *Agent) SheetConfigParse() (map[string]Sheet, error) {

	// Get all sheets and create a map of last row to use as a place to start writing from
	workSpaceSheets := a.Excel.GetSheetList()
	for _, sheetName := range workSpaceSheets {
		rows, err := a.Excel.GetRows(sheetName)
		if err != nil {
			log.Printf("Could not get sheet %s \n", sheetName)
		}
		a.Config[sheetName] = Sheet{
			Name:     sheetName,
			WriteRow: len(rows) + 1,
		}
	}

	// If we have a sheet override
	if slices.Contains(workSpaceSheets, a.ConfigSheet) {
		sheetConfigSheet, err := a.Excel.GetRows(a.ConfigSheet)
		if err != nil {
			return nil, errors.New("could not get config sheet for reading")
		}
		for _, row := range sheetConfigSheet[1:] {
			sheet := Sheet{}
			// Expect index zero to contain Template
		columnParsing:
			for columnIndex, colCell := range row {
				switch columnIndex {
				case 0:
					sheet.Name = colCell
				case 1:
					// Ensure the entry is valid number
					for _, r := range colCell {
						if !unicode.IsDigit(r) {
							return nil, errors.New("unable to convert starting row. given: " + colCell)
						}
					}
					sheet.WriteRow, _ = strconv.Atoi(colCell)
				default:
					break columnParsing
				}
			}
			a.Config[sheet.Name] = sheet
		}
	}
	return a.Config, nil
}
