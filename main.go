package main

import (
	"fmt"
	"math"
	"regexp"
	"strings"
)

// 1. Buat sebuah fungsi pangkat, yang dimana ada 2 parameter.
//     Contoh:
//     - 2 pangkat 3 = 8
//     - 3 pangkat 3 = 27
//     - 5 pangkat -1 = 0,5 (? -> 0.2)

// 2. Buat sebuah fungsi konversi dari CamelCase menjadi snake_case
//     Contoh:
//     - SatuDuaTiga = satu_dua_tiga

type power struct {
	x int
	y int
}

func main() {

	powCases := []power{
		{2, 3},
		{3, 3},
		{5, -1},
	}

	scCases := []string{
		"SatuDuaTiga",
		"dua_tiga_Empat",
		"dua__tiga_Empat",
		"_Dua_tiga_Empat_",
		"____dua_tiga__Empat_____",
	}

	fmt.Println("\n+--------------+")
	fmt.Println("Pangkat (Power)")
	fmt.Println("+--------------+")
	for i, v := range powCases {
		fmt.Println(fmt.Sprintf("%d.", i+1), "Manual :", pow(float64(v.x), v.y), "|", "With math.Pow() :", Pow(float64(v.x), float64(v.y)))
	}

	fmt.Println("\n+----------------------+")
	fmt.Println("Camel Case to Snack Case")
	fmt.Println("+----------------------+")
	for i, v := range scCases {
		fmt.Println(fmt.Sprintf("%d.", i+1), "Manual :", toSnakeCase(v), "|", "With Regex :", ToSnakeCase(v))
	}
}

// With Manual -> easy to understand, but does not support powering by float number)
// Algorithm pow(x float64, y int) :
// if (y == 0) => 1
// if (y == even) => pow(x, y / 2) * pow(x, y / 2)
// else (y == odd)
//     if (y > 0) => x * pow(x, y / 2) * pow(x, y / 2)
//     else (y < 0) => pow(x, y / 2) * pow(x, y / 2) / x
func pow(x float64, y int) float64 {
	if y == 0 {
		return 1
	}
	temp := pow(x, y/2)
	if y%2 == 0 {
		return temp * temp
	} else {
		if y > 0 {
			return x * temp * temp
		} else {
			return temp * temp / x
		}
	}
}

// With built in library math Pow -> More complex algorithm (support powering by float64)
func Pow(x, y float64) float64 {
	return math.Pow(x, y)
}

// Manual
func toSnakeCase(s string) string {
	const (
		upper  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		lower  = "abcdefghijklmnopqrstuvwxyz"
		uscore = '_'
		empty  = ""
	)
	rs, urunes, lrunes := []rune{}, []rune(upper), []rune(lower)
	var merge = true

	for idx, val := range s {
		for i, v := range urunes {
			if val == uscore {
				break
			}
			if val == v {
				if idx != 0 {
					if s[idx-1] != uscore {
						rs = append(rs, uscore)
					}
				}
				rs = append(rs, lrunes[i])
				merge = false
				break
			}
		}
		if merge {
			rs = append(rs, val)
		}
		merge = true
	}

	// Trim first underscore
	for {
		if rs[0] != uscore {
			break
		}
		rs = rs[1:]
	}

	// Trim last underscore
	for {
		if rs[len(rs)-1] != uscore {
			break
		}
		rs = rs[:len(rs)-1]
	}

	return string(rs)
}

// With Regex
func ToSnakeCase(str string) string {
	const (
		uscore = '_'
	)

	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")

	rs := []rune(strings.ToLower(snake))

	// Trim first underscore
	for {
		if rs[0] != uscore {
			break
		}
		rs = rs[1:]
	}

	// Trim last underscore
	for {
		if rs[len(rs)-1] != uscore {
			break
		}
		rs = rs[:len(rs)-1]
	}
	return string(rs)
}
