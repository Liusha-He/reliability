package main

import "encoding/json"

func GetRandomFact() (RandomFact, error) {
	var fact RandomFact

	resp, err := CLIENT.Get("https://uselessfacts.jsph.pl/random.json?language=en")

	if err != nil {
		return RandomFact{}, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&fact)
	if err != nil {
		return RandomFact{}, err
	}

	return fact, nil
}
