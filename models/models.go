package models

// import (
//	"encoding/json"
// )

// Sections struct contains list of sections
type Sections struct {
	Sections []Section `json:"sections"`
}

// Section struct contains list of Features
type Section struct {
	SectionName string    `json:"sectionname"`
	Features    []Feature `json:"features"`
}

// Feature struct contains list of topics
type Feature struct {
	FeatureName string  `json:"featurename"`
	Topics      []Topic `json:"topics"`
}

// Topic struct contains list of comments
type Topic struct {
	TopicName string    `json:"topicname"`
	Comments  []Comment `json:"comments"`
}

// comment struct
type Comment struct {
	CommentId string `json:"commentid"`
	Comment   string `json:"comment"`
}
