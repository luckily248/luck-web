package test

import (
	"bytes"
	"io/ioutil"
	"luck-web/handlers"
	"testing"
)

const (
	testData = "test it"
	testDir  = "mydir"
)

func TestFtphandler(t *testing.T) {
	h := handlers.Ftphandler{}

	//测试存储
	data := bytes.NewBufferString(testData)
	err := h.SetContent("/test", data)
	if err != nil {
		t.Error(err)
	}

	//是否存在
	isExist, err := h.IsExist("/test")
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("isexist,%s\n", isExist)
	}

	//列目录
	lists, err := h.GetList("/")
	if err != nil {
		t.Error(err)
	} else {
		for num, name := range lists {
			t.Logf("name,%s,%d\n", name, num)
		}
	}

	//获得长度
	lenght, err := h.GetLength("/test")
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("len,%d\n", lenght)
	}

	//获得哈希
	hashs, err := h.GetHash("/test")
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("hash,%s\n", hashs)
	}

	//获得文件内容
	r, err := h.GetContent("/test")
	if err != nil {
		t.Error(err)
	} else {
		buf, err := ioutil.ReadAll(r)
		if err != nil {
			t.Error(err)
		}
		if string(buf) != testData {
			t.Errorf("'%s'", buf)
		} else {
			t.Logf("%s", buf)
		}
		r.Close()
	}
}
