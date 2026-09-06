package principal

import "nut/src/component/navbar"

type Model struct {
	// Dimensiones de la ventana
	maxWidth  int
	maxHeight int

	// Nav bar
	ActiveTab navbar.Tab

	// Terminal
	EventTerminal string

	// Atributos
	Energy      int // En %
	Temperature int // En Grados C
	Oxygen      int // En %
}
