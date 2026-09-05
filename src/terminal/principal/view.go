package principal

import (
	"fmt"
	"strconv"

	tea "charm.land/bubbletea/v2"
)

func (m Model) View() tea.View {
	pEnergy := PanelEnergy().Render(strconv.Itoa(m.Energy))

	s := fmt.Sprintf("%s", pEnergy)
	return tea.NewView(s)
}
