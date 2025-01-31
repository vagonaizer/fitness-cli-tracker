package plot

import (
	"fitness-cli-tracker/internal/models"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func PlotWeight(records []models.Record) error {
	p := plot.New()
	p.Title.Text = "График веса"
	p.X.Label.Text = "Дата"
	p.Y.Label.Text = "Вес"

	points := make(plotter.XYs, len(records))
	for i, r := range records {
		points[i].X = float64(r.Date.Unix())
		points[i].Y = r.Weight
	}

	line, err := plotter.NewLine(points)
	if err != nil {
		return err
	}
	p.Add(line)

	if err := p.Save(10*vg.Inch, 5*vg.Inch, "weight.png"); err != nil {
		return err
	}
	return nil
}
