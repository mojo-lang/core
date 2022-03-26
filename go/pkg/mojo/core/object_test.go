package core

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestObject_From(t *testing.T) {
    image := &Image{
        Title: "test",
        Type:  "png",
        Properties: map[string]*Value{
            "p1": NewInt32Value(1234),
        },
    }

    object := &Object{}
    err := object.From(image)
    assert.NoError(t, err)

    assert.Equal(t, image.Title, object.GetString("title"))
    assert.Equal(t, int32(1234), object.GetValue("properties").GetObject().GetInt32("p1"))
}

func TestObject_To(t *testing.T) {
    obj := NewObject().SetString("title", "test").
        SetString("type", "png").
        SetObject("properties", NewObject().SetInt32("p1", 1234))

    image := &Image{}
    err := obj.To(image)
    assert.NoError(t, err)

    assert.Equal(t, "test", image.Title)
    assert.Equal(t, int32(1234), image.Properties["p1"].GetInt32())
}
