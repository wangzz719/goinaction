package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func isValidIpv4(s string) bool {
	// 判断最小长度
	if len(s) < 7 {
		return false
	}

	// 判断 token 数
	tokens := strings.Split(s, ".")
	if len(tokens) != 4 {
		return false
	}

	for _, t := range tokens {
		// 每个 token 长度在 1-3
		if len(t) > 3 || len(t) == 0 {
			return false
		}

		// 每个 token 数值大小在 0-255
		i, err := strconv.Atoi(t)
		if err != nil {
			return false
		}

		if i < 0 || i > 255 {
			return false
		}

		// 每个 token 不能以 '0' 为起始
		if len(t) > 1 && t[0] == '0' {
			return false
		}
	}

	return true
}

func getIpV4(in string, out string) {
	inFile, err := os.Open(in)
	if err != nil {
		panic(err)
	}

	defer inFile.Close()

	outFile, err := os.Create(out)
	if err != nil {
		panic(err)
	}

	defer outFile.Close()

	inputReader := bufio.NewReader(inFile)
	for {
		l, err := inputReader.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}

		ip := l[:len(l)-1]
		if isValidIpv4(ip) {
			outFile.WriteString(ip + "\n")
		}
	}
}

func main() {
	getIpV4("ipin.txt", "ipout.txt")
}
