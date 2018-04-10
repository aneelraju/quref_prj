/*
package controllers

import (
	"fmt"
	"github.com/aneelraju/quref_prj/models"
)

func (c *Comment) PrintComment() {
	fmt.Printf("%s: %s\n", CommentId, Comment)
}

func (t *Topic) PrintTopic() {
	if t.TopicName == nil {
		fmt.Fatalln("TopicName not found")
	} else {
		fmt.Printf("Topic: %s:\n", t.TopicName)
		fmt.Println("Comment list:")
		for _, c := range t.Comments {
			c.PrintComment()
		}
	}
}

func (t *Topic) AddCommentToTopic(c Comment) {
	t.Comment = append(t.Comment, c)
}

func (t *Topic) DeleteCommentFromTopic(id int) {
	found := 0
	index := 0
	for i, c := range t.Comments {
		if t.Comments[i].CommentId == id {
			found := 1
			index := i
			break
		}
	}
	if found == 1 {
		// delete item from slice, @FIXME move to separate func
		copy(t.Comments[index:], t.Comments[index+1:])
		t.Comments[len(a)-1] = ""
		t.Comments = t.Comments[:len(t.Comments)-1]
	}
}

func (f *Feature) PrintFeature() {
	fmt.Printf("Feature: %s:\n", f.FeatureName)
	fmt.Println("Topic list:")
	for _, t := range f.Topics {
		fmt.Printf("	%s\n", t.TopicName)
	}
}

func (s *Section) PrintSection() {
	fmt.Printf("Section: %s\n", s.SectionName)
	fmt.Println("Feature list:")
	for _, f := range s.Features {
		fmt.Printf("	%s\n", f.FeatureName)
	}
}
*/
