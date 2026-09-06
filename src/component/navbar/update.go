package navbar

func TabUpdate(key string) Tab {
	switch key {
	case "1":
		return Principal
	case "2":
		return Energy
	case "3":
		return Healt
	case "4":
		return Map

	}
	return Principal
}
