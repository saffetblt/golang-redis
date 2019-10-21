package main

import (
	"fmt"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/mitchellh/mapstructure"
)

type Rule struct {
	Source      string `json:"Source"`
	Active      string `json:"Active"`
	SrcIP       string `json:"SrcIP"`
	SrcPort     string `json:"SrcPort"`
	DstIP       string `json:"DstIP"`
	DstPort     string `json:"DstPort"`
	Protocol    string `json:"Protocol"`
	Function    string `json:"Function"`
	Signature   string `json:"Signature"`
	Category    string `json:"Category"`
	Severity    string `json:"Severity"`
	Mitigation  string `json:"Mitigation"`
	SignatureID string `json:"SignatureID"`
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	SrcIP := "any"
	SrcPort := "any"
	DstIP := "any"
	DstPort := "any"
	Protocol := "MODBUS"
	Function := "5"
	var rule Rule

	for i := 105010; i <= 105018; i++ {
		m, err := client.HGetAll(strconv.Itoa(i)).Result()
		if err != nil {
			fmt.Println(err)
		}
		err = mapstructure.Decode(m, &rule)
		if err != nil {
			fmt.Println(err)
		}
		if rule.SrcIP == SrcIP && rule.SrcPort == SrcPort &&
			rule.DstIP == DstIP && rule.DstPort == DstPort &&
			rule.Protocol == Protocol && rule.Function == Function {
			fmt.Println(rule)
		}
	}
}
