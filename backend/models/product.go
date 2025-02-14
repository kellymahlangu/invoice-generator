package models

import (
	"strconv"

	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"

	"github.com/johnfercher/maroto/v2/pkg/components/text"
)

type Product struct {
	Item        string
	Description string
	Quantity    int
	UnitPrice   float64
	Subtotal    float64
}

var background = &props.Color{
	Red:   200,
	Green: 200,
	Blue:  200,
}

func (p Product) GetHeader() core.Row {
	return row.New(10).Add(
		text.NewCol(4, "Item", props.Text{Style: fontstyle.Bold}),
		text.NewCol(8, "Description", props.Text{Style: fontstyle.Bold}),
		text.NewCol(8, "Quantity", props.Text{Style: fontstyle.Bold}),
		text.NewCol(8, "UnitPrice", props.Text{Style: fontstyle.Bold}),
		text.NewCol(8, "Subtotal", props.Text{Style: fontstyle.Bold}),
	)
}

func (p Product) GetContent(i int) core.Row {
	r := row.New(8).Add(
		text.NewCol(2, p.Item),
		text.NewCol(4, p.Description),
		text.NewCol(2, strconv.Itoa(p.Quantity)),
		text.NewCol(2, strconv.FormatFloat(p.UnitPrice, 'g', 5, 64)),
		text.NewCol(2, strconv.FormatFloat(p.Subtotal, 'g', 5, 64)),
	)

	if i%2 == 0 {
		r.WithStyle(&props.Cell{
			BackgroundColor: background,
		})
	}

	return r
}
