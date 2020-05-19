package utils

// Função para facilitar o tratamento do erro
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
