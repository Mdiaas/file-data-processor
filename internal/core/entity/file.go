package entity

import "io"

type File struct {
	Header string
	File   io.Reader
	Name   string
}
