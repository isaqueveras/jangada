package newapp

const tmplDBSeeds = `package db

import "gorm.io/gorm"

type Seeds struct{}

func NewSeeds() *Seeds {
	return &Seeds{}
}

func (s *Seeds) Seed(db *gorm.DB) error {
	return nil
}`
