package principal

import (
	"nut/src/component/navbar"

	tea "charm.land/bubbletea/v2"
)

// Inicializa los valores devolviendo el struct directamente (igual que el ejemplo)
func InitialModelTerminal() Model {
	return Model{
		ActiveTab: navbar.Principal,

		EventTerminal: "Bienvenido al terminal",
		Energy:        0,
		Temperature:   0,
		Oxygen:        0,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
