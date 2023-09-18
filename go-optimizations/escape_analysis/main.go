package main

type counter int

func (c *counter) incrementByReference() {
	*c++
}

func (c counter) incrementByValue() counter {
	c++
	return c
}

func incrementByReference(c *counter) {
	*c++
}

func incrementByValue(c counter) *counter {
	c++
	return &c
}

func main() {
	c1 := new(counter)
	var c2 counter

	// increment c1
	c1.incrementByReference()
	*c1 = c1.incrementByValue()
	incrementByReference(c1)
	c1 = incrementByValue(*c1)

	println(*c1)

	// increment c2
	c2.incrementByReference()
	c2 = c2.incrementByValue()
	incrementByReference(&c2)
	c2 = *incrementByValue(c2)

	println(c2)
}
