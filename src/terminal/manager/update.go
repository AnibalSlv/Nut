package manager

import (
	"nut/src/terminal/energy"
	"nut/src/terminal/principal"

	tea "charm.land/bubbletea/v2"
)

func (m managerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.maxWidth = msg.Width
		m.maxHeight = msg.Height

	case tea.KeyPressMsg:

		switch msg.String() {

		case "ctrl+c", "i":
			return m, tea.Quit

			// Logica de la barra de navegacion
		case "1":
			m.currentView = viewPrincipal
			return m, nil
		case "2":
			m.currentView = viewEnergy

			// Pasa el tamano de la pantalla al modelo
			m.energyModel.MaxWidth = m.maxWidth
			m.energyModel.MaxHeight = m.maxHeight
			return m, nil
		}
	}

	switch m.currentView {
	case viewPrincipal:
		var cmd tea.Cmd
		updated, cmd := m.principalModel.Update(msg)
		m.principalModel = updated.(principal.Model)
		return m, cmd

	case viewEnergy:
		m.energyModel.MaxWidth = m.maxWidth
		m.energyModel.MaxHeight = m.maxHeight

		updated, cmd := m.energyModel.Update(msg)
		m.energyModel = updated.(energy.ModelEnergy)
		return m, cmd
	}

	return m, nil
}
