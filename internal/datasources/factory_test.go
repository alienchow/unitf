package datasources

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
)

type DummyModel struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	Age  types.Int64  `tfsdk:"age"`
}

type DummyDataSourceModel struct {
	SiteID types.String  `tfsdk:"site_id"`
	Filter []FilterModel `tfsdk:"filter"`
	Items  []DummyModel  `tfsdk:"items"`
}

func TestApplyFilters_NoFilter(t *testing.T) {
	model := DummyDataSourceModel{
		Items: []DummyModel{
			{ID: types.StringValue("1"), Name: types.StringValue("A")},
			{ID: types.StringValue("2"), Name: types.StringValue("B")},
		},
	}

	err := applyFilters(&model)
	assert.NoError(t, err)
	assert.Len(t, model.Items, 2)
}

func TestApplyFilters_Match(t *testing.T) {
	model := DummyDataSourceModel{
		Filter: []FilterModel{
			{
				Name: types.StringValue("name"),
				Values: []types.String{
					types.StringValue("B"),
				},
			},
		},
		Items: []DummyModel{
			{ID: types.StringValue("1"), Name: types.StringValue("A")},
			{ID: types.StringValue("2"), Name: types.StringValue("B")},
			{ID: types.StringValue("3"), Name: types.StringValue("C")},
		},
	}

	err := applyFilters(&model)
	assert.NoError(t, err)
	assert.Len(t, model.Items, 1)
	assert.Equal(t, "2", model.Items[0].ID.ValueString())
}

func TestApplyFilters_MultipleValues(t *testing.T) {
	model := DummyDataSourceModel{
		Filter: []FilterModel{
			{
				Name: types.StringValue("name"),
				Values: []types.String{
					types.StringValue("A"),
					types.StringValue("C"),
				},
			},
		},
		Items: []DummyModel{
			{ID: types.StringValue("1"), Name: types.StringValue("A")},
			{ID: types.StringValue("2"), Name: types.StringValue("B")},
			{ID: types.StringValue("3"), Name: types.StringValue("C")},
		},
	}

	err := applyFilters(&model)
	assert.NoError(t, err)
	assert.Len(t, model.Items, 2)
	assert.Equal(t, "1", model.Items[0].ID.ValueString())
	assert.Equal(t, "3", model.Items[1].ID.ValueString())
}

func TestApplyFilters_MultipleFilters(t *testing.T) {
	model := DummyDataSourceModel{
		Filter: []FilterModel{
			{
				Name: types.StringValue("name"),
				Values: []types.String{
					types.StringValue("A"),
				},
			},
			{
				Name: types.StringValue("id"),
				Values: []types.String{
					types.StringValue("1"),
				},
			},
		},
		Items: []DummyModel{
			{ID: types.StringValue("1"), Name: types.StringValue("A")},
			{ID: types.StringValue("2"), Name: types.StringValue("A")},
			{ID: types.StringValue("3"), Name: types.StringValue("C")},
		},
	}

	err := applyFilters(&model)
	assert.NoError(t, err)
	assert.Len(t, model.Items, 1)
	assert.Equal(t, "1", model.Items[0].ID.ValueString())
}
