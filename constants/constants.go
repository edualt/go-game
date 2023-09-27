package constants

import "github.com/oakmound/oak/v4/collision"

const (
	FieldWidth  = 640
	FieldHeight = 500

	EnemyLabel   collision.Label = 1
	EnemyRefresh                 = 50
	EnemySpeed                   = 2

	BoostLabel   collision.Label = 2
	BoostRefresh                 = 120
	BoostSpeed                   = 1
)
