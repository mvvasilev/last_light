package model

import (
	"fmt"
	"mvvasilev/last_light/engine"
	"mvvasilev/last_light/game/systems"

	"github.com/gdamore/tcell/v2"
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

		// TODO: calculate additional projectile damage
		ExecuteAttack(eventLog, projectileData.Source, nextTileEntityData.Entity, true)

		return false
	}

	dungeon.CurrentLevel().AddEntity(npc)

	return
}

func HostileMeleeNPCBehavior(eventLog *engine.GameEventLog, dungeon *Dungeon, player *Player) func(npc Entity) (complete bool, requeue bool) {
	return func(npc Entity) (complete bool, requeue bool) {
		CalcPathToPlayerAndMove(25, eventLog, dungeon, npc, player)

		return true, true
	}
}

func HostileRangedNPCBehavior(eventLog *engine.GameEventLog, dungeon *Dungeon, player *Player) func(npc Entity) (complete bool, requeue bool) {
	return func(npc Entity) (complete bool, requeue bool) {
		return true, true
	}
}

func CalcPathToPlayerAndKeepDistance(simulationDistance int, eventLog *engine.GameEventLog, dungeon *Dungeon, npc Entity, player *Player) {

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
		ExecuteAttack(eventLog, npc, player, false)
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

func ExecuteAttack(eventLog *engine.GameEventLog, attacker, victim Entity, isRanged bool) {
	hit, precision, evasion, dmg, dmgType := CalculateAttack(attacker, victim, isRanged)

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

func CalculateAttack(attacker, victim Entity, isRanged bool) (hit bool, precisionRoll, evasionRoll int, damage int, damageType DamageType) {
	if attacker.Equipped() != nil && attacker.Equipped().Inventory.AtSlot(EquippedSlotDominantHand) != nil {
		weapon := attacker.Equipped().Inventory.AtSlot(EquippedSlotDominantHand)

		if weapon.Damaging() != nil {
			// If the weapon is ranged, but the combat isn't, do unarmed
			if isRanged && !weapon.Damaging().IsRanged {
				return UnarmedAttack(attacker, victim)
			}

			// Doing melee damage from ranged? Don't think so.
			if !isRanged && weapon.Damaging().IsRanged {
				return false, 0, 0, 0, DamageType_Physical_Unarmed
			}
		}

		return PhysicalWeaponAttack(attacker, weapon, victim)
	} else {
		return UnarmedAttack(attacker, victim)
	}
}

func ShootProjectile(shooter Entity, target engine.Position, eventLog *engine.GameEventLog, dungeon *Dungeon, turnSystem *systems.TurnSystem) (success bool) {
	success = false

	logMessage := func(msg string) {
		if eventLog != nil {
			eventLog.Log(msg)
		}
	}

	if shooter.Equipped() == nil || shooter.Positioned() == nil {
		return
	}

	shooterName := "Unknown"

	if shooter.Named() != nil {
		shooterName = shooter.Named().Name
	}

	weapon := shooter.Equipped().Inventory.AtSlot(EquippedSlotDominantHand)

	if weapon == nil {
		logMessage(fmt.Sprintf("%s wants to shoot, but doesn't have anything equipped!", shooterName))

		return
	}

	if weapon.Damaging() == nil || !weapon.Damaging().IsRanged {
		itemName := "dominant hand"

		if weapon.Named() != nil {
			itemName = weapon.Named().Name
		}

		logMessage(fmt.Sprintf("%s wants to use %s for this, but can't!", shooterName, itemName))

		return
	}

	projectileItem := shooter.Equipped().Inventory.AtSlot(EquippedSlotOffhand)

	if projectileItem == nil {
		logMessage(fmt.Sprintf("%s doesn't have any projectiles equipped!", shooterName))

		return
	}

	if projectileItem.ProjectileData() == nil {
		projectileItemName := "off hand"

		if projectileItem.Named() != nil {
			projectileItemName = projectileItem.Named().Name
		}

		logMessage(fmt.Sprintf("%s can't use %s as ammo", shooterName, projectileItemName))

		return
	}

	distance := target.Distance(shooter.Positioned().Position)

	if distance > 12 {
		// logMessage("Can't see in the dark that far")

		return
	}

	path := engine.LinePath(
		shooter.Positioned().Position,
		target,
	)

	if path == nil {
		// logMessage("Can't shoot there, something is in the way")
		return
	}

	direction := map[engine.Position]ProjectileDirection{
		engine.PositionAt(-1, -1): ProjectileDirection_NorthWestSouthEast,
		engine.PositionAt(+1, +1): ProjectileDirection_NorthWestSouthEast,
		engine.PositionAt(-1, +1): ProjectileDirection_NorthEastSouthWest,
		engine.PositionAt(+1, -1): ProjectileDirection_NorthEastSouthWest,
		engine.PositionAt(0, +1):  ProjectileDirection_NorthSouth,
		engine.PositionAt(0, -1):  ProjectileDirection_NorthSouth,
		engine.PositionAt(-1, 0):  ProjectileDirection_EastWest,
		engine.PositionAt(+1, 0):  ProjectileDirection_EastWest,
	}[shooter.Positioned().Position.Diff(target).Sign()]

	projectile := Entity_Projectile(
		"Projectile",
		projectileItem.ProjectileData().Sprites[direction],
		tcell.StyleDefault,
		shooter,
		path,
		eventLog,
		dungeon,
	)

	turnSystem.Schedule(
		projectile.Behavior().Speed,
		projectile.Behavior().Behavior,
	)

	if projectileItem.Quantifiable() == nil {
		shooter.Equipped().Inventory.Equip(nil, EquippedSlotOffhand)
	} else {
		projectileItem.Quantifiable().CurrentQuantity--

		if projectileItem.Quantifiable().CurrentQuantity <= 0 {
			shooter.Equipped().Inventory.Equip(nil, EquippedSlotOffhand)
		}
	}

	return true
}
