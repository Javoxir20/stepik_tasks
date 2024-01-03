package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Student struct {
	LastName   string   `json:"LastName"`
	FirstName  string   `json:"FirstName"`
	MiddleName string   `json:"MiddleName"`
	Birthday   string   `json:"Birthday"`
	Address    string   `json:"Address"`
	Phone      string   `json:"Phone"`
	Rating     []int    `json:"Rating"`
}
