package zfscli

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

var header = regexp.MustCompile(`(\w+)(\s*)`)

func ScanTable(raw []byte, each func(i int, row []string) error) error {
	lines := bytes.Split(raw, []byte{'\n'})

	m := header.FindAllSubmatch(lines[0], -1)
	if m == nil {
		return fmt.Errorf("unable to parse header: %s", lines[0])
	}

	widths := make([]int, len(m))
	names := make([]string, len(m))
	for i, field := range m {
		w := len(field[2])
		if w != 0 {
			w += len(field[1])
		}
		widths[i] = w
		names[i] = string(field[1])
	}

	for i := 0; i < len(lines); i++ {
		row := make([]string, len(names))

		if i == 0 {
			copy(row, names)
		} else {
			lines[i] = bytes.TrimSpace(lines[i])
			if len(lines[i]) == 0 {
				row = nil
			} else {
				for f, w := range widths {
					if w != 0 {
						row[f] = string(lines[i][0:w])
						lines[i] = lines[i][w:]
					} else {
						row[f] = string(lines[i][0:])
					}
					row[f] = strings.TrimSpace(row[f])
				}
			}
		}

		if err := each(i, row); err != nil {
			return err
		}
	}

	return nil
}
