package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Kg float64 //公斤
type Lb float64 //磅
type Ft float64 //英尺
type M float64  //米

func (kg Kg) String() string { return fmt.Sprintf("%g kg", kg) }
func (lb Lb) String() string { return fmt.Sprintf("%g lb", lb) }
func (ft Ft) String() string { return fmt.Sprintf("%g ft", ft) }
func (m M) String() string   { return fmt.Sprintf("%g m", m) }

func KgToLb(kg Kg) Lb { return Lb(kg / 2.2046226) }
func LbToKg(lb Lb) Kg { return Kg(lb * 2.2046226) }
func FtToM(ft Ft) M   { return M(ft * 0.3048) }
func MToFt(m M) Ft    { return Ft(m / 0.3048) }

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
	Begin:
		fmt.Printf("1.%s\t2.%s\t3.%s\n", "长度", "重量", "退出")
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			switch input.Text() {
			case "1":
				fmt.Printf("1.%s\t2.%s\t3.%s\n", "公斤转磅", "磅转公斤", "返回")
				in1 := bufio.NewScanner(os.Stdin)
				if in1.Scan() {
					switch in1.Text() {
					case "1":
						fmt.Println("请输入需要转换的值: ")
						invalue1 := bufio.NewScanner(os.Stdin)
						if invalue1.Scan() {
							t, err := strconv.ParseFloat(invalue1.Text(), 64)
							if err != nil {
								fmt.Fprintf(os.Stderr, "transform: %v\n", err)
							}
							kg := Kg(t)
							fmt.Printf("%s= %s", kg, KgToLb(kg))
						}
					case "2":
						fmt.Println("请输入需要转换的值: ")
						invalue2 := bufio.NewScanner(os.Stdin)
						if invalue2.Scan() {
							t, err := strconv.ParseFloat(invalue2.Text(), 64)
							if err != nil {
								fmt.Fprintf(os.Stderr, "transform: %v\n", err)
							}
							lb := Lb(t)
							fmt.Printf("%s = %s", lb, LbToKg(lb))
						}
					case "3":
						fallthrough
					default:
						goto Begin
					}
				} else {
					goto Begin
				}
			case "2":
				fmt.Printf("1.%s\t2.%s\t3.%s\n", "米转英尺", "英尺转米", "返回")
				in2 := bufio.NewScanner(os.Stdin)
				if in2.Scan() {
					switch in2.Text() {
					case "1":
						fmt.Println("请输入需要转换的值: ")
						invalue1 := bufio.NewScanner(os.Stdin)
						if invalue1.Scan() {
							t, err := strconv.ParseFloat(invalue1.Text(), 64)
							if err != nil {
								fmt.Fprintf(os.Stderr, "transform: %v\n", err)
							}
							m := M(t)
							fmt.Printf("%s = %s", m, MToFt(m))
						}
					case "2":
						fmt.Println("请输入需要转换的值: ")
						invalue2 := bufio.NewScanner(os.Stdin)
						if invalue2.Scan() {
							t, err := strconv.ParseFloat(invalue2.Text(), 64)
							if err != nil {
								fmt.Fprintf(os.Stderr, "transform: %v\n", err)
							}
							ft := Ft(t)
							fmt.Printf("%s = %s", ft, FtToM(ft))
						}
					case "3":
						fallthrough
					default:
						goto Begin
					}
				} else {
					goto Begin
				}
			case "3":
				return
			default:
				goto Begin
			}
		}
	}

}
