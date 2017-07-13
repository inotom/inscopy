package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// PackageJSON package.json のデコード用構造体
type PackageJSON struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Author  string `json:"author"`
	License string `json:"license"`
}

func main() {
	var isSmallFmt bool

	// オプション引数のパース
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage of %s:
  %s [OPTIONS]

Options
  -h	Show this message
`, os.Args[0], os.Args[0])
		flag.PrintDefaults()
	}
	flag.BoolVar(&isSmallFmt, "s", false, "insert small formatted copyright")
	flag.Parse()

	// package.json の読み込み
	bytes, err := ioutil.ReadFile("package.json")
	if err != nil {
		log.Fatal(err)
	}
	var pj PackageJSON
	if err := json.Unmarshal(bytes, &pj); err != nil {
		log.Fatal(err)
	}

	// copyright 文字列の作成
	var copyStr string
	if isSmallFmt {
		copyStr = fmt.Sprintf("/*! %s v%s */ ", pj.Name, pj.Version)
	} else {
		copyStr = fmt.Sprintf("/*! %s v%s %s | %s */ ", pj.Name,
			pj.Version, pj.Author, pj.License)
	}

	// 標準入力文字列の先頭行に copyright 文字列を追加
	fmt.Print(copyStr)
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		if err := stdin.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		fmt.Println(stdin.Text())
	}
}
