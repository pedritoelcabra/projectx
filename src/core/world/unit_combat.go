package world

func (u *Unit) GetAttackCoolDown() float64 {
	return 6000 / u.GetF(AttackSpeed)
}

func (u *Unit) PerformAttackOn(target *Unit) {
	if u.IsBusy() {
		return
	}
	u.SetF(LastCombatAction, float64(theWorld.GetTick()))
	attackSpeed := u.GetAttackCoolDown()
	u.SetBusyTimeF(attackSpeed)
	x, y := target.GetPos()
	u.QueueAttackAnimation(x, y, int(attackSpeed))
	u.StopMovement()
	//logger.General("Attacking "+target.GetName(), nil)
	attack := NewAttack()
	attack.Damage = u.GetF(AttackDamage)
	attack.Attacker = u
	attack.Defender = target
	target.ReceiveAttack(attack)
}

func (u *Unit) SetBusyTime(duration int) {
	u.SetBusyTimeF(float64(duration))
}

func (u *Unit) SetBusyTimeF(duration float64) {
	u.SetF(BusyTime, duration)
}

func (u *Unit) ReceiveAttack(attack *Attack) {
	u.SetF(LastCombatAction, float64(theWorld.GetTick()))
	u.Attributes.ApplyF(HitPoints, -attack.Damage)
	if u.GetHealth() <= 0 {
		u.Die()
	}
}

func (u *Unit) Die() {
	u.Alive = false
	//logger.General(u.GetName()+" died", nil)
	home := u.GetHome()
	if home != nil {
		home.CheckForDeceasedUnits()
	}
	if u.Work != nil {
		u.Work.Destroy()
	}
}

func (u *Unit) QueueAttackAnimation(x, y float64, speed int) {
	u.Sprite.QueueAttackAnimation((x-u.GetX())/2, (y-u.GetY())/2, speed)
	u.OrientateTowardsPoint(x, y)
}

func (u *Unit) SetToMaxHealth() {
	u.SetF(HitPoints, u.GetF(MaxHitPoints))
}

func (u *Unit) PassiveHeal() {
	if u.GetMaxHealth() <= u.GetHealth() {
		return
	}
	if theWorld.GetTick()-int(u.GetF(LastCombatAction)) < PassiveHealthCooldown {
		return
	}
	if !u.IsInOwnedSector() {
		return
	}
	u.Attributes.ApplyF(HitPoints, u.GetF(HitPoints)/100)
	if u.GetHealth() > u.GetMaxHealth() {
		u.SetToMaxHealth()
	}
}

func (u *Unit) IsInOwnedSector() bool {
	currentSector := theWorld.Grid.Tile(u.GetTileCoord()).GetSector()
	if currentSector == nil {
		return false
	}
	currentSectorFaction := currentSector.GetFaction()
	if currentSectorFaction == nil || currentSectorFaction.GetId() != FactionKey(u.Get(FactionId)) {
		return false
	}
	return true
}

func (u *Unit) GetMaxHealth() float64 {
	return u.GetF(MaxHitPoints)
}

func (u *Unit) ClosestVisibleEnemy() UnitKey {
	closestEnemy := UnitKey(-1)
	closestDistance := 999999
	for key, unit := range theWorld.GetUnits() {
		if !unit.IsAlive() {
			continue
		}
		if key == u.Id {
			continue
		}
		thisDistance := u.DistanceToUnit(unit)
		if !u.DistanceWithinVision(thisDistance) {
			continue
		}
		if !u.GetFaction().IsHostileTowards(unit.GetFaction()) {
			continue
		}
		if thisDistance < closestDistance {
			closestDistance = thisDistance
			closestEnemy = key
		}
	}
	return closestEnemy
}
