package dicom

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jamesshenjian/dicom/pkg/tag"
)

func TestDataset_FindElementByTag(t *testing.T) {
	data := Dataset{
		Elements: map[tag.Tag]*Element{
			tag.Rows: {
				Tag:                 tag.Rows,
				ValueRepresentation: tag.VRInt32List,
				Value: &intsValue{
					value: []int{100},
				},
			},
			tag.Columns: &Element{
				Tag:                 tag.Columns,
				ValueRepresentation: tag.VRInt32List,
				Value: &intsValue{
					value: []int{200},
				},
			},
		},
	}

	elem, err := data.FindElementByTag(tag.Rows)
	if err != nil {
		t.Errorf("FindElementByTag(%v): unexpected err: %v", tag.Rows, err)
	}

	if rows := MustGetInts(elem.Value)[0]; rows != 100 {
		t.Errorf("FindElementByTag(%v): got: %v, want: %v", tag.Rows, rows, 100)
	}
}

func TestDataset_FlatStatefulIterator(t *testing.T) {
	cases := []struct {
		name                 string
		dataset              Dataset
		expectedFlatElements []*Element
	}{
		{
			//we do not allow same tag to appear multiple times in same file
			name: "flat dataset",
			dataset: Dataset{Elements: map[tag.Tag]*Element{
				tag.PatientName: MustNewElement(tag.PatientName, []string{"Bob", "Smith"}),
				//tag.PatientName: MustNewElement(tag.PatientName, []string{"Bob", "Jones"}),
			}},
			expectedFlatElements: []*Element{
				MustNewElement(tag.PatientName, []string{"Bob", "Smith"}),
				//MustNewElement(tag.PatientName, []string{"Bob", "Jones"}),
			},
		},
		{
			name: "nested dataset",
			dataset: Dataset{Elements: map[tag.Tag]*Element{
				tag.AddOtherSequence: MakeSequenceElement(tag.AddOtherSequence, [][]*Element{
					// Item 1
					{
						MustNewElement(tag.PatientName, []string{"Bob", "Jones"}),
						// Nested Sequence.
						MakeSequenceElement(tag.AnatomicRegionSequence, [][]*Element{
							{
								MustNewElement(tag.PatientName, []string{"Bob", "Smith"}),
							},
						}),
					},
				}),
			}},
			expectedFlatElements: []*Element{
				// First, expect the entire SQ element
				MakeSequenceElement(tag.AddOtherSequence, [][]*Element{
					// Item 1
					{
						MustNewElement(tag.PatientName, []string{"Bob", "Jones"}),
						// Nested Sequence.
						MakeSequenceElement(tag.AnatomicRegionSequence, [][]*Element{
							{
								MustNewElement(tag.PatientName, []string{"Bob", "Smith"}),
							},
						}),
					},
				}),
				// Then expect the inner elements
				MustNewElement(tag.PatientName, []string{"Bob", "Jones"}),
				// Inner SQ element
				MakeSequenceElement(tag.AnatomicRegionSequence, [][]*Element{
					{
						MustNewElement(tag.PatientName, []string{"Bob", "Smith"}),
					},
				}),
				// Inner element of the inner SQ
				MustNewElement(tag.PatientName, []string{"Bob", "Smith"}),
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var gotElems []*Element
			for iter := tc.dataset.FlatStatefulIterator(); iter.HasNext(); {
				gotElems = append(gotElems, iter.Next())
			}
			if diff := cmp.Diff(SortElements(tc.expectedFlatElements), SortElements(gotElems), cmp.AllowUnexported(allValues...)); diff != "" {
				t.Errorf("FlatStatefulIterator(%v) returned unexpected set of elements: %v", tc.dataset, diff)
			}
		})
	}
}

