package mltext

import (
	"os"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"io/ioutil"
)

func TestHtmlText(t *testing.T) {
	file, err := os.OpenFile("evernote_test.xml", os.O_RDONLY, 0)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	expected, err := ioutil.ReadFile("evernote_expected_text.txt")
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	actual, err := ToText(file)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	assert.Equal(t, string(expected), actual)
}