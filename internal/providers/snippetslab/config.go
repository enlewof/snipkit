package snippetslab

import (
	"github.com/lemoony/snippet-kit/internal/utils"
)

type Config struct {
	Enabled     bool     `yaml:"enabled" head_comment:"Set to true if you want to use SnippetsLab."`
	LibraryPath string   `yaml:"libraryPath" head_comment:"Path to your *.snippetslablibrary file.\nSnipKit will try to detect this file automatically when generating the config."`
	IncludeTags []string `yaml:"includeTags" head_comment:"If this list is not empty, only those snippets that match the listed tags will be provided to you."`
	ExcludeTags []string `yaml:"excludeTags" head_comment:"If this list is not empty, snippets that have one of the listed tags will not be provided to you."`
}

func AutoDiscoveryConfig(system *utils.System) Config {
	result := Config{
		Enabled:     false,
		LibraryPath: "/path/to/main.snippetslablibrary",
	}

	libraryURL, err := getLibraryURL(system)
	if err != nil {
		return result
	}

	if ok, err := libraryURL.validate(); err != nil || !ok {
		return result
	} else if basePath, err := libraryURL.basePath(); err == nil {
		result.Enabled = true
		result.LibraryPath = basePath
	}

	return result
}
