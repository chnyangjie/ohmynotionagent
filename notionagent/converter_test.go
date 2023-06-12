package notionagent

import (
	"log"
	"testing"
	"time"

	"github.com/jomei/notionapi"
)

func TestRichText(t *testing.T) {
	t.Log("test RichTextProperty")
	property, err := VarToProperty("test", "")
	if err != nil {
		t.Fatalf("err: %+v", err)
	}
	if property == nil {
		t.Fatalf("property is nil")
	}
	typedProperty, ok := property.(*notionapi.RichTextProperty)
	if !ok {
		t.Fatalf("property is not a RichTextProperty")
	}
	text := typedProperty.RichText[0].Text.Content
	if text != "test" {
		t.Fatalf("typedProperty is not a RichTextProperty")
	}
	log.Printf("typedProperty: %+v", typedProperty)
}
func TestRichText2(t *testing.T) {
	t.Log("test RichTextProperty")
	fieldType := notionapi.PropertyTypeRichText
	property, err := VarToProperty("test", fieldType)
	if err != nil {
		t.Fatalf("err: %+v", err)
	}
	if property == nil {
		t.Fatalf("property is nil")
	}
	typedProperty, ok := property.(*notionapi.RichTextProperty)
	if !ok {
		t.Fatalf("property is not a RichTextProperty")
	}
	text := typedProperty.RichText[0].Text.Content
	if text != "test" {
		t.Fatalf("typedProperty is not a RichTextProperty")
	}
	log.Printf("typedProperty: %+v", typedProperty)
}
func TestTitle(t *testing.T) {
	t.Log("test TitleProperty")
	fieldType := notionapi.PropertyTypeTitle
	property, err := VarToProperty("test", fieldType)
	if err != nil {
		t.Fatalf("err: %+v", err)
	}
	if property == nil {
		t.Fatalf("property is nil")
	}
	typedProperty, ok := property.(*notionapi.TitleProperty)
	if !ok {
		t.Fatalf("property is not a TitleProperty")
	}
	text := typedProperty.Title[0].Text.Content
	if text != "test" {
		t.Fatalf("typedProperty is not a TitleProperty")
	}
	log.Printf("typedProperty: %+v", typedProperty)
}

func TestMultiSelect(t *testing.T) {
	t.Log("test TitleProperty")
	fieldType := notionapi.PropertyTypeMultiSelect
	property, err := VarToProperty([]string{"test"}, fieldType)
	if err != nil {
		t.Fatalf("err: %+v", err)
	}
	if property == nil {
		t.Fatalf("property is nil")
	}
	typedProperty, ok := property.(*notionapi.MultiSelectProperty)
	if !ok {
		t.Fatalf("property is not a TitleProperty")
	}
	text := typedProperty.MultiSelect[0].Name
	if text != "test" {
		t.Fatalf("typedProperty is not a TitleProperty")
	}
	log.Printf("typedProperty: %+v", typedProperty)
}

func TestDate(t *testing.T) {
	t.Log("test TitleProperty")
	fieldType := notionapi.PropertyTypeDate
	now := time.Now()
	property, err := VarToProperty([]time.Time{now}, fieldType)
	if err != nil {
		t.Fatalf("err: %+v", err)
	}
	if property == nil {
		t.Fatalf("property is nil")
	}
	typedProperty, ok := property.(*notionapi.DateProperty)
	if !ok {
		t.Fatalf("property is not a TitleProperty")
	}
	text := *typedProperty.Date.Start
	if time.Time(text) != now {
		t.Fatalf("typedProperty is not a TitleProperty")
	}
	log.Printf("typedProperty: %+v", typedProperty)
}
func TestUrl(t *testing.T) {
	t.Log("test TitleProperty")
	fieldType := notionapi.PropertyTypeURL
	property, err := VarToProperty("http://baidu.com", fieldType)
	if err != nil {
		t.Fatalf("err: %+v", err)
	}
	if property == nil {
		t.Fatalf("property is nil")
	}
	typedProperty, ok := property.(*notionapi.URLProperty)
	if !ok {
		t.Fatalf("property is not a TitleProperty")
	}
	text := typedProperty.URL
	if text != "http://baidu.com" {
		t.Fatalf("typedProperty is not a TitleProperty")
	}
	log.Printf("typedProperty: %+v", typedProperty)
}
