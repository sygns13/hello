package greet

//variables a nivel de packetes no se pueden crear con la asignacion corta
var emoji = "🙋‍"

//a nivel global exportable primera letra con mayuscula en este caso funciones
var Emoji2 = "🙋‍"

// English retorna saludo en ingles
func English() string {
	return "Hi " + emoji
}

// Italian retorna saludo en italiano
func Italian() string {
	return "Ciao " + emoji
}
