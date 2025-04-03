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

func TestStorageRoomeShare(t *testing.T) {

	mockctrl := gomock.NewController(t)
	defer mockctrl.Finish()

	room := NewMockStorageRoom(mockctrl)
	room.EXPECT().Capacity(gomock.Any()).Return(float32(100.)).AnyTimes()

	tests := []struct {
		floor  int
		share  float32
		output float32
	}{
		{1, 0.5, 50.}, {2, 0.5, 50.}, {3, 0.75, 75.},
	}

	for _, tc := range tests {

		output := storage.StorageShare(room, tc.floor, tc.share)

		if output != tc.output {
			t.Errorf("Storage share for floor %d and share %.2f: get %.0f, expect %.0f",
				tc.floor,
				tc.share,
				output,
				tc.output,
			)
		}
	}

}

func TestItemsStored(t *testing.T) {

	mockctrl := gomock.NewController(t)

	counter := NewMockStoredItemsCount(mockctrl)

	first := counter.EXPECT().Count().Return(1)
	second := counter.EXPECT().Count().Return(2).After(first)
	counter.EXPECT().Count().Return(42).After(second)

	calls := []int{1, 2, 42}

	for _, c := range calls {
		output := storage.ItemsStored(counter)
		if output != c {
			t.Errorf("get %d, expect %d", output, c)
		}
	}

}

func TestItemsStoredWithInOrder(t *testing.T) {

	mockctrl := gomock.NewController(t)

	counter := NewMockStoredItemsCount(mockctrl)

	gomock.InOrder(
		counter.EXPECT().Count().Return(1),
		counter.EXPECT().Count().Return(2),
		counter.EXPECT().Count().Return(42),
	)

	// does not work:
	// expected call has already been called the max number of times
	// calls := []int{1, 2, 42, 42}
	calls := []int{1, 2, 42}

	for _, c := range calls {
		output := storage.ItemsStored(counter)
		if output != c {
			t.Errorf("get %d, expect %d", output, c)
		}
	}

}
