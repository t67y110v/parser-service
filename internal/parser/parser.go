package parser

import (
	"github.com/t67y110v/parser-service/internal/models"

	"github.com/sirupsen/logrus"
)

type Parser struct {
	Parse Parsing
}

type Parsing interface {
	Page(query string, page int) (*models.Response, error)
	All(query string) (*models.Response, error)
}

func NewCyberleninka(l *logrus.Logger) *Parser {

	return &Parser{
		Parse: NewCyberleninkaParser(l),
	}
}

func NewLvrach(l *logrus.Logger) *Parser {

	return &Parser{
		Parse: NewLvrachParser(l),
	}
}
