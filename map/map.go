package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func mapFunc[T any,R any](list []T, f func(T, int) R) []R {
	result := []R{}

	for index,item := range list {
		result = append(result, f(item,index))
	}

	return result
}

func main()  {
	arr := []int {1,3,5,7,9}

	list := mapFunc(arr, func (item int, _ int) Person {
		return Person{
			Name: fmt.Sprintf("name%v", item),
			Age: item,
		}
	})

	for _,item := range list {
		fmt.Println(item.Name, item.Age)
	}
}