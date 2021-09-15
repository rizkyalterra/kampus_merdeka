package featurea

import (
	"errors"
	"fmt"
)

func FeatureA() {
	fmt.Println("Hello World")
	fmt.Println("FeatureA")
}

var (
	ErrorFailedInsert = errors.New("failed insert to database")
)
