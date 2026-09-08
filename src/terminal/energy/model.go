package energy

type ModelEnergy struct {
	// Dimensiones de la ventana
	MaxWidth  int
	maxHeight int

	// Atributos de la ventana
	energyTotal   int
	energyHealt   int
	energyEngine  int
	energyShield  int
	energyWeapons int
}
