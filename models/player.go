package models

import (
	"can-u-scape/constants"
	"log"

	"github.com/oakmound/oak/v4"
	"github.com/oakmound/oak/v4/alg/floatgeom"
	"github.com/oakmound/oak/v4/alg/intgeom"
	"github.com/oakmound/oak/v4/entities"
	"github.com/oakmound/oak/v4/event"
	"github.com/oakmound/oak/v4/key"
	"github.com/oakmound/oak/v4/render"
	"github.com/oakmound/oak/v4/scene"
)

type Player struct {
	Score int
	*entities.Entity
}

func NewPlayer(ctx *scene.Context) *Player {
	playerRender, err := render.LoadSprite("assets/images/Terran.png")
	if err != nil {
		log.Fatal(err)
	}

	entity := entities.New(ctx,
		entities.WithRect(floatgeom.NewRect2WH(100, 100, 32, 32)),
		entities.WithRenderable(playerRender),
		entities.WithSpeed(floatgeom.Point2{3, 3}),
		entities.WithDrawLayers([]int{1, 2}),
	)

	return &Player{Entity: entity}
}

func (p *Player) Run(ctx *scene.Context, player *entities.Entity) {
	screenCenter := ctx.Window.Bounds().DivConst(2)
	event.Bind(ctx, event.Enter, player, func(char *entities.Entity, ev event.EnterPayload) event.Response {
		if oak.IsDown(key.W) {
			char.Delta[1] += (-char.Speed.Y() * ev.TickPercent)
		}
		if oak.IsDown(key.A) {
			char.Delta[0] += (-char.Speed.X() * ev.TickPercent)
		}
		if oak.IsDown(key.S) {
			char.Delta[1] += (char.Speed.Y() * ev.TickPercent)
		}
		if oak.IsDown(key.D) {
			char.Delta[0] += (char.Speed.X() * ev.TickPercent)
		}
		ctx.Window.(*oak.Window).DoBetweenDraws(func() {
			char.ShiftDelta()
			oak.SetViewport(
				intgeom.Point2{int(char.X()), int(char.Y())}.Sub(screenCenter),
			)
			char.Delta = floatgeom.Point2{}
		})
		if char.X() < 0 {
			char.SetX(0)
		} else if char.X() > constants.FieldWidth-char.W() {
			char.SetX(constants.FieldWidth - char.W())
		}
		if char.Y() < 0 {
			char.SetY(0)
		} else if char.Y() > constants.FieldHeight-char.H() {
			char.SetY(constants.FieldHeight - char.H())
		}

		hitEnemy := char.HitLabel(constants.EnemyLabel)
		if hitEnemy != nil {
			p.SetScore(0)
			ctx.Window.NextScene()
		}

		hitBoost := char.HitLabel(constants.BoostLabel)
		if hitBoost != nil {
			p.AddBoost()
			event.TriggerForCallerOn(ctx, hitBoost.CID, Destroy, struct{}{})
		}

		return 0
	})
}

func (p *Player) GetScore() int {
	return p.Score
}

func (p *Player) SetScore(score int) {
	p.Score = score
}

func (p *Player) AddBoost() {
	p.Score += 10
}
