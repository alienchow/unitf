package datasources

import (
	"context"
	"fmt"
	"reflect"

	"github.com/alienchow/unitf/internal/client"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type FilterModel struct {
	Name   types.String   `tfsdk:"name"`
	Values []types.String `tfsdk:"values"`
}

type GenericDataSource[TModel any] struct {
	TypeName string
	TFSchema schema.Schema
	ReadFunc func(ctx context.Context, c *client.Client, model *TModel) error

	client *client.Client
}

func (d *GenericDataSource[TModel]) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + d.TypeName
}

func (d *GenericDataSource[TModel]) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = d.TFSchema
}

func (d *GenericDataSource[TModel]) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Data Source Configure Type", "Expected *client.Client")
		return
	}
	d.client = c
}

func (d *GenericDataSource[TModel]) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var model TModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := d.ReadFunc(ctx, d.client, &model)
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("Error Reading %s", d.TypeName), fmt.Sprintf("%v", err))
		return
	}

	if err := applyFilters(&model); err != nil {
		resp.Diagnostics.AddError("Error Filtering", err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func applyFilters(model any) error {
	v := reflect.ValueOf(model)
	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}

	filterField := getFieldByTfsdkTag(v, "filter")
	if !filterField.IsValid() || filterField.IsNil() || filterField.Len() == 0 {
		return nil
	}

	var filters []FilterModel
	for i := 0; i < filterField.Len(); i++ {
		val := filterField.Index(i).Interface()
		if f, ok := val.(FilterModel); ok {
			filters = append(filters, f)
		}
	}

	if len(filters) == 0 {
		return nil
	}

	itemsField := getFieldByTfsdkTag(v, "items")
	if !itemsField.IsValid() || itemsField.IsNil() {
		return nil
	}

	filteredItems := reflect.MakeSlice(itemsField.Type(), 0, itemsField.Len())
	for i := 0; i < itemsField.Len(); i++ {
		item := itemsField.Index(i)
		if matchFilters(item, filters) {
			filteredItems = reflect.Append(filteredItems, item)
		}
	}

	if itemsField.CanSet() {
		itemsField.Set(filteredItems)
	}

	return nil
}

func matchFilters(item reflect.Value, filters []FilterModel) bool {
	if item.Kind() == reflect.Pointer {
		item = item.Elem()
	}

	for _, filter := range filters {
		fieldName := filter.Name.ValueString()
		fieldVal := getFieldByTfsdkTag(item, fieldName)
		if !fieldVal.IsValid() {
			return false
		}

		var strVal string
		switch v := fieldVal.Interface().(type) {
		case types.String:
			strVal = v.ValueString()
		case types.Int64:
			strVal = fmt.Sprintf("%d", v.ValueInt64())
		case types.Float64:
			strVal = fmt.Sprintf("%g", v.ValueFloat64())
		case types.Bool:
			strVal = fmt.Sprintf("%t", v.ValueBool())
		default:
			strVal = fmt.Sprintf("%v", v)
		}

		match := false
		for _, allowed := range filter.Values {
			if strVal == allowed.ValueString() {
				match = true
				break
			}
		}
		if !match {
			return false
		}
	}
	return true
}

func getFieldByTfsdkTag(v reflect.Value, tagName string) reflect.Value {
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if tag := field.Tag.Get("tfsdk"); tag == tagName {
			return v.Field(i)
		}
	}
	return reflect.Value{}
}
