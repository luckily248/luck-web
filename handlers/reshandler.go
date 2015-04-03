package handlers

import (
	"io"
)

type Reshandler interface {
	IsExist(src string) (isExist bool, err error)
	GetContent(src string) (reader io.ReadCloser, err error)
	GetList(dir string) (lists []string, err error)
	GetLength(src string) (length uint64, err error)
	GetHash(src string) (hash string, err error)
	SetContent(src string, r io.Reader) (err error)
}
