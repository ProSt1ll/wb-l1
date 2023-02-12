package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	a := new(big.Int)
	b := new(big.Int)
	a.Set(big.NewInt(int64(math.Pow(2, 30))))
	b.Set(big.NewInt(int64(math.Pow(2, 25))))
	//сложение
	fmt.Println(Add(a, b))
	//вычитание
	fmt.Println(Sub(a, b))
	//умножение
	fmt.Println(Mul(a, b))
	//деление
	fmt.Println(Div(a, b))
}

func Add(a *big.Int, b *big.Int) *big.Int {
	return new(big.Int).Add(big.NewInt(a.Int64()), big.NewInt(b.Int64()))
}

func Sub(a *big.Int, b *big.Int) *big.Int {
	return new(big.Int).Sub(big.NewInt(a.Int64()), big.NewInt(b.Int64()))
}

func Mul(a *big.Int, b *big.Int) *big.Int {
	return new(big.Int).Mul(big.NewInt(a.Int64()), big.NewInt(b.Int64()))
}

func Div(a *big.Int, b *big.Int) *big.Int {
	return new(big.Int).Div(big.NewInt(a.Int64()), big.NewInt(b.Int64()))
}
