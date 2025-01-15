package helloworld

const (
	spanish = "spanish"
	french  = "french"

	englishHelloPrefix = "Hello"
	spanishHelloPrefix = "Hola"
	frenchHelloPrefix  = "Bonjour"
)

func Hello(name string, lang string) string {
	if name == "" {
		name = "world"
	}
	return greetingPrefix(lang) + ", " + name
}

func greetingPrefix(lang string) string {
	prefix := ""
	switch lang {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = "Hello"
	}
	return prefix
}