func ExampleDataset_FlatIterator() {
	nestedData := [][]*Element{
		{
			MustNewElement(tag.PatientName, []string{"Bob"}),
		},
	}

	data := Dataset{
		Elements: map[tag.Tag]*Element{
			tag.Rows:             MustNewElement(tag.Rows, []int{100}),
			tag.Columns:          MustNewElement(tag.Columns, []int{100}),
			tag.AddOtherSequence: MakeSequenceElement(tag.AddOtherSequence, nestedData),
		},
	}

	// Use this style if you will always exhaust all of the elements in the
	// channel. Otherwise, you must call ExhaustElementChannel. See the
	// FlatIteratorWithExhaustAllElements example for that. If you don't need
	// a channel API (just want to loop over items), use FlatStatefulIterator
	// instead, which is much simpler.
	for elem := range data.FlatIterator() {
		fmt.Println(elem.Tag)
	}

	// Note the output below includes all three leaf elements __as well as__ the sequence element's tag

	// Unordered output:
	// (0028,0010)
	// (0028,0011)
	// (0010,0010)
	// (0046,0102)
}

func ExampleDataset_FlatIteratorWithExhaustAllElements() {
	nestedData := [][]*Element{
		{
			MustNewElement(tag.PatientName, []string{"Bob"}),
		},
	}

	data := Dataset{
		Elements: map[tag.Tag]*Element{
			tag.Rows:             MustNewElement(tag.Rows, []int{100}),
			tag.Columns:          MustNewElement(tag.Columns, []int{100}),
			tag.AddOtherSequence: MakeSequenceElement(tag.AddOtherSequence, nestedData),
		},
	}

	// Because we read only one element from the channel, we want to make sure
	// the channel is exhausted to ensure it is closed properly under the hood.
	// This is also needed if you have any situation in which you may not
	// read all the elements in the channel (e.g. if you are looping over it,
	// but might return early if there's an error).
	elemChan := data.FlatIterator()
	defer ExhaustElementChannel(elemChan)
	fmt.Println((<-elemChan).Tag)

	// Note the output below includes all three leaf elements __as well as__ the sequence element's tag

	// Unordered output:
	// (0028,0010)
}

func ExampleDataset_FlatStatefulIterator() {
	nestedData := [][]*Element{
		{
			{
				Tag:                 tag.PatientName,
				ValueRepresentation: tag.VRString,
				Value: &stringsValue{
					value: []string{"Bob"},
				},
			},
		},
	}

	data := Dataset{
		Elements: map[tag.Tag]*Element{
			tag.Rows: {
				Tag:                 tag.Rows,
				ValueRepresentation: tag.VRInt32List,
				Value: &intsValue{
					value: []int{100},
				},
			},
			tag.Columns: {
				Tag:                 tag.Columns,
				ValueRepresentation: tag.VRInt32List,
				Value: &intsValue{
					value: []int{200},
				},
			},
			tag.AddOtherSequence: MakeSequenceElement(tag.AddOtherSequence, nestedData),
		},
	}

	for iter := data.FlatStatefulIterator(); iter.HasNext(); {
		fmt.Println(iter.Next().Tag)
	}

	// Note the output below includes all three leaf elements __as well as__ the sequence element's tag

	// Unordered output:
	// (0028,0010)
	// (0028,0011)
	// (0010,0010)
	// (0046,0102)
}

func ExampleDataset_String() {
	d := Dataset{
		Elements: map[tag.Tag]*Element{
			tag.Rows: {
				Tag:                    tag.Rows,
				ValueRepresentation:    tag.VRInt32List,
				RawValueRepresentation: "UL",
				Value: &intsValue{
					value: []int{100},
				},
			},
			tag.Columns: {
				Tag:                    tag.Columns,
				ValueRepresentation:    tag.VRInt32List,
				RawValueRepresentation: "UL",
				Value: &intsValue{
					value: []int{200},
				},
			},
		},
	}

	fmt.Println(d.String())

	// Output:
	// [
	//   Tag: (0028,0010)
	//   Tag Name: Rows
	//   VR: VRInt32List
	//   VR Raw: UL
	//   VL: 0
	//   Value: [100]
	// ]
	//
	// [
	//   Tag: (0028,0011)
	//   Tag Name: Columns
	//   VR: VRInt32List
	//   VR Raw: UL
	//   VL: 0
	//   Value: [200]
	// ]

}
