package main

// import "fmt"

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

func (p *Person)  changeName(newName string) {
	p.Name = newName
}

func main() {
	myPerson := NewPerson("Ramy", 23)
	myPerson.changeName("Youssef")

	a :=7
	b := &a
	*b = 9

	// mySlice  := []int{1,2,3,4,5}

	// for index, _ := range mySlice {
	// 	mySlice[index]++;
	// }

	// fmt.Println(mySlice)
	// fmt.Printf("Hello, %s %d", myPerson.Name, myPerson.Age)


}