package abcutil_test

import (
	"fmt"
	"math"

	"github.com/abcsuite/abcutil"
)

func ExampleAmount() {

	a := abcutil.Amount(0)
	fmt.Println("Zero Atom:", a)

	a = abcutil.Amount(1e8)
	fmt.Println("100,000,000 Atoms:", a)

	a = abcutil.Amount(1e5)
	fmt.Println("100,000 Atoms:", a)
	// Output:
	// Zero Atom: 0 ABC
	// 100,000,000 Atoms: 1 ABC
	// 100,000 Atoms: 0.001 ABC
}

func ExampleNewAmount() {
	amountOne, err := abcutil.NewAmount(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountOne) //Output 1

	amountFraction, err := abcutil.NewAmount(0.01234567)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountFraction) //Output 2

	amountZero, err := abcutil.NewAmount(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountZero) //Output 3

	amountNaN, err := abcutil.NewAmount(math.NaN())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(amountNaN) //Output 4

	// Output: 1 ABC
	// 0.01234567 ABC
	// 0 ABC
	// invalid coin amount
}

func ExampleAmount_unitConversions() {
	amount := abcutil.Amount(44433322211100)

	fmt.Println("Atom to kCoin:", amount.Format(abcutil.AmountKiloCoin))
	fmt.Println("Atom to Coin:", amount)
	fmt.Println("Atom to MilliCoin:", amount.Format(abcutil.AmountMilliCoin))
	fmt.Println("Atom to MicroCoin:", amount.Format(abcutil.AmountMicroCoin))
	fmt.Println("Atom to Atom:", amount.Format(abcutil.AmountAtom))

	// Output:
	// Atom to kCoin: 444.333222111 kABC
	// Atom to Coin: 444333.222111 ABC
	// Atom to MilliCoin: 444333222.111 mABC
	// Atom to MicroCoin: 444333222111 μABC
	// Atom to Atom: 44433322211100 Atom
}
