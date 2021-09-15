package featurea

import (
	"errors"
	"fmt"
)

func FeatureA() {
	fmt.Println("FeatureA")
}

var (
	ErrorFailedInsert = errors.New("failed insert to database")
)
