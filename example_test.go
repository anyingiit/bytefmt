package bytefmt_test

import (
	"fmt"

	"code.cloudfoundry.org/bytefmt"
)

func ExampleByteSize() {
	fmt.Println(bytefmt.ByteSize(100.5 * bytefmt.MEGABYTE))
	fmt.Println(bytefmt.ByteSize(uint64(1024)))
	fmt.Println(bytefmt.ByteSize(0))
	// Output:
	// 100.5M
	// 1K
	// 0B
}

func ExampleToBytes() {
	b, err := bytefmt.ToBytes("1.5GiB")
	fmt.Println(b, err)

	_, err = bytefmt.ToBytes("12")
	fmt.Println(err)
	// Output:
	// 1610612736 <nil>
	// byte quantity must be a positive integer with a unit of measurement like M, MB, MiB, G, GiB, or GB
}

func ExampleToMegabytes() {
	m, err := bytefmt.ToMegabytes("2g")
	fmt.Println(m, err)
	// Output:
	// 2048 <nil>
}
