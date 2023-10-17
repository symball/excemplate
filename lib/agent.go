package lib

import (
	"bufio"
	"errors"
	"github.com/xuri/excelize/v2"
	"os"
)

type Agent struct {
	Excel           *excelize.File
	Config          map[string]Sheet
	ConfigSheet     string
	TemplateSheet   string
	Templates       map[string]Template
	GeneratorsSheet string
	Generators      []GeneratorItem
}

type Config struct {
	File            string
	ConfigSheet     string
	TemplateSheet   string
	GeneratorsSheet string
}

func NewAgent(config Config) (*Agent, error) {
	agent := Agent{
		ConfigSheet:     config.ConfigSheet,
		TemplateSheet:   config.TemplateSheet,
		GeneratorsSheet: config.GeneratorsSheet,
		Config:          make(map[string]Sheet),
		Templates:       make(map[string]Template),
		Generators:      make([]GeneratorItem, 0),
	}

	err := agent.OpenExcel(config.File)
	if err != nil {
		return nil, err
	}
	return &agent, nil
}

func (a *Agent) CloseExcel() {
	_ = a.Excel.Close()
}
func (a *Agent) OpenExcel(filePath string) error {
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		return errors.New("could not find input file")
	}

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	fileHandle := bufio.NewReader(file)

	a.Excel, err = excelize.OpenReader(fileHandle)
	if err != nil {
		return err
	}
	return nil
}
