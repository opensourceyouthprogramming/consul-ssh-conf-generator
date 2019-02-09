package consul2ssh

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetEnvKey(t *testing.T) {
	var testParameters = []struct {
		message  string
		keyVal   string
		defVal   string
		expected string
	}{
		{"Test non exists key", "", "def_key_val", "def_key_val"},
		{"Test exists key", "key_val", "def_key_val", "key_val"},
	}
	for _, tp := range testParameters {
		if tp.keyVal != "" {
			os.Setenv("FOO", tp.keyVal)
		}
		val := GetEnvKey("FOO", tp.defVal)
		if val != tp.expected {
			t.Errorf("%v: got %q, want %q.", tp.message, val, tp.expected)
		}
	}

}

func TestGetFilePath(t *testing.T) {
	workingDir, _ := os.Getwd()
	fileName := "dummy_file.txt"
	filePath := getFilePath(fileName)
	expectedFilePath := fmt.Sprintf("%s/%s", workingDir, fileName)
	assert.Equal(t, expectedFilePath, filePath)
}

func TestMergeMaps(t *testing.T) {
	sourceMap := mapInterface{
		"key01": "val",
		"key02": true,
		"key03": 1,
	}
	distMap := mapInterface{}
	mergeMaps(sourceMap, distMap)
	sourceMapLen := make([]int, len(sourceMap))
	for keyID := range sourceMapLen {
		key := fmt.Sprintf("key%02d", keyID)
		assert.Equal(t, sourceMap[key], distMap[key])
	}
}

func TestSetStrVal(t *testing.T) {
	var testParameters = []struct {
		message  string
		inVal    string
		defVal   string
		expected string
	}{
		{"Test empty val", "", "def_val", "def_val"},
		{"Test exists val", "val", "", "val"},
	}
	for _, tp := range testParameters {
		val := setStrVal(tp.inVal, tp.defVal)
		if val != tp.expected {
			t.Errorf("%v: got %q, want %q.", tp.message, val, tp.expected)
		}
	}

}
