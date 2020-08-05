package ddb_test

import (
	"encoding/json"
	"fmt"
	"github.com/applike/gosoline/pkg/ddb"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

type TestModel struct {
	MyName         int    `json:"myName" ddb:"key=hash"`
	MyNameWithAttr string `json:"myNameWithAttr,omitempty" ddb:"global=hash"`
	MyNamePlain    string `json:",omitempty" ddb:"key=range"`
	MyIgnoredField int64  `json:"-"`
	DashField      int    `json:"-," ddb:"global=range"`
	DefaultField   uint   `ddb:"global=range"`
}

type TestModelEmptyDDB struct {
	Field int `ddb:""`
}

type TestModelNoKV struct {
	Field int `ddb:"not a value"`
}

type TestModelEmptyJSON struct {
	Field int `json:"" ddb:"key=hash"`
}

type CustomModel struct {
	IntField          int          `ddb:"key=int"`
	Int8Field         int8         `ddb:"key=int8"`
	Int16Field        int16        `ddb:"key=int16"`
	Int32Field        int32        `ddb:"key=int32"`
	Int64Field        int64        `ddb:"key=int64"`
	UintField         uint         `ddb:"key=uint"`
	Uint8Field        uint8        `ddb:"key=uint8"`
	Uint16Field       uint16       `ddb:"key=uint16"`
	Uint32Field       uint32       `ddb:"key=uint32"`
	Uint64Field       uint64       `ddb:"key=uint64"`
	Float32Field      float32      `ddb:"key=float32"`
	Float64Field      float64      `ddb:"key=float64"`
	StringField       string       `ddb:"key=string"`
	TimeField         time.Time    `ddb:"key=time"`
	CustomTimeField   CustomTime   `ddb:"key=customTime"`
	CustomIntField    CustomInt    `ddb:"key=customInt"`
	CustomStringField CustomString `ddb:"key=customString"`
	CustomStructField CustomStruct `ddb:"key=customStruct"`
}

type CustomTime struct {
	time.Time
}

type CustomInt int

type CustomString string

type CustomStruct struct {
	a string
	b string
}

func (f CustomStruct) MarshalJSON() ([]byte, error) {
	return json.Marshal(fmt.Sprintf("%s|%s", f.a, f.b))
}

func (f *CustomStruct) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)

	if err == nil {
		f.a = strings.Split(s, "|")[0]
		f.b = strings.Split(s, "|")[1]
	}

	return err
}

func TestReadAttributes(t *testing.T) {
	attributes, err := ddb.ReadAttributes(TestModel{})
	assert.NoError(t, err)
	assert.Equal(t, ddb.Attributes{
		"myName": &ddb.Attribute{
			FieldName:     "MyName",
			AttributeName: "myName",
			Tags: map[string]string{
				"key": "hash",
			},
			Type: "N",
		},
		"myNameWithAttr": &ddb.Attribute{
			FieldName:     "MyNameWithAttr",
			AttributeName: "myNameWithAttr",
			Tags: map[string]string{
				"global": "hash",
			},
			Type: "S",
		},
		"MyNamePlain": &ddb.Attribute{
			FieldName:     "MyNamePlain",
			AttributeName: "MyNamePlain",
			Tags: map[string]string{
				"key": "range",
			},
			Type: "S",
		},
		"-": &ddb.Attribute{
			FieldName:     "DashField",
			AttributeName: "-",
			Tags: map[string]string{
				"global": "range",
			},
			Type: "N",
		},
		"DefaultField": &ddb.Attribute{
			FieldName:     "DefaultField",
			AttributeName: "DefaultField",
			Tags: map[string]string{
				"global": "range",
			},
			Type: "N",
		},
	}, attributes)
	_, err = ddb.ReadAttributes(TestModelEmptyDDB{})
	assert.Error(t, err)
	_, err = ddb.ReadAttributes(TestModelEmptyJSON{})
	assert.Error(t, err)
	_, err = ddb.ReadAttributes(TestModelNoKV{})
	assert.Error(t, err)
	attributes, err = ddb.ReadAttributes(CustomModel{})
	assert.NoError(t, err)
	expectedAttributes := ddb.Attributes{}
	for _, s := range []string{
		"int",
		"int8",
		"int16",
		"int32",
		"int64",
		"uint",
		"uint8",
		"uint16",
		"uint32",
		"uint64",
		"float32",
		"float64",
		"customInt",
	} {
		fieldName := fmt.Sprintf("%s%sField", strings.ToUpper(s[0:1]), s[1:])
		expectedAttributes[fieldName] = &ddb.Attribute{
			FieldName:     fieldName,
			AttributeName: fieldName,
			Tags: map[string]string{
				"key": strings.ToLower(s),
			},
			Type: "N",
		}
	}
	for _, s := range []string{
		"string",
		"time",
		"customTime",
		"customString",
		"customStruct",
	} {
		fieldName := fmt.Sprintf("%s%sField", strings.ToUpper(s[0:1]), s[1:])
		expectedAttributes[fieldName] = &ddb.Attribute{
			FieldName:     fieldName,
			AttributeName: fieldName,
			Tags: map[string]string{
				"key": strings.ToLower(s),
			},
			Type: "S",
		}
	}
	assert.Equal(t, expectedAttributes, attributes)
}

func TestMetadataReadFields(t *testing.T) {
	fields, err := ddb.MetadataReadFields(TestModel{})
	assert.NoError(t, err)
	assert.Equal(t, []string{
		"myName",
		"myNameWithAttr",
		"MyNamePlain",
		// MyIgnoredField is... ignored, so there is no entry for it
		"-",
		"DefaultField",
	}, fields)
	_, err = ddb.MetadataReadFields(TestModelEmptyJSON{})
	assert.Error(t, err)
}
