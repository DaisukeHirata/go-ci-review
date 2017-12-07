package cart

import "testing"

func TestAddAndGetProductsInCarrt(t *testing.T) {
       c := New()
       c.Add("りんご")
       c.Add("みかん")

       products := c.GetAll()
       if len(products) != 2 {
        t.Fatalf("商品の数が想定と違う。(商品数 :%d)", len(products))
       }
       if products[0] != "りんご" && products[1] != "りんご" {
        t.Error("りんごがカートに入っていない。")
        t.Log("カートの中身：", products)
       }
       if products[0] != "みかん" && products[1] != "みかん" {
        t.Error("みかんがカートに入っていない。")
        t.Log("カートの中身：", products)
       }

}

