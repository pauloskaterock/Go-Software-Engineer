package greeting // Define um pacote chamado "greeting", que pode ser importado em outros arquivos Go.

const helloGreetings = "Hello, World!" // Declaração de uma constante de string contendo "Hello, World!".

// HelloWorld retorna uma string "Hello, World!".
// Essa função pode ser chamada em outros arquivos que importem o pacote "greeting".
func HelloWorld() string {
	return helloGreetings // Retorna o valor armazenado na constante helloGreetings.
}
