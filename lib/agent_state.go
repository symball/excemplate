package lib

import "errors"

func (a *Agent) GetTemplate(templateIdentifier string) (*Template, error) {
	template, found := a.Templates[templateIdentifier]
	if !found {
		return nil, errors.New("template not found: " + templateIdentifier)
	}
	return &template, nil
}
func (a *Agent) SetTemplate(template Template) {
	a.Templates[template.Identifier] = template
}

func (a *Agent) AddItem(item GeneratorItem) {
	a.Generators = append(a.Generators, item)
}
