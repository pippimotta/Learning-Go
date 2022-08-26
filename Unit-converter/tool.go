package main

import "fmt"

type memory struct {
	converter *converter
	from      float64
	to        float64
}

type tool struct {
	converters []*converter

	memories []*memory
}

func (m *memory) String() string {
	return fmt.Sprintf("%.2f[%s] -> %.2f[%s]", m.from, m.converter.fromUnit, m.to, m.converter.toUnit)
}
func (t *tool) start() {
	for {

		t.printMemories()
		no, ok := t.inputConverterNo()
		if !ok {
			return
		}

		c := t.converters[no]
		from := t.inputValue(c)
		to := c.calc(from)
		fmt.Printf("%.2f[%s] -> %.2f[%s]\n\n", from, c.fromUnit, to, c.toUnit)

	}
}

func (t *tool) printMemories() {
	if len(t.memories) < 1 {
		return
	}

	fmt.Println("History")
	for _, m := range t.memories {
		fmt.Println(m.String())
	}
	fmt.Println()
}

func (t *tool) inputConverterNo() (int, bool) {
	fmt.Println("Please select a converter.")
	for i, c := range t.converters {
		fmt.Printf("%d: %s\n", i+1, c.name)
	}
	fmt.Printf("%d: End\n", len(t.converters)+1)

	var no int
	for no <= 0 || no > len(t.converters)+1 {
		fmt.Printf("Please select from 1〜%d>", len(t.converters)+1)
		fmt.Scanln(&no)
	}

	// end
	if no == len(t.converters)+1 {

		return 0, false
	}

	return no - 1, true
}

func (t *tool) inputValue(c *converter) float64 {
	var v float64
	fmt.Println(c.name)
	fmt.Printf("Convert from [%s]>", c.fromUnit)
	fmt.Scanln(&v)
	return v
}
