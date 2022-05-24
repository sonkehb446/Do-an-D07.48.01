package view

import (
	"html/template"
	"strings"
	"time"
)

func Funcs() template.FuncMap {
	return template.FuncMap{
		"FormatTime":  formatTime,
		"StringLower": formatLower,
		"FormatRole":  formatRole,
	}
}
func formatTime(t time.Time) string {
	// return t.Format("2006/01/02 15:04:05")
	return t.Format("2006/01/02")
}

func formatLower(s string) string {
	low := strings.ToLower(s)
	return low
}

func formatRole(r *bool) string {
	if *r == true {
		return "Admin"
	}
	return "User"
}
