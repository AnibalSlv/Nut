package manager

import (
	"nut/src/terminal/energy"
	"nut/src/terminal/principal"

	tea "charm.land/bubbletea/v2"
)

func InitialModelManager() managerModel {
	return managerModel{
		currentView:    viewPrincipal,
		principalModel: principal.InitialModelTerminal(),
		energyModel:    energy.InitialModelEnergy(),
	}
}

func (m managerModel) Init() tea.Cmd {
	return nil
}
