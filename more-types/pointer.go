package main
import "fmt"

func main() {
	i, j := 44, 32

	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
  *p = *p / 37
  fmt.Println(j)
}
