package resources

import (
	"testing"

	"github.com/jagregory/cfval/schema"
)

func TestPeriodValidation(t *testing.T) {
	template := &schema.Template{}
	tr := schema.NewTemplateResource(template)
	context := []string{}

	if _, errs := period(schema.Schema{}, "", tr, context); errs == nil {
		t.Error("Period should fail on empty string")
	}

	if _, errs := period(schema.Schema{}, "abc", tr, context); errs == nil {
		t.Error("Period should fail on anything which isn't a period")
	}

	for _, ex := range []string{"0", "10", "119", "260"} {
		if _, errs := period(schema.Schema{}, ex, tr, context); errs == nil {
			t.Errorf("Period should fail on number which isn't a multiple of 60 (ex: %s)", ex)
		}
	}

	for _, ex := range []string{"60", "120", "240"} {
		if _, errs := period(schema.Schema{}, ex, tr, context); errs != nil {
			t.Errorf("Cidr should pass with numbers which are multiples of 60 (ex: %s)", ex)
		}
	}
}
