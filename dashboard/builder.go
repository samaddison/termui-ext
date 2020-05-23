package dashboard

import (
	"encoding/json"
	"io/ioutil"
	termui_ext "termui-ext"
)

func BuildGridFromDefinition(definition *Definition) (*termui_ext.Grid, *[]error) {
	widgets, errors := buildWidgets(definition.Widgets)
	for _, v := range *errors {
		if v != nil {
			return nil, errors
		}
	}
	grid, err := buildGrid(definition.Grid, widgets)
	if err != nil {
		result := []error{err}
		return nil, &result
	} else {
		return grid, nil
	}
}

func BuildDashboardDefinitionFromJSON(jsonString string) (*Definition, error) {
	d := Definition{}
	err := json.Unmarshal([]byte(jsonString), &d)
	if err != nil {
		return nil, err
	} else {
		return &d, nil
	}
}

func BuildGrid(jsonFilePath string) (*termui_ext.Grid, *[]error) {
	jsonString, err := ioutil.ReadFile(jsonFilePath)
	if err != nil {
		result := []error{err}
		return nil, &result
	} else {
		dashboardFromJSON, err := BuildDashboardDefinitionFromJSON(string(jsonString))
		if err != nil {
			result := []error{err}
			return nil, &result
		} else {
			return BuildGridFromDefinition(dashboardFromJSON)
		}
	}
}
