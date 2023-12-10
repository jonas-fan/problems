// #array

type ProductOfNumbers struct {
    products []int
}


func Constructor() ProductOfNumbers {
    return ProductOfNumbers{
        products: []int{},
    }
}


func (this *ProductOfNumbers) Add(num int)  {
    for i := range this.products {
        this.products[i] *= num
    }

    this.products = append(this.products, num)
}


func (this *ProductOfNumbers) GetProduct(k int) int {
    return this.products[len(this.products)-k]
}


/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
