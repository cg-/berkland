package cl

import (
	"fmt"
	"log"
	"net/http"
)

/*
CraigsListParser is an object that scrapes CraigsList.

We'll use it to allow us to search for lost and found posts for
animals.
*/
type CraigsListParser struct {
	city     string
	category string
	terms    []string
}

func NewCraigsListParser() *CraigsListParser {
	log.Printf("Created new CraigsListParser\n")
	return &CraigsListParser{}
}

func (c *CraigsListParser) Get() ([]string, error) {
	endPath := ""
	url := fmt.Sprintf("http://%s.craigslist.org/%s", c.city, endPath)
	resp, err := http.Get(url)
	if err != nil {
		return []string{}, err
	}

	log.Println(resp)
	return []string{}, nil
}

// SetCity changes the city we're looking at
func (c *CraigsListParser) SetCity(name string) {
	c.city = name
}

// SetCategory changes the category we're looking at
func (c *CraigsListParser) SetCategory(name string) {
	c.category = name
}

// SetTerms changes the search terms that we're interested in
func (c *CraigsListParser) SetTerms(name []string) {
	c.terms = name
}

// AddTerms adds some strings to search for
func (c *CraigsListParser) AddTerms(name ...string) {
	for i := range name {
		c.terms = append(c.terms, name[i])
	}
}

// DeleteTerms gets rid of some search terms
func (c *CraigsListParser) DeleteTerms(name ...string) {
	for i := range name {
		for j := range c.terms {
			if c.terms[j] == name[i] {
				c.terms = append(c.terms[:j], c.terms[:j+1]...)
			}
		}
	}
}
