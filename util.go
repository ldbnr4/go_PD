package main

func ifErr(operation error) {
	if operation != nil {
		panic(operation)
	}
}
