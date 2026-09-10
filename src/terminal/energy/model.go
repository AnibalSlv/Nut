package energy

type ModelEnergy struct {
	// Dimensiones de la ventana
	MaxWidth  int
	MaxHeight int

	// Atributos de la ventana
	energyTotal       int
	energyLifeSupport int
	energyEngine      int
	energyShield      int
	energyWeapons     int
}
