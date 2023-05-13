package main

import "fmt"

func formatCurrency(n int) string {
	f := float32(n / 100)
	return fmt.Sprintf("$%.2f", f)
}
