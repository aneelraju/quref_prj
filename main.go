package main

import (
	"fmt"
	// "github.com/aneelraju/quref_prj/controllers"
	"encoding/json"
	"github.com/aneelraju/quref_prj/models"
	"io/ioutil"
	"os"
)

func main() {

	// Open data file (json format)
	jsonFile, err := os.Open("data/data.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully opened data/data.json")
	defer jsonFile.Close()

	// Read opened jsonFile as byte array
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// Initialize sections array
	var sections models.Sections

	// Unmarshal jsonFile's content to Sections
	json.Unmarshal(byteValue, &sections)

	// Iterate and print jsonFile values
	s := sections.Sections
	for i := 0; i < len(s); i++ {
		fmt.Println("Section Name: " + s[i].SectionName)
		f := sections.Sections[i].Features
		for j := 0; j < len(f); j++ {
			t := f[j].Topics
			fmt.Println("Feature Name: " + f[j].FeatureName)
			for k := 0; k < len(t); k++ {
				c := t[k].Comments
				fmt.Println("Topic Name: " + t[k].TopicName)
				for l := 0; l < len(c); l++ {
					fmt.Printf("Comment: <%s>:%s\n", c[l].CommentId, c[l].Comment)
				}
			}
		}
	}
}
