package main

import (
	"gopkg.in/AlecAivazis/survey.v1"
	"gopkg.in/AlecAivazis/survey.v1/tests/util"
)

var answer = ""

var goodTable = []TestUtil.TestTableEntry{
	{
		"standard", &survey.Select{
			Message: "Choose a color:",
			Options: []survey.SelectOption{survey.StringOption{"red"}, survey.StringOption{"blue"}, survey.StringOption{"green"}},
		}, &answer,
	},
	{
		"short", &survey.Select{
			Message: "Choose a color:",
			Options: []survey.SelectOption{survey.StringOption{"red"}, survey.StringOption{"blue"}},
		}, &answer,
	},
	{
		"default", &survey.Select{
			Message: "Choose a color (should default blue):",
			Options: []survey.SelectOption{survey.StringOption{"red"}, survey.StringOption{"blue"}, survey.StringOption{"green"}},
			Default: survey.StringOption{"blue"},
		}, &answer,
	},
	{
		"one", &survey.Select{
			Message: "Choose one:",
			Options: []survey.SelectOption{survey.StringOption{"hello"}},
		}, &answer,
	},
	{
		"no help, type ?", &survey.Select{
			Message: "Choose a color:",
			Options: []survey.SelectOption{survey.StringOption{"red"}, survey.StringOption{"blue"}},
		}, &answer,
	},
	{
		"passes through bottom", &survey.Select{
			Message: "Choose one:",
			Options: []survey.SelectOption{survey.StringOption{"red"}, survey.StringOption{"blue"}},
		}, &answer,
	},
	{
		"passes through top", &survey.Select{
			Message: "Choose one:",
			Options: []survey.SelectOption{survey.StringOption{"red"}, survey.StringOption{"blue"}},
		}, &answer,
	},
	{
		"can navigate with j/k", &survey.Select{
			Message: "Choose one:",
			Options: []survey.SelectOption{survey.StringOption{"red"}, survey.StringOption{"blue"}, survey.StringOption{"green"}},
		}, &answer,
	},
}

var badTable = []TestUtil.TestTableEntry{
	{
		"no options", &survey.Select{
			Message: "Choose one:",
		}, &answer,
	},
}

func main() {
	TestUtil.RunTable(goodTable)
	TestUtil.RunErrorTable(badTable)
}
