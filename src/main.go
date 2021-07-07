package main

import (
	"bufio"
	"crypto/md5"
	"flag"
	"fmt"
	"os"
)

func main() {

	op := flag.Int("O", -1, "this is operating")

	flag.Parse()

	if *op == -1 {

		fmt.Println("=================================================")
		fmt.Println("this is Tools Description")
		fmt.Println("1.Md5 : input 1 ")
		fmt.Println("2.Url Encode : 2 ")
		fmt.Println("3.Url DeEncode : 3 ")
		fmt.Println("=================================================")
		fmt.Println("Please Chose Your Tools ")
		var inputOp int
		fmt.Scan(&inputOp)
		op = &inputOp
	}
	startUp(*op)
}

func startUp(op int) {
	switch op {
	case 1:
		{
			fmt.Print("请输入字符串  回车确定 ")
			var input = getInput()
			b := md5.Sum(([]byte(input)))
			fmt.Printf("%x", b)
		}
	case 2:
		{
		}
	case 3:
		{

		}
	}
}

func getInput() string {
	var intputBuffer = bufio.NewReader(os.Stdin)
	input, err := intputBuffer.ReadString('\n')
	if err == nil {
		fmt.Print("Your Input is ", input)
		return input
	}
	return ""
}
