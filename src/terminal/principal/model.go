package principal

type Model struct {
	// Dimensiones de la ventana
	MaxWidth  int
	MaxHeight int

	// Terminal
	EventTerminal string

	// Atributos
	Energy      int // En %
	Temperature int // En Grados C
	Oxygen      int // En %
}
