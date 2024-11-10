// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Article struct {
	ID           string       `json:"id"`
	Title        string       `json:"title"`
	Body         string       `json:"body"`
	Author       string       `json:"author"`
	Views        int          `json:"views"`
	Likes        int          `json:"likes"`
	Shares       int          `json:"shares"`
	ArticleStyle ArticleStyle `json:"articleStyle"`
	Created      string       `json:"created"`
	Updated      string       `json:"updated"`
}

type Mutation struct {
}

type NewArticle struct {
	Title       string       `json:"title"`
	Body        string       `json:"body"`
	ArticleType ArticleStyle `json:"articleType"`
	Author      string       `json:"author"`
}

type Query struct {
}

type ArticleStyle string

const (
	ArticleStyleGreenViper      ArticleStyle = "GreenViper"
	ArticleStyleMediumVioletRed ArticleStyle = "MediumVioletRed"
	ArticleStyleDodgerBlue      ArticleStyle = "DodgerBlue"
	ArticleStyleDimGray         ArticleStyle = "DimGray"
)

var AllArticleStyle = []ArticleStyle{
	ArticleStyleGreenViper,
	ArticleStyleMediumVioletRed,
	ArticleStyleDodgerBlue,
	ArticleStyleDimGray,
}

func (e ArticleStyle) IsValid() bool {
	switch e {
	case ArticleStyleGreenViper, ArticleStyleMediumVioletRed, ArticleStyleDodgerBlue, ArticleStyleDimGray:
		return true
	}
	return false
}

func (e ArticleStyle) String() string {
	return string(e)
}

func (e *ArticleStyle) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ArticleStyle(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ArticleStyle", str)
	}
	return nil
}

func (e ArticleStyle) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
