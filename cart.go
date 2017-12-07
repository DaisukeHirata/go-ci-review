package cart

type Cart struct {
       products []string
}

func New() *Cart {
       c := new(Cart)
       c.products = make ([]string, 0)
       return c
}

func (c *Cart) Add(s string) {
       c.products = append(c.products, s)
}

func (c *Cart) GetAll() []string {
       return c.products
}

