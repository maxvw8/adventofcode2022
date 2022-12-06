package advent

import "testing"

func TestSectionPair_isOverlap(t *testing.T) {
	type fields struct {
		e1 elfAssignment
		e2 elfAssignment
	}
	tests := []struct {
		name string
		e1   elfAssignment
		e2   elfAssignment
		want bool
	}{
		{name: "2-32,1-55", e1: elfAssignment{from: 2, to: 32}, e2: elfAssignment{from: 1, to: 55}, want: true},
		{name: "39-91,38-90", e1: elfAssignment{from: 39, to: 91}, e2: elfAssignment{from: 38, to: 90}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SectionPair{
				e1: tt.e1,
				e2: tt.e2,
			}
			if got := s.isOverlap(); got != tt.want {
				t.Errorf("isOverlap() = %v, want %v", got, tt.want)
			}
		})
	}
}
