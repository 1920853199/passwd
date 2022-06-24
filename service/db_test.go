package service

import (
	"fmt"
	"os"
	"testing"
)

type Account struct {
	Ip       string `json:"ip"`
	User     string `json:"user"`
	Password string `json:"password"`
	Remark   string `json:"remark"`
}

func TestDbGetAndSet(t *testing.T) {
	os.RemoveAll("~/.passwd/default")
	os.MkdirAll("~/.passwd/default", 0777)

	db, err := NewDB("~/.passwd/default")
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}

	db.Set([]byte("a"), []byte("aaaa"))
	db.Set([]byte("b"), []byte("bbbb"))
	db.Set([]byte("c"), []byte("cccc"))

	val, err := db.Get([]byte("a"))
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	fmt.Printf("a value is : %s\n", string(val))
	val, err = db.Get([]byte("b"))
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	fmt.Printf("b value is : %s\n", string(val))
	val, err = db.Get([]byte("c"))
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	fmt.Printf("c value is : %s\n", string(val))
}

func TestDbCodecGetAndCodecSet(t *testing.T) {
	os.RemoveAll("~/.passwd/default")
	os.MkdirAll("~/.passwd/default", 0777)

	db, err := NewDB("~/.passwd/default")
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	err = db.SetWithCodec([]byte("a"), Account{
		Ip:       "192.168.31.123",
		User:     "lee",
		Password: "2014abc",
		Remark:   "阿里云账号",
	})
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}

	var v Account
	err = db.GetWithCodec([]byte("a"), &v)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	fmt.Printf("a value is : %v\n", v)
}

func TestDbClear(t *testing.T) {
	os.RemoveAll("~/.passwd/default")
	os.MkdirAll("~/.passwd/default", 0777)

	db, err := NewDB("~/.passwd/default")
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}

	db.Set([]byte("a"), []byte("aaaa"))
	db.Set([]byte("b"), []byte("bbbb"))
	db.Set([]byte("c"), []byte("cccc"))

	val, err := db.Get([]byte("a"))
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	fmt.Printf("a value is : %s\n", string(val))
	val, err = db.Get([]byte("b"))
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	fmt.Printf("b value is : %s\n", string(val))
	val, err = db.Get([]byte("c"))
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	fmt.Printf("c value is : %s\n", string(val))

	db.SetWithCodec([]byte("a"), Account{
		Ip:       "192.168.31.123",
		User:     "lee",
		Password: "2014abc",
		Remark:   "阿里云账号",
	})
	var v Account
	err = db.GetWithCodec([]byte("a"), &v)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	fmt.Printf("a value is : %v\n", v)

	err = db.Clear()

	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}

	val, err = db.Get([]byte("a"))
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	fmt.Printf("a value is : %s\n", string(val))
	val, err = db.Get([]byte("b"))
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	fmt.Printf("b value is : %s\n", string(val))
	val, err = db.Get([]byte("c"))
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	fmt.Printf("c value is : %s\n", string(val))

	var vv Account
	err = db.GetWithCodec([]byte("a"), vv)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	fmt.Printf("a value is : %v\n", vv)
}
