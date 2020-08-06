package faker_test

import (
	"testing"

	"github.com/pioz/faker"
	"github.com/pioz/faker/data"
	"github.com/stretchr/testify/assert"
)

func TestGetAndSetPool(t *testing.T) {
	faker.SetSeed(802)
	faker.SetPool("ns1", "grp1", data.Pool{"foo", "bar"})
	value, err := faker.GetData("ns1", "grp1")
	assert.Nil(t, err)
	t.Log(value)
	assert.Equal(t, "foo", value)

	faker.SetPool("ns2", "grp1", data.Pool{"foo", "bar"})
	value, err = faker.GetData("ns2", "grp1")
	assert.Nil(t, err)
	t.Log(value)
	assert.Equal(t, "bar", value)

	faker.SetPool("ns2", "grp2", data.Pool{"foo", "bar"})
	value, err = faker.GetData("ns2", "grp2")
	assert.Nil(t, err)
	t.Log(value)
	assert.Equal(t, "foo", value)

	faker.SetPoolGroup("ns3", data.PoolGroup{"grp1": {"foo", "bar"}})
	value, err = faker.GetData("ns3", "grp1")
	assert.Nil(t, err)
	t.Log(value)
	assert.Equal(t, "bar", value)
}

func TestGetDataWithNamespaceError(t *testing.T) {
	faker.SetPool("ns4", "grp1", data.Pool{"foo", "bar"})
	_, err := faker.GetData("not-exist", "grp1")
	assert.NotNil(t, err)
	assert.Equal(t, "The namespace 'not-exist' does not exist", err.Error())
}

func TestGetDataWithGroupError(t *testing.T) {
	faker.SetPool("ns5", "grp1", data.Pool{"foo", "bar"})
	_, err := faker.GetData("ns5", "not-exist")
	assert.NotNil(t, err)
	assert.Equal(t, "The group 'not-exist' in namespace 'ns5' does not exist", err.Error())
}
