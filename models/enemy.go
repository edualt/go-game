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

type Enemy struct {
	*entities.Entity
}

var (
	Destroy = event.RegisterEvent[struct{}]()
)

func NewEnemy(ctx *scene.Context) *Enemy {
	x, y := getEnemyPos()
	enemyRender, err := render.LoadSprite("assets/images/Black_hole.png") // Reemplaza "enemy.png" con el nombre de tu imagen
	if err != nil {
		log.Fatal(err)
	}

	entity := entities.New(ctx,
		entities.WithRect(floatgeom.NewRect2WH(x, y, 40, 40)),
		entities.WithRenderable(enemyRender),
		entities.WithDrawLayers([]int{1, 2}),
		entities.WithLabel(constants.EnemyLabel),
		entities.WithSpeed(floatgeom.Point2{constants.EnemySpeed, 0}),
	)

	return &Enemy{Entity: entity}
}

func (e *Enemy) Run(ctx *scene.Context, entity *entities.Entity) {
	event.Bind(ctx, event.Enter, entity, func(e *entities.Entity, ev event.EnterPayload) event.Response {
		e.Shift(e.Speed)
		if e.X() <= 0 || e.X()+e.W() >= constants.FieldWidth {
			e.Destroy()
		}
		return 0
	})

	event.Bind(ctx, Destroy, entity, func(e *entities.Entity, nothing struct{}) event.Response {
		e.Destroy()
		return 0
	})
}

func getEnemyPos() (float64, float64) {
	perimeter := constants.FieldWidth*2 + constants.FieldHeight*2
	pos := int(rand.Float64() * float64(perimeter))
	if pos < constants.FieldWidth {
		return float64(pos), 0
	}
	pos -= constants.FieldWidth
	if pos < constants.FieldHeight {
		return constants.FieldWidth, float64(pos)
	}
	pos -= constants.FieldHeight
	if pos < constants.FieldWidth {
		return float64(constants.FieldWidth - pos), constants.FieldHeight
	}
	return 0, float64(constants.FieldHeight - pos + constants.FieldWidth)
}
