package action

import (
	"reflect"
	"testing"

	"github.com/aruncveli/gerritr/pkg/git"
)

func TestPatch(t *testing.T) {
	// Mock/Hack
	git.ExecName = "echo"

	type args struct {
		branch string
		state  string
		msg    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{

		{"No edit", args{"", "", ""},
			"commit --amend --no-edit\npush origin HEAD:refs/for/main%\n"},

		{"With message", args{"", "", "message"},
			"commit --amend --message message\npush origin HEAD:refs/for/main%\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotB := Patch(tt.args.branch, tt.args.state, tt.args.msg)
			gotS := string(gotB)
			if !reflect.DeepEqual(gotS, tt.want) {
				t.Errorf("Patch() = %v, want %v", gotS, tt.want)
			}
		})
	}
}
