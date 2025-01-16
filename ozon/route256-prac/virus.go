package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var entries int
	fmt.Fscan(in, &entries)

	var jsonLines int
	var jsonString []byte
	var jsonMap map[string]any

	for range entries {
		res := 0

		fmt.Fscan(in, &jsonLines)
		in.ReadLine()

		for range jsonLines {
			jsonLine, _ := in.ReadBytes('\n')
			jsonString = append(jsonString, jsonLine[:len(jsonLine)-1]...)
		}

		err := json.Unmarshal(jsonString, &jsonMap)
		if err != nil {
			panic(err)
		}

		traverse(false, jsonMap, &res)

		out.WriteString(fmt.Sprintf("%v\n", strconv.Itoa(res)))

		// cleanup
		jsonString = []byte{}
		jsonLines = 0
		jsonMap = map[string]any{}
	}
}

func traverse(isInfected bool, jsonMap map[string]any, res *int) {
	if v, ok := jsonMap["files"]; ok {
		// marking infected files
		vArr := v.([]any)
		if isInfected {
			*res += len(vArr)
		} else {
			// iterating over files in order to find infected ones
			for _, file := range vArr {
				strFile := file.(string)
				// found infected -> mark all files as infected and quit
				if strings.HasSuffix(strFile, ".hack") {
					isInfected = true
					*res += len(vArr)
					break
				}
			}
		}
	}

	if v, ok := jsonMap["folders"]; ok {
		vMap := v.([]any)
		for _, folder := range vMap {
			folderMap := folder.(map[string]any)
			traverse(isInfected, folderMap, res)
		}
	}
}
