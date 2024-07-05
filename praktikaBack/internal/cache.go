package internal;

import (

)

type ICache interface {
	GetVacancies() (*vacancies, error)
}

type cache struct {
	vacanciesArray *vacancies
	db IDB
}

func NewCache(db IDB) ICache {
	return &cache{
		vacanciesArray: &vacancies{},
		db: db,
	}
}

func (c *cache) addVacancies() error {
	vacancies, err := c.db.GetVacancies()
	if err != nil {
		return err;
	}

	c.vacanciesArray = vacancies;
	return nil;
}

func (c *cache) GetVacancies() (*vacancies, error) {
	err := c.addVacancies()
	if err != nil {
		return nil, err
	}
	return c.vacanciesArray, nil;
}