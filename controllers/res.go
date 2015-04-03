package controllers

import (
	"fmt"
	//	"fmt"
	"github.com/astaxie/beego"
	"log"
	. "luck-web/handlers"
)

type ResController struct {
	beego.Controller
	handler Reshandler
}

func (c *ResController) Post() {

	action := c.GetString("action")
	c.TplNames = "index.tpl"
	c.handler = &Ftphandler{}
	//log.Printf("action %s", action)
	switch action {
	case "hash":
		getHash(c)

	case "content":
		//log.Println("content")
		getContent(c)

	case "404":
		is404(c)

	case "list":
		list(c)

	case "length":
		getLength(c)

	}
	c.Ctx.WriteString("forbiden")
	return

}

//列文件信息
func getHash(c *ResController) {
	src := c.GetString("src")
	if "" == src {
		c.Ctx.Output.SetStatus(404)
		return
	}
	isexist, err := c.handler.IsExist(src)
	if err != nil || !isexist {
		log.Fatal(err)
		c.Ctx.Output.SetStatus(404)
		return
	}

	hash, err := c.handler.GetHash(src)
	if err != nil {
		log.Fatal(err)
		return
	}
	c.Ctx.WriteString(hash)
	return
}

//获取文件
func getContent(c *ResController) {
	src := c.GetString("src")
	if "" == src {
		c.Ctx.Output.SetStatus(404)
		return
	}
	r, err := c.handler.GetContent(src)
	defer r.Close()
	if err != nil {
		log.Fatal(err)
		c.Ctx.Output.SetStatus(404)
		return
	}
	log.Printf("isexist %s\n", src)
	//content, _ := ioutil.ReadAll(r)
	//log.Printf("content %s\n", content)
	buf := make([]byte, 1024)
	c.Ctx.Output.ContentType("application/octet-stream")
	//n, err := r.Read(buf)
	//log.Printf("%-20s %-2v %v\n", buf[:n], n, err)
	for offset, _ := r.Read(buf); offset > 0; offset, _ = r.Read(buf) {

		//log.Printf("buf %s\n", buf)
		//log.Printf("off %d\n", offset)
		_, err := c.Ctx.ResponseWriter.Write(buf)
		if err != nil {
			log.Fatal(err)
		}

	}
	r.Close()
	//log.Printf("%s", buf)
	return

}

//是否存在
func is404(c *ResController) {
	src := c.GetString("src")
	if "" == src {
		c.Ctx.Output.SetStatus(404)
		return
	}
	isexist, err := c.handler.IsExist(src)
	if err != nil {
		log.Fatal(err)
		c.Ctx.Output.SetStatus(404)
		return
	}
	c.Ctx.WriteString(fmt.Sprintf("%v", !isexist))
	return
}

//列表
func list(c *ResController) {
	//log.Println("list")
	dir := c.GetString("dir")
	if "" == dir {
		c.Ctx.Output.SetStatus(404)
		return
	}
	//	isexist, err := c.handler.IsExist(dir)
	//	if err != nil {
	//		log.Fatal(err)
	//		c.Ctx.Output.SetStatus(404)
	//		return
	//	}
	//	if !isexist {
	//		c.Ctx.Output.SetStatus(404)
	//		return
	//	}
	//	log.Println("!404")
	lists, err := c.handler.GetList(dir)
	if err != nil {
		log.Fatal(err)
		c.Ctx.WriteString(err.Error())
		return
	}
	//log.Printf("%d\n", len(lists))
	if len(lists) == 0 {
		c.Ctx.Output.SetStatus(404)
		return
	}
	for _, e := range lists {
		c.Ctx.WriteString(fmt.Sprintf("%s\n", e))
		//log.Printf("%s\n", e)
	}
	return
}

//获得长度
func getLength(c *ResController) {
	src := c.GetString("src")
	if "" == src {
		c.Ctx.Output.SetStatus(404)
		return
	}
	isexist, err := c.handler.IsExist(src)
	if err != nil {
		log.Fatal(err)
		c.Ctx.Output.SetStatus(404)
		return
	}
	if !isexist {
		c.Ctx.Output.SetStatus(404)
		return
	}
	l, err := c.handler.GetLength(src)
	if err != nil {
		log.Fatal(err)
		c.Ctx.Output.SetStatus(404)
		return
	}
	c.Ctx.WriteString(fmt.Sprintf("%d", l))
	return
}
