package action_test

import (
	"gerritr/pkg/action"
	"gerritr/pkg/git"
	"reflect"
	"testing"
)

func TestPush(t *testing.T) {
	git.ExecName = "echo"

	type args struct {
		branch    string
		state     string
		msg       string
		reviewers []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Default values",
			args{"", "", "", nil},
			"push origin HEAD:refs/for/main%\n"},

		{"Custom target branch",
			args{"develop", "", "", nil},
			"push origin HEAD:refs/for/develop%\n"},

		{"Custom change state",
			args{"", "wip", "", nil},
			"push origin HEAD:refs/for/main%wip\n"},

		{"Reviewers",
			args{"", "", "", []string{"a@a.a", "b@b.b"}},
			"push origin HEAD:refs/for/main%,r=a@a.a,r=b@b.b\n"},

		{"With message",
			args{"", "", "message", nil},
			"commit --message message\npush origin HEAD:refs/for/main%\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotB := action.Push(tt.args.branch, tt.args.state, tt.args.msg, tt.args.reviewers)
			gotS := string(gotB)
			if !reflect.DeepEqual(gotS, tt.want) {
				t.Errorf("Push() = %v, want %v", gotS, tt.want)
			}
		})
	}
}
