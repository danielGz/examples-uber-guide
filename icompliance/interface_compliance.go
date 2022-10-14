package icompliance

import "fmt"

type Dog interface {
	Bark()
	Play()
	//Sleep()
}

type GermanShepperd struct {}

func (g *GermanShepperd) Bark(){
	fmt.Println("ruff ruff")
}

func (g *GermanShepperd) Play(){
	fmt.Println("<some play representation I dunno>")
}

type Chihuahua struct {}

func (g *Chihuahua) Bark(){
	fmt.Println("yip yip")
}

func (g *Chihuahua) Play(){
	fmt.Println("<...>")
}

// this works but it is not necessary if you're using GermanShepperd where the interface is expected
// could be a good practice if you want to ensure a struct to implement an interface but haven't really started using it as such
var _ Dog = (*GermanShepperd)(nil)


func ScareTrespassers(dog Dog){
	dog.Bark()
	dog.Bark()
	dog.Bark()
}