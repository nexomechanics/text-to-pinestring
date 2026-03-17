// It's useful for cases where you want to write a long tooltip for an input or other applications like alert messages in pine script.
// Pine script has a weird syntax when it comes to multiline text and it can be hard to track. Hopefully this app makes it easier to convert regular text into pine script. If you do this a lot, it should save you a lot of time.
// You can also do this with AI / ChatGPT, but the input text should be perfect, otherwise the mistakes will be inherited or you will risk undesired porting. Often time, when doing with task with AI, you need a couple of tries before getting right, and that's a waste of tokens, context, and most importantly, time!
// For this to work beautifully, it must be paired with an UI that builds the expected json format. Please see more about that in the README file.

package main

import (
	"strings"
)

type block struct {
	Type    string     `json:"type"`
	Content string     `json:"content"`
	Headers []string   `json:"headers"`
	Rows    [][]string `json:"rows"`
}

func convert(input []block) (string, error) {
	var pine_string = ""
	for _, b := range input {

		if b.Type == "line" {
			pine_string += b.Content
		}

		if b.Type == "table" {
			colWidths := make([]int, len(b.Headers))
			for i, h := range b.Headers {
				colWidths[i] = len(h)
			}
			for _, row := range b.Rows {
				for i, cell := range row {
					if i < len(colWidths) && len(cell) > colWidths[i] {
						colWidths[i] = len(cell)
					}
				}
			}

			padRow := func(cells []string) string {
				padded := make([]string, len(cells))
				for i, cell := range cells {
					if i < len(colWidths)-1 {
						target := ((colWidths[i] / 8) + 1) * 8
						padded[i] = cell + strings.Repeat(" ", target-len(cell))
					} else {
						padded[i] = cell
					}
				}
				return strings.Join(padded, "\t")
			}

			allRows := make([]string, 0, 1+len(b.Rows))
			allRows = append(allRows, padRow(b.Headers))
			for _, row := range b.Rows {
				allRows = append(allRows, padRow(row))
			}
			pine_string += strings.Join(allRows, "\n")
		}
	}

	pine_string = strings.ReplaceAll(pine_string, `"`, `\"`) // catches all `"` in pine_string that might break the string output. Reason is that if someone types something like "hello" in the text block, the output would be ""hello"", which is not valid syntax in pine script.

	return `"` + pine_string + `"`, nil
}
