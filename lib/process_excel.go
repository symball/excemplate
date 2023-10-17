package lib

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"text/template"
)

func (a *Agent) ProcessExcelFile(outputFile string) error {

	// Start cycling through questions
	log.Println("Starting output")
	for _, generatorItem := range a.Generators {
		templateObject, notFoundError := a.GetTemplate(generatorItem.Template)
		if notFoundError != nil {
			return errors.New(fmt.Sprintf("Could not find template object: %s \n", generatorItem.Template))
		}

		// Ensure sheet available and get sheet
		_, found := a.Config[generatorItem.OutputSheet]
		if !found {
			_, err := a.Excel.NewSheet(generatorItem.OutputSheet)
			if err != nil {
				return errors.New("could not create sheet: " + generatorItem.OutputSheet)
			}
			newSheet := Sheet{
				Name:     generatorItem.OutputSheet,
				WriteRow: 1,
			}
			a.Config[generatorItem.OutputSheet] = newSheet
		}

		for _, templateRow := range templateObject.Content {

			for columnIndex, columnText := range templateRow {
				coordinates, err := excelize.CoordinatesToCellName(columnIndex+1, a.Config[generatorItem.OutputSheet].WriteRow)
				if err != nil {
					return errors.New("unable to get coordinates")
				}

				tmpl, err := template.New("question").Option("missingkey=zero").Parse(columnText)
				if err != nil {
					panic(err)
				}
				var renderBuffer bytes.Buffer
				err = tmpl.Execute(&renderBuffer, generatorItem.Variables)
				if err != nil {
					panic(err)
				}

				err = a.Excel.SetCellValue(generatorItem.OutputSheet, coordinates, renderBuffer.String())
				if err != nil {
					panic(err)
				}
			}
			// Update the sheet
			currentSheet := a.Config[generatorItem.OutputSheet]
			currentSheet.WriteRow += 1
			a.Config[generatorItem.OutputSheet] = currentSheet
		}
	}

	// Cleanup the generator sheets
	_, configSheetExists := a.Config[a.ConfigSheet]
	if configSheetExists {
		_ = a.Excel.DeleteSheet(a.ConfigSheet)
	}
	_, templateSheetExists := a.Config[a.TemplateSheet]
	if templateSheetExists {
		_ = a.Excel.DeleteSheet(a.TemplateSheet)
	}
	_, itemsSheetExists := a.Config[a.GeneratorsSheet]
	if itemsSheetExists {
		_ = a.Excel.DeleteSheet(a.GeneratorsSheet)
	}
	if err := a.Excel.SaveAs(outputFile); err != nil {
		return err
	}
	return nil
}
