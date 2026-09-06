package principal

import (
	"nut/src/component/navbar"

	tea "charm.land/bubbletea/v2"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.maxWidth = msg.Width
		m.maxHeight = msg.Height

	case tea.KeyPressMsg:

		switch msg.String() {

		case "ctr+c", "q":
			return m, tea.Quit

		// Logica de la barra de navegacion
		case "1", "2", "3", "4":
			m.ActiveTab = navbar.TabUpdate(msg.String())

		}
	}
	return m, nil
}
