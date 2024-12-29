package functions

import (
	"fmt"
	"strconv"
	"strings"
)

func ProcessFunc(line string) string {
	chars := strings.Fields(line)
	// fmt.Println(chars)
	for i := 0; i < len(chars); i++ {
		switch {
		///////////////////////////// FUNCTIONS OF (HEX)      ///////////////////////
		case chars[i] == "(hex)":
			pass := true
			if i > 0 {
				for j := 0; j < i; j++ {
					if chars[i-j-1] != "" && pass {
						chars[i-j-1] = HextoByte(chars[i-j-1])
						chars[i] = ""
						pass = false
					}
				}
			}
			chars[i] = ""
			///////////////////////////// FUNCTIONS OF (BIN)      ///////////////////////
		case chars[i] == "(bin)":
			pass := true
			if i > 0 {
				for j := 0; j < i; j++ {
					if chars[i-j-1] != "" && pass {
						chars[i-j-1] = Bin(chars[i-j-1])
						chars[i] = ""
						pass = false
					}
				}
			}
			chars[i] = ""
			///////////////////////////// FUNCTIONS OF (CAP)      ///////////////////////
		case chars[i] == "(cap)":
			pass := true
			if i > 0 {
				for j := 0; j < i; j++ {
					if chars[i-j-1] != "" && pass {
						chars[i-j-1] = Capilize(chars[i-j-1])
						chars[i] = ""
						pass = false
					}
				}
			}
			chars[i] = ""
			///////////////////////////// FUNCTIONS OF (UP)      ///////////////////////
		case chars[i] == "(up)":
			pass := true
			if i > 0 {
				for j := 0; j < i; j++ {
					if chars[i-j-1] != "" && pass {
						chars[i-j-1] = ToUpper(chars[i-j-1])
						chars[i] = ""
						pass = false
					}
				}
			}
			chars[i] = ""
			///////////////////////////// FUNCTIONS OF (LOW)      ///////////////////////
		case chars[i] == "(low)":
			pass := true
			if i > 0 {
				for j := 0; j < i; j++ {
					if chars[i-j-1] != "" && pass {
						chars[i-j-1] = ToLower(chars[i-j-1])
						chars[i] = ""
						pass = false
					}
				}
			}
			chars[i] = ""
		// /////////////////////////// FUNCTIONS OF (LOW , N)       ///////////////////////
		case chars[i] == "(low,":
			if i+1 < len(chars) {
				word := chars[i+1]
				num, err := strconv.Atoi(word[:len(word)-1])
				count := 0
				// fmt.Println(num)
				if err != nil {
					fmt.Println("Error : invalid number")
				} else {
					for a := i; a >= 0; a-- {
						if chars[a] == "" {
							count++
						}
					}
					if num+count >= i {
						num = i - count
					}
					for j := 0; j < num+count; j++ {
						if chars[i-j-1] != "" && strings.HasSuffix(word, ")") {
							chars[i-j-1] = ToLower(chars[i-j-1])
							chars[i] = ""
							chars[i+1] = ""
						}
					}
					chars[i] = ""
					chars[i+1] = ""
				}

			}
			///////////////////////////// FUNCTIONS OF (UP , N)       ///////////////////////
		case chars[i] == "(up,":
			if i+1 < len(chars) {
				word := chars[i+1]
				num, err := strconv.Atoi(word[:len(word)-1])
				count := 0
				// fmt.Println(num)
				if err != nil {
					fmt.Println("Error : invalid number")
				} else {
					for a := i; a >= 0; a-- {
						if chars[a] == "" {
							count++
						}
					}
					if num+count >= i {
						num = i - count
					}
					for j := 0; j < num+count; j++ {
						if chars[i-j-1] != "" && strings.HasSuffix(word, ")") {
							chars[i-j-1] = ToUpper(chars[i-j-1])
							chars[i] = ""
							chars[i+1] = ""
						}
					}
					chars[i] = ""
					chars[i+1] = ""
				}
			}
			///////////////////////////// FUNCTIONS OF (CAP , N)       ///////////////////////
		case chars[i] == "(cap,":
			if i+1 < len(chars) {
				word := chars[i+1]
				num, err := strconv.Atoi(word[:len(word)-1])
				count := 0
				// fmt.Println(num)
				if err != nil {
					fmt.Println("Error : invalid number")
				} else {
					for a := i; a >= 0; a-- {
						if chars[a] == "" {
							count++
						}
					}
					if num+count >= i {
						num = i - count
					}
					for j := 0; j < num+count; j++ {
						if chars[i-j-1] != "" && strings.HasSuffix(word, ")") {
							chars[i-j-1] = Capilize(chars[i-j-1])
							chars[i] = ""
							chars[i+1] = ""
						}
					}
					chars[i] = ""
					chars[i+1] = ""
				}
			}
		}
	}
	return strings.Join(chars, " ")
}
