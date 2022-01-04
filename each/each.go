package main

import "fmt"

type Person struct {
	Name string
	Age int
}

// FIXME: Each
func each[T any](list []T, f func(T,int)) {
	for index, item := range list {
		f(item,index)
	}
}

func main() {
	arr := []string{"hello", "world"}

	each(arr, func( item string,index int) {
		fmt.Printf("index is %v, item is %v \n", index, item)
	})



	list := []Person{
		Person{Name: "tom",Age: 12},
		Person{Name: "lily",Age: 23},
	}

	each(list,func (p Person,_ int)  {
		fmt.Printf("name is %v, age is %v \n",p.Name,p.Age)
	})
}
