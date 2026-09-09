package energy

import (
	"nut/src/component/navbar"
	"strconv"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func (m ModelEnergy) View() tea.View {

	navbar := navbar.TabView(1, m.MaxWidth)

	pShield := sPanelEngine().Render(strconv.Itoa(m.energyShield) + "%")
	panelUp := lipgloss.Place(m.MaxWidth, 3, lipgloss.Center, lipgloss.Center,
		lipgloss.JoinVertical(lipgloss.Top, "Escudos", pShield))

	pEngine := sPanelEngine().Render(strconv.Itoa(m.energyEngine) + "%")
	pLifeSupport := sPanelEngine().Render(strconv.Itoa(m.energyLifeSupport) + "%")

	// Para poder separar los dos paneles es: el tamano maximo - (el tamano del cuadro * 2)
	separatorWidht := lipgloss.NewStyle().Width(m.MaxWidth - (50 * 2)).Render("")
	panelMid := lipgloss.JoinHorizontal(lipgloss.Top,
		lipgloss.JoinVertical(
			lipgloss.Top, "Motores", pEngine),
		separatorWidht,
		lipgloss.JoinVertical(
			lipgloss.Top, "Soporte Vital", pLifeSupport))

	pWeapon := sPanelEngine().Render(strconv.Itoa(m.energyWeapons) + "%")
	panelDown := lipgloss.Place(m.MaxWidth, 3, lipgloss.Center, lipgloss.Center,
		lipgloss.JoinVertical(lipgloss.Top, "Armas", pWeapon))

	separatorHeight := lipgloss.NewStyle().Height(2).Render("")
	content := lipgloss.JoinVertical(lipgloss.Top, navbar, separatorHeight, panelUp, separatorHeight, panelMid, separatorHeight, panelDown)

	var v tea.View
	v.SetContent(content)
	v.AltScreen = true
	return v
}
