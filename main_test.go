package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strings"
	"testing"
)

func TestGetAge(t *testing.T) {
	_, err := http.Get("http://localhost:8080/age")
	if err != nil {
		fmt.Println(err)
		t.Error("测试失败")
	}
}

func TestGetCar(t *testing.T) {
	_, err := http.Get("http://localhost:8080/car")
	if err != nil {
		fmt.Println(err)
		t.Error("测试失败")
	}
}
func TestGetBuffer(t *testing.T) {
	_, err := http.Get("http://localhost:8080/buffer")
	if err != nil {
		fmt.Println(err)
		t.Error("测试失败")
	}
}
func TestGetRate(t *testing.T) {
	_, err := http.Get("http://localhost:8080/rate")
	if err != nil {
		fmt.Println(err)
		t.Error("测试失败")
	}
}

func TestAddAge(t *testing.T) {
	pushUlr := "http://localhost:8080/age"
	type AgeModel struct {
		Age int `json:"age"`
	}

	var ageModel AgeModel
	ageModel.Age = int(math.Round(60))
	data, _ := json.Marshal(ageModel)
	_, err := http.Post(pushUlr, "application/json", strings.NewReader(string(data)))
	if err != nil {
		t.Error("测试失败")
	}
	r, err1 := http.Get("http://localhost:8080/age")
	if err1 != nil {
		t.Error("测试失败")
	}
	b, errb := ioutil.ReadAll(r.Body)
	if errb != nil {
		t.Error("测试失败")
	}
	var ageModelRet AgeModel
	err = json.Unmarshal(b, &ageModelRet)
	if err != nil {
		t.Error("测试失败")
	}
	if ageModelRet.Age != ageModel.Age {
		t.Error("测试失败")
	}
}
