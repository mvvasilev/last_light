package model

import (
	"fmt"
	"mvvasilev/last_light/engine"
)

type ArrowSprite rune

//
// \  |  /
//
// ─  +  ─
//
// /  |  \

const (
	ProjectileSprite_NorthSouth         ArrowSprite = '|'
	ProjectileSprite_EastWest           ArrowSprite = '─'
	ProjectileSprite_NorthEastSouthWest ArrowSprite = '/'
	ProjectileSprite_NorthWestSouthEast ArrowSprite = '\\'
)

func ProjectileBehavior(eventLog *engine.GameEventLog, dungeon *Dungeon) func(npc Entity) (complete bool, requeue bool) {
	return func(npc Entity) (complete bool, requeue bool) {
		hasNext := ProjectileFollowPathNext(npc, eventLog, dungeon)

		return !hasNext, false
	}
}

func ProjectileFollowPathNext(npc Entity, eventLog *engine.GameEventLog, dungeon *Dungeon) (hasNext bool) {
	projectileData := npc.Projectile()
	positionData := npc.Positioned()

	if projectileData == nil || positionData == nil {
		return false
	}

	path := projectileData.Path
	next, hasNext := path.Next()

	nextTile := dungeon.CurrentLevel().TileAt(next.XY())

	nextTileEntityData := nextTile.Entity()

	dungeon.CurrentLevel().DropEntity(npc.UniqueId())

	positionData.Position = next

	// The next tile is impassable ( wall, void, etc. ) and contains no entity to damage
	// This is the end of the path
	if nextTileEntityData == nil && !nextTile.Passable() {
		return false
	}

	// Otherwise, if the tile is passible, but also the end of the path, stop here and despawn the projectile
	if nextTileEntityData == nil && next == projectileData.Path.To() {
		return false
	}

	// The next tile contains an entity, do damage to it if we have damage data
	if nextTileEntityData != nil {

		// The arrow strikes against its master, but to no avail, for I decree it to be illegal
		if nextTileEntityData.Entity == projectileData.Source {
			return
		}

		// Futher I decree, that should the arrow striketh at thyself, it shall be blocked from doing so
		if nextTileEntityData.Entity == npc {
			return
		}

		if projectileData.Source == nil {
			return false
		}

		ExecuteAttack(eventLog, projectileData.Source, nextTileEntityData.Entity)

		return false
	}

	dungeon.CurrentLevel().AddEntity(npc)

	return
}

func HostileNPCBehavior(eventLog *engine.GameEventLog, dungeon *Dungeon, player *Player) func(npc Entity) (complete bool, requeue bool) {
	return func(npc Entity) (complete bool, requeue bool) {
		CalcPathToPlayerAndMove(25, eventLog, dungeon, npc, player)

		return true, true
	}
}

func CalcPathToPlayerAndMove(simulationDistance int, eventLog *engine.GameEventLog, dungeon *Dungeon, npc Entity, player *Player) {
	if npc.Positioned().Position.DistanceSquared(player.Position()) > simulationDistance*simulationDistance {
		return
	}

	if npc.HealthData().IsDead {
		dungeon.CurrentLevel().DropEntity(npc.UniqueId())
		return
	}

	playerVisibleAndInRange := false
	hasLos, _ := HasLineOfSight(dungeon, npc.Positioned().Position, player.Position())

	if npc.Positioned().Position.DistanceSquared(player.Position()) < 144 && hasLos {
		playerVisibleAndInRange = true
	}

	if !playerVisibleAndInRange {
		randomMove := Direction(engine.RandInt(int(DirectionNone), int(East)))

		nextPos := npc.Positioned().Position

		switch randomMove {
		case North:
			nextPos = nextPos.WithOffset(0, -1)
		case South:
			nextPos = nextPos.WithOffset(0, +1)
		case West:
			nextPos = nextPos.WithOffset(-1, 0)
		case East:
			nextPos = nextPos.WithOffset(+1, 0)
		default:
			return
		}

		if dungeon.CurrentLevel().IsTilePassable(nextPos.XY()) {
			dungeon.CurrentLevel().MoveEntityTo(
				npc.UniqueId(),
				nextPos.X(),
				nextPos.Y(),
			)
		}

		return
	}

	if WithinHitRange(npc.Positioned().Position, player.Position()) {
		ExecuteAttack(eventLog, npc, player)
	}

	pathToPlayer := engine.FindPath(
		npc.Positioned().Position,
		player.Position(),
		12,
		func(x, y int) bool {
			if x == player.Position().X() && y == player.Position().Y() {
				return true
			}

			return dungeon.CurrentLevel().IsTilePassable(x, y)
		},
	)

	if pathToPlayer == nil {
		return
	}

	nextPos, hasNext := pathToPlayer.Next()

	if !hasNext {
		return
	}

	if nextPos.Equals(player.Position()) {
		return
	}

	dungeon.CurrentLevel().MoveEntityTo(npc.UniqueId(), nextPos.X(), nextPos.Y())
}

func HasLineOfSight(dungeon *Dungeon, start, end engine.Position) (hasLos bool, lastTile Tile) {
	positions := engine.CastRay(start, end)
	tile := dungeon.CurrentLevel().TileAt(positions[0].XY())

	for _, p := range positions {
		tile = dungeon.CurrentLevel().TileAt(p.XY())

		if tile.Opaque() {
			return false, tile
		}
	}

	return true, tile
}

func WithinHitRange(pos engine.Position, otherPos engine.Position) bool {
	return pos.WithOffset(-1, 0) == otherPos || pos.WithOffset(+1, 0) == otherPos || pos.WithOffset(0, -1) == otherPos || pos.WithOffset(0, +1) == otherPos
}

func ExecuteAttack(eventLog *engine.GameEventLog, attacker, victim Entity) {
	hit, precision, evasion, dmg, dmgType := CalculateAttack(attacker, victim)

	if attacker.Projectile() != nil {
		attacker = attacker.Projectile().Source
	}

	attackerName := "Unknown"

	if attacker.Named() != nil {
		attackerName = attacker.Named().Name
	}

	victimName := "Unknown"

	if victim.Named() != nil {
		victimName = victim.Named().Name
	}

	if !hit {
		eventLog.Log(fmt.Sprintf("%s attacked %s, but missed ( %v Evasion vs %v Precision)", attackerName, victimName, evasion, precision))
		return
	}

	if victim.HealthData() == nil {
		return
	}

	victim.HealthData().Health -= dmg

	if victim.HealthData().Health <= 0 {
		victim.HealthData().IsDead = true
		eventLog.Log(fmt.Sprintf("%s attacked %s, and was victorious ( %v Evasion vs %v Precision)", attackerName, victimName, evasion, precision))
		return
	}

	eventLog.Log(fmt.Sprintf("%s attacked %s, and hit for %v %v damage", attackerName, victimName, dmg, DamageTypeName(dmgType)))
}

func CalculateAttack(attacker, victim Entity) (hit bool, precisionRoll, evasionRoll int, damage int, damageType DamageType) {
	if attacker.Equipped() != nil && attacker.Equipped().Inventory.AtSlot(EquippedSlotDominantHand) != nil {
		weapon := attacker.Equipped().Inventory.AtSlot(EquippedSlotDominantHand)

		return PhysicalWeaponAttack(attacker, weapon, victim)
	} else {
		return UnarmedAttack(attacker, victim)
	}
}
