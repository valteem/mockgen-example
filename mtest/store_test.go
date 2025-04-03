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

func TestFindProductDescription(t *testing.T) {

	mockctrl := gomock.NewController(t)
	defer mockctrl.Finish()

	pc := NewMockProductCatalog(mockctrl)

	p1 := NewMockProduct(mockctrl)
	p1.EXPECT().Description().Return("some product").AnyTimes()
	p2 := NewMockProduct(mockctrl)
	p2.EXPECT().Description().Return("some other product").AnyTimes()

	pc.EXPECT().Find(gomock.Eq(uint(1))).Return(p1).AnyTimes()
	pc.EXPECT().Find(gomock.Eq(uint(2))).Return(p2).AnyTimes()

	tests := []struct {
		id     uint
		output string
	}{
		{1, "some product"},
		{2, "some other product"},
	}

	for _, tc := range tests {

		output := storage.FindProductDescription(pc, tc.id)
		if output != tc.output {
			t.Errorf("Product description:\nget\n%s\nexpect\n%s\n", output, tc.output)
		}
	}

}
