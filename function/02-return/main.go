package main

func sayHello(name string) string {
	return "Hello " + name
}

func main() {
	res := sayHello("Psyon")
	println(res)
}
