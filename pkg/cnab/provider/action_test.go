package cnabprovider

import (
	"io/ioutil"
	"testing"

	"github.com/cnabio/cnab-go/bundle"
	"github.com/cnabio/cnab-go/driver"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddReloccation(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/relocation-mapping.json")
	require.NoError(t, err)

	d := NewTestRuntime(t)

	args := ActionArguments{
		RelocationMapping: "/cnab/app/relocation-mapping.json",
	}

	d.TestConfig.TestContext.AddTestFile("testdata/relocation-mapping.json", "/cnab/app/relocation-mapping.json")

	opConf := d.AddRelocation(args)

	invoImage := bundle.InvocationImage{}
	invoImage.Image = "gabrtv/microservice@sha256:cca460afa270d4c527981ef9ca4989346c56cf9b20217dcea37df1ece8120687"

	op := &driver.Operation{
		Files: make(map[string]string),
		Image: invoImage,
	}
	err = opConf(op)
	assert.NoError(t, err)

	mapping, ok := op.Files["/cnab/app/relocation-mapping.json"]
	assert.True(t, ok)
	assert.Equal(t, string(data), mapping)
	assert.Equal(t, "my.registry/microservice@sha256:cca460afa270d4c527981ef9ca4989346c56cf9b20217dcea37df1ece8120687", op.Image.Image)

}
