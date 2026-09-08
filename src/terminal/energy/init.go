package energy

import (
	tea "charm.land/bubbletea/v2"
)

// Inicializa los valores devolviendo el struct directamente (igual que el ejemplo)
func InitialModelEnergy() ModelEnergy {
	return ModelEnergy{

		energyTotal:   0,
		energyHealt:   0,
		energyEngine:  0,
		energyShield:  0,
		energyWeapons: 0,
	}
}

func (m ModelEnergy) Init() tea.Cmd {
	return nil
}
