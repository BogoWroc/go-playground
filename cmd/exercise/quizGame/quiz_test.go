package main

import (
	"errors"
	"reflect"
	"testing"
)

func Test_readCsvFile(t *testing.T) {
	type args struct {
		pathToFile string
	}
	tests := []struct {
		name string
		args args
		want map[string]int
		error
	}{
		{
			name:  "Parse correct csv file",
			args:  args{"./testdata/sampleExpressions.csv"},
			want:  map[string]int{"1+1": 2, "1+2": 3, "1+4": 5},
			error: nil,
		},
		{
			name:  "Parse empty csv file",
			args:  args{"./testdata/empty.csv"},
			want:  map[string]int{},
			error: nil,
		},
		{
			name:  "Parse broken csv file",
			args:  args{"./testdata/broken.csv"},
			want:  nil,
			error: errors.New("line '[1+1 a]' contains invalid expression"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readCsvFile(tt.args.pathToFile)
			if err != nil && err.Error() != tt.error.Error() {
				t.Errorf("readCsvFile() -> gotError: %v, wantError: %v", err, tt.error)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readCsvFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shouldReportAnErrorWhenFileDoesNotExist(t *testing.T) {
	_, err := readCsvFile("fileDoesNotExist.csv")
	if err == nil {
		t.Error("File does not exist. Error should be reported!")
	}
}
