// SPDX-FileCopyrightText: 2022 Free Mobile
// SPDX-License-Identifier: AGPL-3.0-only

package console

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Configuration describes the configuration for the console component.
type Configuration struct {
	// ServeLiveFS serve files from the filesystem instead of the embedded versions.
	ServeLiveFS bool `yaml:"-"`
	// Version is the version to display to the user.
	Version string `yaml:"-"`
	// DefaultVisualizeOptions define some defaults for the "visualize" tab.
	DefaultVisualizeOptions VisualizeOptionsConfiguration
	// HomepageTopWidgets defines the list of widgets to display on the home page.
	HomepageTopWidgets []string `validate:"dive,oneof=src-as dst-as src-country dst-country exporter protocol etype src-port dst-port"`
	// DimensionsLimit put an upper limit to the number of dimensions to return.
	DimensionsLimit int `validate:"min=10"`
}

// VisualizeOptionsConfiguration defines options for the "visualize" tab.
type VisualizeOptionsConfiguration struct {
	// Start is the start time (as a string)
	Start string `json:"start"`
	// End is the end time (as string)
	End string `json:"end"`
	// Filter  is the the filter string
	Filter string `json:"filter"`
	// Dimensions is the array of dimensions to use
	Dimensions []queryColumn `json:"dimensions"`
}

// DefaultConfiguration represents the default configuration for the console component.
func DefaultConfiguration() Configuration {
	return Configuration{
		DefaultVisualizeOptions: VisualizeOptionsConfiguration{
			Start:      "6 hours ago",
			End:        "now",
			Filter:     "InIfBoundary = external",
			Dimensions: []queryColumn{queryColumnSrcAS},
		},
		HomepageTopWidgets: []string{"src-as", "src-port", "protocol", "src-country", "etype"},
		DimensionsLimit:    50,
	}
}

func (c *Component) configHandlerFunc(gc *gin.Context) {
	gc.JSON(http.StatusOK, gin.H{
		"version":                 c.config.Version,
		"defaultVisualizeOptions": c.config.DefaultVisualizeOptions,
		"topWidgets":              c.config.HomepageTopWidgets,
	})
}
