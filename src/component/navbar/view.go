package navbar

import "charm.land/lipgloss/v2"

/*
Contiene toda la logica de visualizacion del navbar
  - Tiene como altura: 5
  - Numero en las vistas:
  - 0 = Principal
  - 1 = Energy
  - 2 = Healt
  - 3 = Map
*/
func TabView(activeTab Tab, maxWidth int) string {
	var tabs []string

	titles := []string{"Principal", "Energia", "Salud", "Mapa"}

	for i, title := range titles {
		if Tab(i) == activeTab {
			tabs = append(tabs, sActiveTab().Render(title))
		} else {
			tabs = append(tabs, sInactiveTab().Render(title))
		}
	}

	navbar := lipgloss.JoinHorizontal(lipgloss.Top, sNavBar(maxWidth).
		Render(lipgloss.JoinHorizontal(lipgloss.Top, tabs...)))

	return navbar
}
