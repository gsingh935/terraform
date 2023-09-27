// Code generated by "stringer -type=Progress progress.go"; DO NOT EDIT.

package moduletest

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Starting-0]
	_ = x[TearDown-1]
	_ = x[Complete-2]
}

const _Progress_name = "StartingTearDownComplete"

var _Progress_index = [...]uint8{0, 8, 16, 24}

func (i Progress) String() string {
	if i < 0 || i >= Progress(len(_Progress_index)-1) {
		return "Progress(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Progress_name[_Progress_index[i]:_Progress_index[i+1]]
}