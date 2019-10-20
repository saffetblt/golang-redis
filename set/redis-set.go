package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fatih/structs"
	"github.com/go-redis/redis"
)

type Rule struct {
	Kaynak      string `json:"Kaynak"`
	Activite    string `json:"Activite"`
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
	client := newClient()
	jsonFile, err := os.Open("../protocol.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var rules []Rule
	json.Unmarshal(byteValue, &rules)
	var rulesM map[string]interface{}

	for i := 0; i < len(rules); i++ {
		rulesM = structs.Map(rules[i])
		err = client.HMSet(rules[i].SignatureID, rulesM).Err()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func newClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return client
}

func hgetall(client *redis.Client, sdi string) {
	m, err := client.HGetAll(sdi).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(m)
}
