package energy

import (
	"nut/src/component/navbar"
	"strconv"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func (m ModelEnergy) View() tea.View {

	navbar := navbar.TabView(1, m.MaxWidth)

	pEngine := sPanelEngine().Render(strconv.Itoa(m.MaxWidth))
	pWeapon := sPanelEngine().Render(strconv.Itoa(m.MaxHeight))
	pShield := sPanelEngine().Render("0%")
	pLifeSupport := sPanelEngine().Render("0%")

	panels := lipgloss.JoinHorizontal(lipgloss.Top, pEngine, "", pLifeSupport)

	content := lipgloss.JoinVertical(lipgloss.Top, navbar, " ", pShield, " ", panels, " ", pWeapon)

	var v tea.View
	v.SetContent(content)
	v.AltScreen = true
	return v
}
