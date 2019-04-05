package dev

import (
	cliexchange "github.com/open-horizon/anax/cli/exchange"
)

const PATTERN_DEFINITION_FILE = "pattern.definition.json"
const PATTERN_DEFINITION_ALL_ARCHES_FILE = "pattern.definition-all-arches.json"

// Sort of like a constructor, it creates an in memory object except that it is created from the patern definition config
// file in the current project. This function assumes the caller has determined the exact location of the file.
func GetPatternDefinition(directory string, name string) (*cliexchange.PatternFile, error) {

	res := new(cliexchange.PatternFile)

	// GetFile will write to the res object, demarshalling the bytes into a json object that can be returned.
	if err := GetFile(directory, name, res); err != nil {
		return nil, err
	}
	return res, nil

}

// It creates a pattern definition config object and writes it to the project
// in the file system.
func CreatePatternDefinition(directory string) error {

	// Create a pattern definition config object with fillins/place-holders for configuration.
	res := new(cliexchange.PatternFile)

	sv := new(cliexchange.ServiceChoiceFile)
	sv.Version = "$SERVICE_VERSION"

	sref := new(cliexchange.ServiceReferenceFile)
	sref.ServiceOrg = "$HZN_ORG_ID"
	sref.ServiceURL = "$SERVICE_NAME"
	sref.ServiceArch = "$ARCH"
	sref.ServiceVersions = []cliexchange.ServiceChoiceFile{*sv}

	res.Label = "Edge $SERVICE_NAME Service Pattern for $ARCH"
	res.Description = "Pattern for $SERVICE_NAME"
	res.Public = true
	res.Services = []cliexchange.ServiceReferenceFile{*sref}

	// Convert the object to JSON and write it into the project.
	return CreateFile(directory, PATTERN_DEFINITION_FILE, res)
}

// It creates a pattern definition config object for all architectures and writes it to the project
// in the file system.
func CreatePatternDefinitionAllArches(directory string) error {

	// Create a pattern definition config object with fillins/place-holders for configuration.
	res := new(cliexchange.PatternFile)

	sv := new(cliexchange.ServiceChoiceFile)
	sv.Version = "$SERVICE_VERSION"

	res.Label = "Edge $SERVICE_NAME Service Pattern for all architectures"
	res.Description = "Pattern for $SERVICE_NAME"
	res.Public = true
	res.Services = []cliexchange.ServiceReferenceFile{}

	sref := new(cliexchange.ServiceReferenceFile)
	sref.ServiceOrg = "$HZN_ORG_ID"
	sref.ServiceURL = "$SERVICE_NAME"
	sref.ServiceVersions = []cliexchange.ServiceChoiceFile{*sv}

	arches := []string{"amd64", "arm", "arm64"}
	for _, arch := range arches {
		sref.ServiceArch = arch
		res.Services = append(res.Services, *sref)
	}

	// Convert the object to JSON and write it into the project.
	return CreateFile(directory, PATTERN_DEFINITION_ALL_ARCHES_FILE, res)
}