package hello

//TestVersion is the version of the unit test that this will pass
const TestVersion = 1

/*HelloWorld greets you or the world if you don't have a name.*/
func HelloWorld(name string) string {
	if name == "" {
		name = "World"
	}
	return "Hello, " + name + "!"
}
