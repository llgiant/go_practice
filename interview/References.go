package main

/*
func main() {
	a := false
	b := true
	c := &a
	d := &b
	fmt.Printf("value = %v, address = %p\n", a, &a)
	fmt.Printf("value = %v, address = %p\n", b, &b)
	fmt.Printf("value = %v, address = %p\n", c, &c)
	fmt.Printf("value = %v, address = %p\n", d, &d)
}
*/
/*
func changeName(name *string) {
	fmt.Printf("changeName value = %v, address = %p\n", name, &name)
}
func main() {
	name := "Hui"
	fmt.Printf("main value = %v, address = %p\n", name, &name)
	changeName(&name)
}
*/
/*
func changeName(name *string) {
	fmt.Printf("changeName value = %v, address = %p\n", name, &name)
}
func main() {
	name := "Hui"
	nameRef := &name
	fmt.Printf("main value = %v, address = %p\n", nameRef, &nameRef)
	changeName(nameRef)
}


func changeName(name *string) {
	newName := "Huila"
	name = &newName
	fmt.Printf("name %v\n", *name)
}

func main() {
	name := "Hui"
	changeName(&name)
	println(name)
}


func changeName(name *string) {
	*name = "Huila"
}

func main() {
	name := "Hui"
	changeName(&name)
	println(name)
}


type Person struct {
	name string
	age  int
}

func (p *Person) ChangeName(name string) {
	p.name = name
}
func main() {
	p := Person{"Bob", 25}
	p.ChangeName("Bob Dilan")
	fmt.Println(p)
}


type Address struct {
	city   string
	street string
	house  int
}

func (a *Address) setCity(city string) {
	a = &Address{
		city: city,
	}
}
func (a *Address) setStreet(street string) {
	a.street = street
}
func setHouse(addr *Address, house int) {
	addr = &Address{
		house: house,
	}
}
func main() {
	addr := Address{
		"New York",
		"Wall",
		10,
	}
	addr.setCity("London")
	addr.setStreet("Picadilla")
	setHouse(&addr, 5)
	fmt.Println(addr)
}


func main() {
	p := getSlice()
	fmt.Println(p)
}

func getSlice() int[] {
	s := []int{1, 2, 3}
	defer changeSlice(s)
	return s
}

func changeSlice(s []int) {
	s[1] = 10
}


func named() (a, b int) {
	a, b = 1, 2
	defer func() {
		a = 10
		b = 20
	}()
	return a, b
}
func unnamed() (int, int) {
	a, b := 1, 2

	defer func() {
		a = 10
		b = 20
	}()
	return a, b
}
func main() {
	a, b := named()
	fmt.Println(a, b)
	a, b = unnamed()
	fmt.Println(a, b)
}
*/
func main() {

}
