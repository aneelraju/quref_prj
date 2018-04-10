package models

import (
	"encoding/json"
	"os"
)

// comment struct
type Comment struct {
	CommentId string `json:"commentid"`
	Comment   string `json:"comment"`
}

// Topic struct contains list of comments
type Topic struct {
	TopicName string `json:"topicname"`
	Comments  []Comment
}

// Feature struct contains list of topics
type Feature struct {
	FeatureName string `json:"FeatureName"`
	Topics      []Topic
}

// Section struct contains list of Features
type Section struct {
	SectionName string `json:"sectionname"`
	Features    []Feature
}
