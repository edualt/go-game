package scenes

import (
	"can-u-scape/constants"
	"can-u-scape/models"
	"image/color"
	"time"

	"github.com/oakmound/oak/v4"
	"github.com/oakmound/oak/v4/alg/intgeom"
	"github.com/oakmound/oak/v4/event"
	"github.com/oakmound/oak/v4/render"
	"github.com/oakmound/oak/v4/scene"
)

var (
	playerX *float64
	playerY *float64
	destroy = event.RegisterEvent[struct{}]()
)

func countScore(player *models.Player) {
	ticker := time.NewTicker(time.Second)

	for {
		select {
		case <-ticker.C:
			player.SetScore(player.GetScore() + 1)
		}
	}
}

func MainScene(ctx *scene.Context) {
	render.Draw(render.NewDrawFPS(0, nil, 10, 10), 2, 0)
	render.Draw(render.NewLogicFPS(0, nil, 10, 20), 2, 0)
	oak.SetViewportBounds(intgeom.NewRect2(0, 0, 0, 0))
	oak.SetTitle("Can U Scape?")

	player := models.NewPlayer(ctx)
	player.Run(ctx, player.Entity)

	playerX = &player.Rect.Min[0]
	playerY = &player.Rect.Min[1]

	t := render.DefaultFont().NewIntText(&player.Score, 200, 30)
	render.Draw(t, 2, 0)
	go countScore(player)

	event.GlobalBind(ctx, event.Enter, func(enterPayload event.EnterPayload) event.Response {
		if enterPayload.FramesElapsed%constants.EnemyRefresh == 0 {
			enemy := models.NewEnemy(ctx)
			go enemy.Run(ctx, enemy.Entity)
		}
		return 0
	})

	event.GlobalBind(ctx, event.Enter, func(enterPayload event.EnterPayload) event.Response {
		if enterPayload.FramesElapsed%constants.BoostRefresh == 0 {
			booster := models.NewBooster(ctx)
			go booster.Run(ctx, booster.Entity)
		}
		return 0
	})

	bgColor := color.RGBA{0, 0, 0, 255}
	bg := render.NewColorBox(constants.FieldWidth, constants.FieldHeight, bgColor)
	render.Draw(bg, 0, 0)
}
