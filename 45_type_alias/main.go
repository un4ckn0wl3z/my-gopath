package main

func main() {
	var (
		byteVal  byte
		uint8Val uint8
		intVal   int
	)
	uint8Val = byteVal

	var (
		runeVal  rune
		int32Val int32
	)

	runeVal = int32Val
	_ = uint8Val
	_ = intVal
	_ = runeVal

}
