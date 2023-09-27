package models

import (
	"can-u-scape/constants"
	"log"
	"math/rand"

	"github.com/oakmound/oak/v4/alg/floatgeom"
	"github.com/oakmound/oak/v4/entities"
	"github.com/oakmound/oak/v4/event"
	"github.com/oakmound/oak/v4/render"
	"github.com/oakmound/oak/v4/scene"
)

type Booster struct {
	*entities.Entity
}

func NewBooster(ctx *scene.Context) *Booster {
	x, y := getBoosterPos()
	boosterRender, err := render.LoadSprite("assets/images/Baren.png")
	if err != nil {
		log.Fatal(err)
	}

	entity := entities.New(ctx,
		entities.WithRect(floatgeom.NewRect2WH(x, y, 32, 32)),
		entities.WithRenderable(boosterRender),
		entities.WithSpeed(floatgeom.Point2{3, 3}),
		entities.WithDrawLayers([]int{1, 2}),
		entities.WithLabel(constants.BoostLabel),
	)

	return &Booster{Entity: entity}
}

func (b *Booster) Run(ctx *scene.Context, entity *entities.Entity) {
	event.Bind(ctx, event.Enter, entity, func(e *entities.Entity, ev event.EnterPayload) event.Response {
		x, y := entity.X(), entity.Y()
		pt := floatgeom.Point2{x, y}
		pt2 := floatgeom.Point2{rand.Float64() * constants.FieldWidth, rand.Float64() * constants.FieldHeight}
		delta := pt2.Sub(pt).Normalize().MulConst(constants.BoostSpeed * ev.TickPercent)
		entity.Shift(delta)
		return 0
	})

	event.Bind(ctx, Destroy, entity, func(e *entities.Entity, nothing struct{}) event.Response {
		e.Destroy()
		return 0
	})
}

func getBoosterPos() (float64, float64) {
	perimeter := constants.FieldWidth*2 + constants.FieldHeight*2
	pos := int(rand.Float64() * float64(perimeter))
	if pos < constants.FieldWidth {
		return float64(pos), 0
	}
	pos -= constants.FieldWidth
	if pos < constants.FieldHeight {
		return float64(constants.FieldWidth), float64(pos)
	}
	pos -= constants.FieldHeight
	if pos < constants.FieldWidth {
		return float64(pos), float64(constants.FieldHeight)
	}
	pos -= constants.FieldWidth
	return 0, float64(pos)
}
