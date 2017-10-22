package main

func ifErr(operation error) {
	if operation != nil {
		panic(operation)
	}
}

func onTheGuestList(album Album, visitor string) bool {
	allowed := false
	for _, guest := range album.Guest {
		if guest.Hex() == visitor {
			allowed = true
			break
		}
	}
	return allowed
}
