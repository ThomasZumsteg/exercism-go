package markdown

// implementation to refactor

import "fmt"

// Render translates markdown to HTML
func Render(markdown string) string {
	strong := false
	em := false
	header := 0
	list := 0
	html := ""
	for pos := 0; pos < len(markdown); pos++ {
		switch char := markdown[pos]; char {
		case '#':
			for ; char == '#'; char = markdown[pos] {
				header++
				pos++
			}
			html += fmt.Sprintf("<h%d>", header)
		case '*':
			if list == 0 {
				html += "<ul>"
			}
			html += "<li>"
			list++
			pos++
		case '\n':
			if list > 0 {
				html += "</li>"
			}
			if header > 0 {
				html += fmt.Sprintf("</h%d>", header)
				header = 0
			}
		case '_':
			if pos+1 < len(markdown) && markdown[pos+1] == '_' {
				strong = !strong
				if strong {
					html += "<strong>"
				} else {
					html += "</strong>"
				}
				pos++
			} else {
				em = !em
				if em {
					html += "<em>"
				} else {
					html += "</em>"
				}
			}
		default:
			html += string(char)
		}
	}
	if header > 0 {
		return html + fmt.Sprintf("</h%d>", header)
	}
	if list > 0 {
		return html + "</li></ul>"
	}
	return "<p>" + html + "</p>"

}
