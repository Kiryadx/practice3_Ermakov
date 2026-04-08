package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Kiryadx/practice3_Ermakov/internal/calculator"
)

func main() {
	var (
		a       = flag.Float64("a", 0, "первое число")
		b       = flag.Float64("b", 0, "второе число")
		op      = flag.String("op", "", "операция (add, sub, mul, div, pow)")
		logging = flag.Bool("log", false, "включить логирование")
		help    = flag.Bool("h", false, "вывести справку")
	)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Калькулятор с плагинами\n")
		fmt.Fprintf(os.Stderr, "Использование: %s [опции]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Опции:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nПримеры:\n")
		fmt.Fprintf(os.Stderr, "  %s -a 10 -b 5 -op add\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -a 10 -b 0 -op div\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -a 2 -b 3 -op pow -log\n", os.Args[0])
	}

	flag.Parse()

	if *help || flag.NFlag() == 0 {
		flag.Usage()
		os.Exit(0)
	}

	if *op == "" {
		fmt.Fprintf(os.Stderr, "Ошибка: не указана операция (-op)\n")
		flag.Usage()
		os.Exit(1)
	}

	calc := calculator.NewCalculator(*logging)

	result, err := calc.Calculate(*op, *a, *b)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Результат: %.2f %s %.2f = %.2f\n", *a, *op, *b, result)
}