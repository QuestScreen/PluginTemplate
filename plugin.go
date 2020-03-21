package main

import (
	"github.com/QuestScreen/PluginTemplate/moduletemplate"
	"github.com/QuestScreen/QuestScreen/generated"
	"github.com/QuestScreen/api"
)

// QSPlugin is the plugin's descriptor.
var QSPlugin = api.Plugin{
	Name: "<TODO: enter name>",
	Modules: []*api.Module{
		&moduletemplate.Descriptor,
		// TODO: add modules here
	},
	AdditionalJS:    generated.MustAsset("web/js/base.js"),
	AdditionalHTML:  generated.MustAsset("web/html/base.html"),
	AdditionalCSS:   nil,
	SystemTemplates: nil,
	GroupTemplates: []api.GroupTemplate{
		{
			Name: "Minimal", Description: "Contains a „Main“ scene with no modules.",
			Config: []byte("name: Minimal"),
			Scenes: []api.SceneTmplRef{
				{Name: "Main", TmplIndex: 0},
			},
		}, {
			Name:        "Base",
			Description: "Contains a „Main“ scene with base modules.",
			Config:      []byte("name: Base"),
			Scenes: []api.SceneTmplRef{
				{Name: "Main", TmplIndex: 1},
			},
		},
	},
	SceneTemplates: []api.SceneTemplate{
		{
			Name: "Empty", Description: "A scene with no modules.",
			Config: []byte("name: Empty"),
		}, {
			Name:        "BaseMain",
			Description: "A scene with background, title, herolist and overlay enabled.",
			Config: []byte(`name: BaseMain
modules:
  background:
    enabled: true
  herolist:
    enabled: true
  overlays:
    enabled: true
  title:
    enabled: true`),
		},
	},
}

// required to compile; although never called
func main() {}
