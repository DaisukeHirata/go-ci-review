package main

//import "errors"

//func examle(code string) (int,err) {
//    if code == "hoge" {
//        return 1,nil
//    }
//    return 0,errors.New("code must be hoge")
//}

//func main() {
//	New()
//}//

type Cart struct {
	products []string
}

func New() *Cart {
	c := new(Cart)
	c.products = make([]string, 0)
	return c
}

func (c *Cart) Add(s string) {
	c.products = append(c.products, s)
}

func (c *Cart) GetAll() []string {
	return c.products
}
