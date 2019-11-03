package main

import (
	"GoPlayWithCompiler/craft"
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	verbose bool
)

func init() {
	flag.BoolVar(&verbose, "v", false, "-v 1 to print detail")
}

func main() {
	flag.Parse()
	fmt.Println("Simple script language!")
	fmt.Println("input exit(); to quit")
	f := bufio.NewReader(os.Stdin) //读取输入的内容
	promt := ">>"
	scriptText := ""
	calculator := craft.SimpleParser{}
	for {
		fmt.Print(promt)
		input, _ := f.ReadString('\n') //定义一行输入的内容分隔符。
		if len(input) == 1 {
			continue //如果用户输入的是一个空行就让用户继续输入。
		}
		input = strings.TrimSpace(input)
		if input == "exit();" {
			fmt.Println("good bye!")
			break
		}
		scriptText += input + "\n"
		if strings.HasSuffix(scriptText, ";\n") {
			fmt.Println("your input is :" + scriptText)
			root := calculator.Parse(scriptText)
			if verbose {
				craft.DumpAST(*root, "")
			}
			script := craft.NewSimpleScript(verbose)
			script.Evaluate(*root, "")
			scriptText = ""
		}
	}
}
