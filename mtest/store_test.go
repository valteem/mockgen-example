package mtest

import (
	"testing"

	"github.com/valteem/mockgen-example/storage"
	"go.uber.org/mock/gomock"
)

func TestProductDescription(t *testing.T) {

	mockctrl := gomock.NewController(t)
	defer mockctrl.Finish()

	p := NewMockProduct(mockctrl)

	tests := []struct {
		mock   func()
		output string
	}{
		{
			mock:   func() { p.EXPECT().Description().Return("some very good product") },
			output: "some very good product",
		},
	}

	for _, tc := range tests {

		tc.mock()

		output := storage.ProductDescription(p)

		if output != tc.output {
			t.Errorf("Product description:\nget\n%q\nexpect\n%q\n", output, tc.output)
		}

	}

}
