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
	// Zero Atom: 0 DCR
	// 100,000,000 Atoms: 1 DCR
	// 100,000 Atoms: 0.001 DCR
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

	// Output: 1 DCR
	// 0.01234567 DCR
	// 0 DCR
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
	// Atom to kCoin: 444.333222111 kDCR
	// Atom to Coin: 444333.222111 DCR
	// Atom to MilliCoin: 444333222.111 mDCR
	// Atom to MicroCoin: 444333222111 μDCR
	// Atom to Atom: 44433322211100 Atom
}
