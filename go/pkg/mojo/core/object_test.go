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

func TestObject_Delete(t *testing.T) {
    obj := NewObject().SetString("title", "test")
    obj.Delete("title")
    assert.True(t, obj.IsEmpty())
}

func TestObject_Merge(t *testing.T) {
    o1, _ := NewObjectFromKeyValues("x", 12, "y", "foo")
    o2, _ := NewObjectFromKeyValues("x", 16, "z", "bar")
    o1.Merge(o2)
    assert.Equal(t, 3, len(o1.Vals))
    assert.Equal(t, int64(16), o1.GetInt64("x"))
}

func TestMergeObjects(t *testing.T) {
    o1, _ := NewObjectFromKeyValues("x", 12, "y", "foo")
    o2, _ := NewObjectFromKeyValues("x", 16, "z", "bar")
    o3, _ := NewObjectFromKeyValues("a", "alpha", "b", "blue")

    o := MergeObjects(o1, o2, o3)
    assert.NotNil(t, o)
    assert.Equal(t, 5, len(o.Vals))
    assert.Equal(t, int64(16), o.GetInt64("x"))
}
