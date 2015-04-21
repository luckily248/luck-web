package test

import (
	"luck-web/models"
	"testing"
	"time"
)

func TestMgoAdd(b *testing.T) {
	bgtime := time.Now()
	b.Logf("%s:new user", bgtime.String())
	user := models.User{}
	enbutime := time.Now()
	b.Logf("%s:end new user use  %v ", enbutime.String(), enbutime.Sub(bgtime))
	err := models.AddUser(user)
	if err != nil {
		b.Errorf("err :%v", err)
	}
	enaddtime := time.Now()
	b.Logf("%s:end add use %v", enaddtime.String(), enaddtime.Sub(enbutime))
}

func Benchmark_Add(b *testing.B) {
	b.StopTimer()
	user := models.User{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		err := models.AddUser(user)
		if err != nil {
			b.Errorf("err :%v", err)
		}
	}

}
