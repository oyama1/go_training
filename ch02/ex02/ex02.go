package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"github.com/oyama1/go_training/ch02/ex02/conv"
)

func main() {
	if (len(os.Args) <= 1) {
		printUnit()
	} else {
		for _, arg := range os.Args[1:] {
			// 引数で値が来た場合は、全ての単位を出力する
			fmt.Printf("All unit for value = %s\n", arg)
			printTemp(arg)
			printLength(arg)
			printWeight(arg)
		}
	}
}

func printUnit() {
	input := bufio.NewScanner(os.Stdin)

	// TODO この方式だと、入力回数が増え、制限も多くなってしまうのでもっと良い方法を模索したい 
    fmt.Printf("Input unit to convert in (%s, %s, %s)  : ", conv.Temp, conv.Length, conv.Weight)

	input.Scan()
	unit := input.Text()

	if (unit == conv.Temp) {
		printTemp(inputValue())
	} else if (unit == conv.Length) {
		printLength(inputValue())
	} else if (unit == conv.Weight) {
		printWeight(inputValue())
	} else {
		fmt.Printf("Unit should be inputted in (%s, %s, %s) \n", conv.Temp, conv.Length, conv.Weight)
		os.Exit(1) // 処理をstatus = 1 で終了
	}
}

func inputValue() string {
	inputStr := ""
	input := bufio.NewScanner(os.Stdin)

    fmt.Print("Input value for convert  : ")

	input.Scan()
	inputStr = input.Text()

	return inputStr
}

func printTemp(inputValue string) {
	f,c := convTemp(inputValue)
	fmt.Printf("%s = %s, %s = %s\n",
		f, conv.FToC(f), c, conv.CToF(c))
}

func printLength(inputValue string) {
	meter,feet := convLength(inputValue)
	fmt.Printf("%s = %s, %s = %s\n",
		meter, conv.MeterToFeet(meter), feet, conv.FeetToMeter(feet))
}

func printWeight(inputValue string) {
	kg,lb := convWeight(inputValue)
	fmt.Printf("%s = %s, %s = %s\n",
		kg, conv.KgToLb(kg), lb, conv.LbToKg(lb))
}

func convTemp(input string) (conv.Fahrenheit, conv.Celsius) {
	temp, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "convTemp: %v\n", err)
		os.Exit(1) // 処理をstatus = 1 で終了
	}
	fahrenheit := conv.Fahrenheit(temp)
	celsius := conv.Celsius(temp)
	return fahrenheit,celsius
}

func convLength(input string) (conv.Meter, conv.Feet) {
	length, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "convLength: %v\n", err)
		os.Exit(1) // 処理をstatus = 1 で終了
	}
	meter := conv.Meter(length)
	feet := conv.Feet(length)
	return meter,feet
}

func convWeight(input string) (conv.Kg, conv.Lb) {
	weight, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "convWeight: %v\n", err)
		os.Exit(1) // 処理をstatus = 1 で終了
	}
	kg := conv.Kg(weight)
	lb := conv.Lb(weight)
	return kg,lb
}