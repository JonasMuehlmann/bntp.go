package helpers_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/JonasMuehlmann/bntp.go/internal/helpers"
	"github.com/JonasMuehlmann/bntp.go/test"
	"github.com/stretchr/testify/assert"
)

func TestUnMarshallFromWhole(t *testing.T) {
	marshaled := []byte(`{"wrapee": 123, "has_value": true}`)

	var optional helpers.Optional[int]

	err := json.Unmarshal(marshaled, &optional)
	assert.NoError(t, err)
	assert.Equal(t, 123, optional.Wrappee)
	assert.Equal(t, true, optional.HasValue)
}

func TestUnMarshallFromWrappee(t *testing.T) {
	targetStruct := struct {
		Foo        string                `json:"foo"`
		MyOptional helpers.Optional[int] `json:"my_optional"`
	}{}
	marshaled, err := json.Marshal(map[string]any{"foo": "bar", "my_optional": 123})
	assert.NoError(t, err)

	err = json.Unmarshal(marshaled, &targetStruct)
	assert.NoError(t, err)
	assert.Equal(t, 123, targetStruct.MyOptional.Wrappee)
	assert.Equal(t, true, targetStruct.MyOptional.HasValue)
}

// ******************************************************************//
//                           fmt.Stringer                           //
// ******************************************************************//.
func TestOptionalToStringHasValue(t *testing.T) {
	optional := helpers.Optional[int]{Wrappee: 123, HasValue: true}

	stringified := fmt.Sprint(optional)
	assert.Equal(t, "123", stringified)
}

func TestOptionalToStringHasNoValue(t *testing.T) {
	optional := helpers.Optional[int]{Wrappee: 123, HasValue: false}

	stringified := fmt.Sprint(optional)
	assert.Equal(t, "empty optional", stringified)
}

// ******************************************************************//
//        database.sql.Scanner and database.sql.driver.Valuer       //
// ******************************************************************//.
func TestSQLScanAndValueHasValue(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	_, err = db.Exec("CREATE TABLE FOO(BAR TEXT)")
	assert.NoError(t, err)

	input := helpers.Optional[string]{Wrappee: "bar", HasValue: true}
	_, err = db.Exec("INSERT INTO FOO VALUES(?)", input)
	assert.NoError(t, err)

	var output helpers.Optional[string]
	err = db.Get(&output, "SELECT * FROM Foo")
	assert.NoError(t, err)
	assert.Equal(t, input, output)
}

func TestSQLScanAndValueHasNoValue(t *testing.T) {
	db, err := test.GetDB(t)
	assert.NoError(t, err)

	_, err = db.Exec("CREATE TABLE FOO(BAR TEXT)")
	assert.NoError(t, err)

	input := helpers.Optional[string]{Wrappee: "bar", HasValue: true}
	_, err = db.Exec("INSERT INTO FOO VALUES(?)", input)
	assert.NoError(t, err)

	var output helpers.Optional[string]
	err = db.Get(&output, "SELECT * FROM Foo")
	assert.NoError(t, err)
	assert.Equal(t, input, output)
}
