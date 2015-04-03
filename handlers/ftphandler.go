package handlers

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"ftp"
	"github.com/astaxie/beego"
	"io"
	"io/ioutil"
	"strings"
	"sync"
	"time"
)

type Ftphandler struct {
}

func initSession() (*ftp.ServerConn, error) {
	url := beego.AppConfig.String("ftpurl")
	username := beego.AppConfig.String("ftpusername")
	password := beego.AppConfig.String("ftppassword")
	//log.Printf("url:%s\n username:%s\n password:%s\n", url, username, password)
	c, err := ftp.DialTimeout(url, 5*time.Second)
	if err != nil {
		return c, err
	}
	err = c.Login(username, password)
	return c, err

}

func (h *Ftphandler) IsExist(src string) (isExist bool, err error) {
	c, err := initSession()
	if err != nil {
		return
	}
	defer c.Quit()
	r, err := c.Retr(src)
	if err != nil {
		return
	}
	r.Close()
	isExist = true
	return

}
func (h *Ftphandler) GetContent(src string) (reader io.ReadCloser, err error) {
	c, err := initSession()
	if err != nil {
		return
	}
	defer c.Quit()
	reader, err = c.Retr(src)
	return
}
func (h *Ftphandler) GetList(dir string) (lists []string, err error) {
	c, err := initSession()
	if err != nil {
		return
	}
	defer c.Quit()
	e, err := c.List(dir)
	if err != nil {
		return
	}
	for _, en := range e {
		lists = append(lists, en.Name)
	}
	return

}
func (h *Ftphandler) GetLength(src string) (length uint64, err error) {
	c, err := initSession()
	if err != nil {
		return
	}
	defer c.Quit()
	e, err := c.List(src)
	if err != nil || len(e) == 0 {
		return
	}
	length = e[0].Size
	return

}
func (h *Ftphandler) GetHash(src string) (hash string, err error) {
	c, err := initSession()
	if err != nil {
		return
	}
	defer c.Quit()
	srchash := src + ".hash"
	isexist, err := h.IsExist(srchash)
	if err != nil {
		return
	}
	if isexist {
		r, err := c.Retr(srchash)
		if err != nil {
			return
		}
		buf, err := ioutil.ReadAll(r)
		if err != nil {
			return
		}
		hash = string(buf)
		return
	} else {
		var mutex sync.Mutex
		mutex.Lock()
		defer mutex.Unlock()
		isexist, err := h.IsExist(srchash)
		if err != nil {
			return
		}
		if isexist {
			r, err := c.Retr(srchash)
			if err != nil {
				return
			}
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				return
			}
			hash = string(buf)
			return

		} else {
			isfileExist, err := h.IsExist(src)
			if err != nil || isfileExist == false {
				if err == nil {
					err = errors.New("file not exist")
				}
				return
			}
			sha1h := sha1.New()
			r, err := h.GetContent(src)
			if err != nil {
				return
			}
			io.Copy(sha1h, r)
			hash = fmt.Sprintf("%x", sha1h.Sum(nil))
			//fmt.Printf("sha1 %s", sha1s)
			hashr := strings.NewReader(sha1s)
			err = h.SetContent(srchash, hashr)
			if err != nil {
				return
			}
			return
		}

	}

}
func (h *Ftphandler) SetContent(src string, r io.Reader) (err error) {
	c, err := initSession()
	if err != nil {
		return
	}
	defer c.Quit()
	err = c.Stor(src, r)
	return

}