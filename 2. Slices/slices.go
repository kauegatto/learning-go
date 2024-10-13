package main

import "fmt"

func main() {
	// slices são declarados de forma similar à arrays, mas o colchete não especifica o tamanho
	var s []int
	s = append(s, 1) // 1 elemento, 1 de capacidade
	// Slices são como "arrays infinitos" - Mas por baixo dos panos, ainda são arrays:
	// Arrays são espaços consecutivos na memória, portanto, tem tamanho fixo
	// Slices funcionam como o ArrayList, dobrando de tamanho quando sua cabacidade ultrapassa um limite especificado
	inform_slice(s)
	s = append(s, 1)
	s = append(s, 1)
	s = append(s, 1)
	inform_slice(s) // 4 elementos, 4 de capacidade
	s = append(s, 1)
	inform_slice(s) // 5 elementos, 8 de capacidade

	s = s[:4]
	inform_slice(s) // 4 elementos, 8 de capacidade
	s = s[:3]
	inform_slice(s) // 3 elementos, 8 de capacidade

	// dica: se você sabe que seu slice tem ~100 espaços, melhor jogar sua capacidade inicial próximo disso (ou até um pouco maior)
	// isso pode fazer com que perca um pouco de espaço, mas evita a criação de um novo array, com o dobro do que você especificou no início.
	// te resulta mais performance e memória, no fim das contas.
}

func inform_slice(s []int) {
	fmt.Printf("tamanho: %d len(s), capacidade: %d\n", len(s), cap(s))
}
