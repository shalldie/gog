package each



func Each(list []T,f func(T) bool )  {
	for i,item := range list{
		f(i,item)
	}
}