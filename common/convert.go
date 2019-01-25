package common

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"math/big"
	"strconv"
	"strings"
	// "github.com/pingcap/tidb/types"
	// "github.com/pingcap/tidb/types"
)

var tenToAny map[int64]string = map[int64]string{0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9", 10: "a", 11: "b", 12: "c", 13: "d", 14: "e", 15: "f", 16: "g", 17: "h", 18: "i", 19: "j", 20: "k", 21: "l", 22: "m", 23: "n", 24: "o", 25: "p", 26: "q", 27: "r", 28: "s", 29: "t", 30: "u", 31: "v", 32: "w", 33: "x", 34: "y", 35: "z", 36: ":", 37: ";", 38: "<", 39: "=", 40: ">", 41: "?", 42: "@", 43: "[", 44: "]", 45: "^", 46: "_", 47: "{", 48: "|", 49: "}", 50: "A", 51: "B", 52: "C", 53: "D", 54: "E", 55: "F", 56: "G", 57: "H", 58: "I", 59: "J", 60: "K", 61: "L", 62: "M", 63: "N", 64: "O", 65: "P", 66: "Q", 67: "R", 68: "S", 69: "T", 70: "U", 71: "V", 72: "W", 73: "X", 74: "Y", 75: "Z"}

// Decimal to binary
func DecBin(n int64) string {
	if n < 0 {
		log.Println("Decimal to binary error: the argument must be greater than zero.")
		return ""
	}
	if n == 0 {
		return "0"
	}
	s := ""
	for q := n; q > 0; q = q / 2 {
		m := q % 2
		s = fmt.Sprintf("%v%v", m, s)
	}
	return s
}

// Decimal to octal
func DecOct(d int64) int64 {
	if d == 0 {
		return 0
	}
	if d < 0 {
		log.Println("Decimal to octal error: the argument must be greater than zero.")
		return -1
	}
	s := ""
	for q := d; q > 0; q = q / 8 {
		m := q % 8
		s = fmt.Sprintf("%v%v", m, s)
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Println("Decimal to octal error:", err.Error())
		return -1
	}
	return int64(n)
}

// Decimal to hexadecimal
func DecHex(n int64) string {
	if n < 0 {
		log.Println("Decimal to hexadecimal error: the argument must be greater than zero.")
		return ""
	}
	if n == 0 {
		return "0"
	}
	hex := map[int64]int64{10: 65, 11: 66, 12: 67, 13: 68, 14: 69, 15: 70}
	s := ""
	for q := n; q > 0; q = q / 16 {
		m := q % 16
		if m > 9 && m < 16 {
			m = hex[m]
			s = fmt.Sprintf("%v%v", string(m), s)
			continue
		}
		s = fmt.Sprintf("%v%v", m, s)
	}
	return s
}

// Binary to decimal
func BinDec(b string) (n int64) {
	s := strings.Split(b, "")
	l := len(s)
	i := 0
	d := float64(0)
	for i = 0; i < l; i++ {
		f, err := strconv.ParseFloat(s[i], 10)
		if err != nil {
			log.Println("Binary to decimal error:", err.Error())
			return -1
		}
		d += f * math.Pow(2, float64(l-i-1))
	}
	return int64(d)
}

// Octal to decimal
func OctDec(o int64) (n int64) {
	s := strings.Split(strconv.Itoa(int(o)), "")
	l := len(s)
	i := 0
	d := float64(0)
	for i = 0; i < l; i++ {
		f, err := strconv.ParseFloat(s[i], 10)
		if err != nil {
			log.Println("Octal to decimal error:", err.Error())
			return -1
		}
		d += f * math.Pow(8, float64(l-i-1))
	}
	return int64(d)
}

// Hexadecimal to decimal
func HexDec(h string) (n int64) {
	h = strings.Replace(h, "0x", "", -1)
	s := strings.Split(strings.ToUpper(h), "")
	l := len(s)
	i := 0
	d := float64(0)
	hex := map[string]string{"A": "10", "B": "11", "C": "12", "D": "13", "E": "14", "F": "15"}
	for i = 0; i < l; i++ {
		c := s[i]
		if v, ok := hex[c]; ok {
			c = v
		}
		f, err := strconv.ParseFloat(c, 10)
		if err != nil {
			log.Println("Hexadecimal to decimal error:", err.Error())
			return -1
		}
		d += f * math.Pow(16, float64(l-i-1))
	}
	return int64(d)
}

// Octal to binary
func OctBin(o int64) string {
	d := OctDec(o)
	if d == -1 {
		return ""
	}
	return DecBin(d)
}

// Hexadecimal to binary
func HexBin(h string) string {
	d := HexDec(h)
	if d == -1 {
		return ""
	}
	return DecBin(d)
}

// Binary to octal
func BinOct(b string) int64 {
	d := BinDec(b)
	if d == -1 {
		return -1
	}
	return DecOct(d)
}

// Binary to hexadecimal
func BinHex(b string) string {
	d := BinDec(b)
	if d == -1 {
		return ""
	}
	return DecHex(d)
}

// func RealValue(value string, unit string) string {
// 	var valueDec, unitDec, resultDec decimal.Decimal
// 	unitUint, _ := strconv.ParseUint(unit, 10, 64)
// 	unitStr := strconv.FormatUint(Exponent(10, unitUint), 10)
// 	valueDec, _ = decimal.NewFromString(value)
// 	unitDec, _ = decimal.NewFromString(unitStr)
// 	resultDec = valueDec.Div(unitDec)
// 	valueDec.Div(unitDec)
// 	return string(resultDec.String())
// }

func RealValue(value string, unit string) string {
	unitInt, _ := strconv.ParseInt(unit, 10, 64)
	unit = BigIntPow(10, unitInt).String()
	valueDecBig := new(big.Int)
	valueDecBig.SetString(value, 10)
	unitUintBig := new(big.Int)
	unitUintBig.SetString(unit, 10)
	return BigIntDiv(valueDecBig, unitUintBig)
}

//精确获取两个大整型数据相除的结果
func BigIntDiv(aV, b *big.Int) string {
	bigA := big.NewInt(0)
	ltZero := false
	if aV.Cmp(big.NewInt(0)) == -1 {
		bigA = big.NewInt(0).Abs(aV)
		ltZero = true
	} else {
		bigA = aV
	}
	ip := big.NewInt(1)
	r := ip.Div(bigA, b)

	ip = big.NewInt(1)
	c := ip.Mul(r, b)

	ip = big.NewInt(1)
	d := ip.Sub(bigA, c)
	e := d.Cmp(big.NewInt(0))
	if e > 0 {
		n := len(b.String()) - len(d.String()) - 1
		var buffer bytes.Buffer
		for i := 0; i < n; i++ {
			buffer.WriteString("0")
		}
		dstr := d.String()
		dstr = strings.TrimRight(dstr, "0")
		buffer.WriteString(dstr)
		if ltZero {
			return fmt.Sprintf("-%v.%s", r, buffer.String())
		}
		return fmt.Sprintf("%v.%s", r, buffer.String())
	}
	if ltZero {
		return fmt.Sprintf("-%v", r)
	}
	return fmt.Sprintf("%v", r)
}

// 10进制转16进制
func DecimalToHex(num *big.Int) string {
	return "0x" + decimalToAny(num, 16)
}

// 10进制转任意进制
func decimalToAny(num *big.Int, n int64) string {
	newNumStr := ""
	var remainder *big.Int
	var remainderString string
	for num.Cmp(big.NewInt(0)) != 0 {
		remainder = big.NewInt(1).Mod(num, big.NewInt(n))
		r76 := remainder.Cmp(big.NewInt(76))
		r9 := remainder.Cmp(big.NewInt(9))
		if r76 == -1 && r9 > 0 {
			remainderString = tenToAny[remainder.Int64()]
		} else {
			remainderString = remainder.String()
		}
		newNumStr = remainderString + newNumStr
		num = big.NewInt(1).Div(num, big.NewInt(n))
	}
	if newNumStr == "" {
		newNumStr = "0"
	}
	return newNumStr
}

//map根据value找key
func findkey(in string) int64 {
	var result int64 = -1
	for k, v := range tenToAny {
		if in == v {
			result = k
		}
	}
	return result
}

// 16进制转10进制
func HexToString(num string) string {
	if num == "" || num == "0" {
		return big.NewInt(0).String()
	}
	r := strings.Index(num, "0x")
	if r == 0 {
		num = num[2:]
	}
	return anyToDecimal(num, 16).String()
}

// 任意进制转10进制
func anyToDecimal(num string, n int64) *big.Int {
	newNum := big.NewInt(0)
	nNum := len(strings.Split(num, "")) - 1
	for _, value := range strings.Split(num, "") {
		tmp := big.NewInt(findkey(value))
		if tmp.Int64() != -1 {
			ip := big.NewInt(1)
			newNum = ip.Mul(tmp, BigIntPow(n, int64(nNum))).Add(ip, newNum)
			nNum = nNum - 1
		} else {
			break
		}
	}
	return newNum
}

/**
 * n 如果为16
 * m 如果为9
 *结果为 16的9次方
 */
func BigIntPow(n, m int64) *big.Int {
	bigSum := big.NewInt(1)
	var i int64
	for i = 0; i < m; i++ {
		bigSum = big.NewInt(1).Mul(bigSum, big.NewInt(n))
	}
	return bigSum
}
