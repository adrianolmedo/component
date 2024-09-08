package components

import (
	"encoding/json"

	"github.com/adrianolmedo/components/alpine"
	"github.com/adrianolmedo/components/staticdata"

	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

type ProvinceAndCityPickerProps struct {
	DefaultProvince string
	DefaultCity     string
	ContainerClass  string
}

func ProvinceAndCityPicker(props ProvinceAndCityPickerProps) gomponents.Node {
	provincesAndCities := map[string][]string{}
	for _, pc := range staticdata.ProvincesWithCities() {
		provincesAndCities[pc.ProvinceName] = pc.Cities
	}

	jsonByte, err := json.Marshal(provincesAndCities)
	if err != nil {
		return html.Div(
			gomponents.Text("[1704216567] Error al procesar las provincias y ciudades, contacte con el administrador"),
		)
	}
	jsonStr := string(jsonByte)

	provincesOptions := make([]gomponents.Node, len(provincesAndCities))
	for provinceName := range provincesAndCities {
		provincesOptions = append(provincesOptions,
			html.Option(
				html.Value(provinceName),
				gomponents.Text(provinceName),
			),
		)
	}

	return html.Div(
		alpine.XData(`{
			"pickedProvince": "`+props.DefaultProvince+`",
			"pickedCity": "`+props.DefaultCity+`",
			"provinces": `+jsonStr+`
		}`),

		components.Classes{
			props.ContainerClass: props.ContainerClass != "",
		},

		SelectControl(SelectControlProps{
			Name:        "province",
			Title:       "Provincia",
			Placeholder: "Selecciona una provincia",
			Required:    true,
			SelectChildren: []gomponents.Node{
				alpine.XModel("pickedProvince"),
				gomponents.Group(provincesOptions),
			},
		}),

		SelectControl(SelectControlProps{
			Name:        "city",
			Title:       "Ciudad",
			Placeholder: "Selecciona una ciudad",
			Required:    true,
			SelectChildren: []gomponents.Node{
				gomponents.El(
					"template",
					gomponents.Attr("x-for", "city in provinces[pickedProvince]"),
					html.Option(
						alpine.XBind("selected", "city === pickedCity"),
						gomponents.Attr("x-text", "city"),
						gomponents.Attr("x-bind:value", "city"),
					),
				),
			},
		}),
	)
}
