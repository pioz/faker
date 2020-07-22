package factory_test

import (
	"testing"

	"github.com/pioz/factory"
	"github.com/pioz/factory/data"
	"github.com/stretchr/testify/assert"
)

func TestGetAndSetPool(t *testing.T) {
	factory.SetSeed(802)
	factory.SetPool("ns1", "grp1", data.Pool{"foo", "bar"})
	value, err := factory.GetData("ns1", "grp1")
	assert.Nil(t, err)
	t.Log(value)
	assert.Equal(t, "foo", value)

	factory.SetPool("ns2", "grp1", data.Pool{"foo", "bar"})
	value, err = factory.GetData("ns2", "grp1")
	assert.Nil(t, err)
	t.Log(value)
	assert.Equal(t, "bar", value)

	factory.SetPool("ns2", "grp2", data.Pool{"foo", "bar"})
	value, err = factory.GetData("ns2", "grp2")
	assert.Nil(t, err)
	t.Log(value)
	assert.Equal(t, "foo", value)

	factory.SetPoolGroup("ns3", data.PoolGroup{"grp1": {"foo", "bar"}})
	value, err = factory.GetData("ns3", "grp1")
	assert.Nil(t, err)
	t.Log(value)
	assert.Equal(t, "bar", value)
}

func TestGetDataWithNamespaceError(t *testing.T) {
	factory.SetPool("ns4", "grp1", data.Pool{"foo", "bar"})
	_, err := factory.GetData("not-exist", "grp1")
	assert.NotNil(t, err)
	assert.Equal(t, "The namespace 'not-exist' does not exist", err.Error())
}

func TestGetDataWithGroupError(t *testing.T) {
	factory.SetPool("ns5", "grp1", data.Pool{"foo", "bar"})
	_, err := factory.GetData("ns5", "not-exist")
	assert.NotNil(t, err)
	assert.Equal(t, "The group 'not-exist' in namespace 'ns5' does not exist", err.Error())
}
