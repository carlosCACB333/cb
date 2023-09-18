package utils

import (
	"fmt"
	"strings"

	"github.com/tmc/langchaingo/schema"
)

func ProcessDocPage(content schema.Document) {

	fmt.Println(content.PageContent)
	content.PageContent = strings.ReplaceAll(content.PageContent, "\n", "")
	fmt.Println(content.PageContent)

}
