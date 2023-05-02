package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/Ferdzzzzzzzz/sys_utils/x/twx"
)

func main() {
	tw := twx.NewWriter(os.Stdout)

	dir := "."
	if len(os.Args) > 1 {
		dir += "/" + os.Args[1]
	}

	panicIfErr(readDir(tw, dir, 0))

	tw.Flush()

}

func readDir(tw *tabwriter.Writer, path string, indent int) error {
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	var indentStr string
	for i := 0; i < indent; i++ {
		indentStr = fmt.Sprintf("%s\t", indentStr)
	}

	for _, e := range entries {
		fmt.Fprintf(tw, "%s%s\n", indentStr, e.Name())
		if e.IsDir() {
			if err := readDir(tw, path+"/"+e.Name(), indent+1); err != nil {
				return err
			}
		}
	}

	return nil
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
