package main

func main() {
	// Array tem espaço fixo :), se comporta como um array no c

	var meuArray [3]int
	meuArray[0] = 2
	meuArray[1] = 4
	meuArray[2] = 6

	for i := 0; i < len(meuArray); i++ {
		println(meuArray[i])
	}

	println("for range")

	for _, v := range meuArray {
		println(v)
	}
}
