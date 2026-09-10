package energy

import (
	tea "charm.land/bubbletea/v2"
)

func (m ModelEnergy) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyPressMsg:

		switch msg.String() {

		case "ctr+c", "q":
			return m, tea.Quit

		}
	}
	return m, nil
}
