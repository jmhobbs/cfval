package resources

import (
	"testing"

	"github.com/jagregory/cfval/schema"
)

func TestCidrValidation(t *testing.T) {
	template := &schema.Template{}
	tr := schema.NewTemplateResource(template)
	context := []string{}

	if _, errs := cidr.Validate(schema.Schema{}, "", tr, context); errs == nil {
		t.Error("Cidr should fail on empty string")
	}

	if _, errs := cidr.Validate(schema.Schema{}, "abc", tr, context); errs == nil {
		t.Error("Cidr should fail on anything which isn't a cidr")
	}

	if _, errs := cidr.Validate(schema.Schema{}, "0.0.0.0/100", tr, context); errs == nil {
		t.Error("Cidr should fail on an invalid mask")
	}

	if _, errs := cidr.Validate(schema.Schema{}, "10.200.300.10/24", tr, context); errs == nil {
		t.Error("Cidr should fail on an invalid IP")
	}

	if _, errs := cidr.Validate(schema.Schema{}, "10.200.30.10/24", tr, context); errs != nil {
		t.Error("Cidr should pass with a valid cidr")
	}
}
