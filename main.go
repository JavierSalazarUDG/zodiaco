package main

import (
	"fmt"
)

func main() {
	var month int
	var day int
	fmt.Scan(&day)
	fmt.Scan(&month)

	switch month {
	case 1:
		if day >= 21 {
			fmt.Println("Acuario")
		} else if day > 1 && day < 21 {
			fmt.Println("Capricornio")
		}
	case 2:
		if day >= 19 {
			fmt.Println("Piscis")
		} else if day > 1 && day < 19 {
			fmt.Println("Acuario")
		}
	case 3:
		if day >= 21 {
			fmt.Println("Aries")
		} else if day > 1 && day < 21 {
			fmt.Println("Piscis")
		}
	case 4:
		if day >= 21 {
			fmt.Println("Tauro")
		} else if day > 1 && day < 21 {
			fmt.Println("Aries")
		}
	case 5:
		if day >= 21 {
			fmt.Println("Géminis")
		} else if day > 1 && day < 21 {
			fmt.Println("Tauro")
		}
	case 6:
		if day >= 22 {
			fmt.Println("Cáncer")
		} else if day > 1 && day < 22 {
			fmt.Println("Géminis")
		}
	case 7:
		if day >= 23 {
			fmt.Println("Leo")
		} else if day > 1 && day < 22 {
			fmt.Println("Cáncer")
		}
	case 8:
		if day >= 23 {
			fmt.Println("Virgo")
		} else if day > 1 && day < 22 {
			fmt.Println("Leo")
		}
	case 9:
		if day >= 23 {
			fmt.Println("Libra")
		} else if day > 1 && day < 23 {
			fmt.Println("Virgo")
		}
	case 10:
		if day > 23 {
			fmt.Println("Escorpio")
		} else if day > 1 && day < 22 {
			fmt.Println("Libra")
		}
	case 11:
		if day > 23 {
			fmt.Println("Sagitario")
		} else if day > 1 && day < 22 {
			fmt.Println("Escorpio")
		}
	case 12:
		if day > 23 {
			fmt.Println("Capricornio")
		} else if day > 1 && day < 22 {
			fmt.Println("Sagitario")
		}
	}

}
