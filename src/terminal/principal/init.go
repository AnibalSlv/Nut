package principal

import (
	tea "charm.land/bubbletea/v2"
)

// Inicializa los valores devolviendo el struct directamente (igual que el ejemplo)
func InitialModelTerminal() Model {
	return Model{
		EventTerminal: "Bienvenido al terminal",
		Energy:        0,
		Temperature:   0,
		Oxygen:        0,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
