package ptrs_interfaces

import "fmt"

type Fooer interface {
	ModifyMe()
	GetValue() string
}

type Foo struct {
	Name string
}

func (f *Foo) ModifyMe() {
	f.Name = "NEW"
}

func (f *Foo) GetValue() string {
	return f.Name
}

/*
implementing the interface without a pointer will not allow you to modify the underlying data of the caller
func (f Foo) ModifyMe() {
	f.Name = "NEW"
}

func (f Foo) GetValue() string {
	return f.Name
}
*/

func Demo(){
	var foox Fooer
	//this doesn't work
	/*
		var foo Fooer
		foo = Foo{}
	*/
	// but this does, in this case you don't need to create an actual pointer to be able to treat
	foo := Foo{}
	foox = &Foo{}
	foox.ModifyMe()

	foo.ModifyMe()
	fmt.Printf("foox %s \n", foox.GetValue())
	fmt.Printf("foo %s \n", foo.Name)
}
