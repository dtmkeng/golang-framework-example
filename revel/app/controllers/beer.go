package controllers

import (
	"github.com/revel/revel"
	"github.com/dtmkeng/framework-exmaple/revel/app/models"
)

type Beers struct {
	*revel.Controller
}
var beers = []models.Beer{
	models.Beer{1, "La Trappe Quadrupel Oak Aged", "Ale", "Bierbrouwerij De Koningshoeven"},
	models.Beer{2, "Cocoa Psycho", "Porter", "BrewDog"},
	models.Beer{3, "American Dream", "Lager", "Mikkeller"},
}

func (c Beers) List() revel.Result {
	return c.RenderJSON(beers)
}

func (c Beers) Show(beerID int) revel.Result {
	var res models.Beer

	for _, beer := range beers {
		if beer.ID == beerID {
			res = beer
		}
	}

	if res.ID == 0 {
		return c.NotFound("Could not find beer")
	}

	return c.RenderJSON(res)
}