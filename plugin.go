package main

import (
	"github.com/QuestScreen/api"
	"github.com/QuestScreen/plugin-template/generated"
	"github.com/QuestScreen/plugin-template/moduletemplate"
)

// QSPlugin is the plugin's descriptor.
var QSPlugin = api.Plugin{
	Name: "My Plugin", // TODO: change
	Modules: []*api.Module{
		&moduletemplate.Descriptor,
		// TODO: add modules here
	},
	AdditionalJS:   generated.MustAsset("web/js/controllers.js"),
	AdditionalHTML: generated.MustAsset("web/html/templates.html"),
	AdditionalCSS:  nil,
	SystemTemplates: []api.SystemTemplate{
		// TODO: add proper system templates here or delete this.
		{
			Name: "MySystem", ID: "myplugin-mysystem",
			Config: []byte("name: MySystem"),
		},
	},
	GroupTemplates: []api.GroupTemplate{
		// TODO: add proper group templates here or delete this.
		{
			Name: "MyTemplate", Description: "Contains a „Main“ scene with no modules.",
			Config: []byte("name: MyTemplate"),
			Scenes: []api.SceneTmplRef{
				{Name: "Main", TmplIndex: 0},
			},
		},
	},
	SceneTemplates: []api.SceneTemplate{
		// TODO: add proper scene templates here or delete this.
		{
			Name: "Empty", Description: "A scene with no modules.",
			Config: []byte("name: Empty"),
		},
	},
}

// required to compile; although never called
func main() {}
