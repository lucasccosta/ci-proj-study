package main

import "testing"

func TestSomaPositives(t *testing, T){

	total := SomaPositives(2, 5)

	if total < 0 {
		t.Error("Resultado da soma deve ser um valor positivo")
	}
}