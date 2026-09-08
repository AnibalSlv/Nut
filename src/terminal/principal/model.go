package principal

type Model struct {
	// Dimensiones de la ventana
	maxWidth  int
	maxHeight int

	// Terminal
	EventTerminal string

	// Atributos
	Energy      int // En %
	Temperature int // En Grados C
	Oxygen      int // En %
}
