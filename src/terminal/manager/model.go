package manager

import (
	"nut/src/terminal/energy"
	"nut/src/terminal/principal"
)

type view int

const (
	viewPrincipal view = iota
	viewEnergy
	viewHealt
	viewMap
)

type managerModel struct {
	currentView    view
	principalModel principal.Model
	energyModel    energy.ModelEnergy
}
