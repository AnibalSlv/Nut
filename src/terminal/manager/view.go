package manager

import tea "charm.land/bubbletea/v2"

func (m managerModel) View() tea.View {
	switch m.currentView {
	case viewPrincipal:
		return m.principalModel.View()

	case viewEnergy:
		return m.energyModel.View()

	default:
		var v tea.View
		v.SetContent("[Error] Pantalla desconocida")
		return v
	}
}
