package main

import (
	"fmt"
	"strings"
)

func main() {
	palavra := "golang"
	tentativas := 6
	contagem_tentativas := 0
	letras_adivinhadas := make([]string, 0)

	fmt.Println("Bem-vindo ao Jogo da Forca!")

	for contagem_tentativas < tentativas {
		fmt.Printf("Palavra: %s\n", exibePalavra(palavra, letras_adivinhadas))
		fmt.Printf("Tentativas restantes: %d\n", tentativas - contagem_tentativas)
		fmt.Print("Digite uma letra: ")
		var tentativa string
		fmt.Scanf("%s", &tentativa)
		tentativa = strings.ToLower(tentativa)

		if len(tentativa) != 1 || !isAlpha(tentativa) {
			fmt.Println("Digite uma unica letra valida!")
			continue
		}

		if !contains(letras_adivinhadas, tentativa) {
			letras_adivinhadas = append(letras_adivinhadas, tentativa)
			if !strings.Contains(palavra, tentativa) {
				tentativas++
			}
		}

		if palavraAdivinhada(palavra, letras_adivinhadas) {
			fmt.Println("Parabens, voce ganhou! a palavra era:", palavra)
			return
		}

		fmt.Println("Fim de jogo! voce perdeu. A palavra era:", palavra)
	}
}

func isAlpha(s string) bool {
	return len(s) == 1 && 'a' <= s[0] && s[0] <= 'z'
}

func contains(arr []string, item string) bool {
	for _, val := range arr {
		if val == item {
			return true
		}
	}
	return false
}

func exibePalavra(txt string, letras_adivinhadas []string) string {
	display := make([]string, len(txt))
	for i, char := range txt {
		if contains(letras_adivinhadas, string(char)) {
			display[i] = string(char)
		} else {
			display[i] = "_"
		}
	}
	return strings.Join(display, " ")
}

func palavraAdivinhada(txt string, letras_adivinhadas []string) bool {
	for _, char := range txt {
		if !contains(letras_adivinhadas, string(char)) {
			return false
		}
	}
	return true
}
