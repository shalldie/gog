package main

type Person struct {
	Name string
	Age int
}

func filter[T any](list []T, f func(T, int) bool) []T {
	result := []T{}

	for index, item := range list {
		if f(item, index) {
			result = append(result, item)
		}
	}

	return result
}

func main()  {
	list := []int{1,3,5,7,9}

	list = filter(list,func(item int, _ int) bool {
		return item >= 5
	})

	for _,n := range list {
		println(n)
	}

	list2 := []Person{
		Person{Name: "tom",Age: 12},
		Person{Name: "lily",Age: 23},
	}

	list2 = filter(list2, func (item Person,_ int) bool {
		return item.Age>18
	})

	for _, n := range list2{
		println(n.Name, n.Age)
	}
}