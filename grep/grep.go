package grep

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Search(pattern string, flags []string, files []string) []string {
	result := []string{}
	args_map := map[string]bool{}
	multi_file := len(files) > 1
	for _, arg := range flags {
		args_map[arg] = true
	}
	for _, file_name := range files {
		file, err := os.Open(file_name)
		if err != nil {
			panic("Could not open file")
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for n := 1; scanner.Scan(); n++ {
			line := scanner.Text()
			matched := line
			if _, ok := args_map["-i"]; ok {
				matched = strings.ToLower(line)
				pattern = strings.ToLower(pattern)
			}
			match := strings.Contains(matched, pattern)
			if _, ok := args_map["-x"]; ok {
				match = matched == pattern
			}

			if _, ok := args_map["-v"]; (!ok && match) || (ok && !match) {
				if _, ok := args_map["-n"]; ok {
					line = fmt.Sprintf("%d:%s", n, line)
				}
				if multi_file {
					line = fmt.Sprintf("%s:%s", file_name, line)
				}
				if _, ok := args_map["-l"]; ok {
					result = append(result, file_name)
					break
				}
				result = append(result, line)
			}
		}
	}
	return result
}
