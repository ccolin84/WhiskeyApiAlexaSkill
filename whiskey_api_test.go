package main

import (
	"testing"
)

func TestGetRandomWhiskeyFactIsntEmpty(t *testing.T) {
	fact := GetRandomWhiskeyFact()

	if len(fact) == 0 {
		t.Error("Random fact is empty")
	}
}

func TestGetRandomWhiskeyFactReturnsDifferentFacts(t *testing.T) {
	factNumber := 4
	var facts []string
	
	for i := 0; i < factNumber; i++ {
		facts = append(facts, GetRandomWhiskeyFact())
	}

	differentFacts := false
	for i := 0; i < factNumber; i++ {
		if differentFacts {
			break
		}
		
		for j := i + 1; j < factNumber; j++ {
			if facts[i] != facts[j] {
				differentFacts = true
				break
			}
		}
	}

	if !differentFacts {
		t.Error("Not getting different facts from the API")
	}
}
