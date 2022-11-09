package review_test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/aruncveli/gerritr/pkg/review"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
)

func TestMain(m *testing.M) {
	fPath := filepath.Join("..", "..", "test", "config.yml")
	fProvider := file.Provider(fPath)
	err := review.Config.Load(fProvider, yaml.Parser())
	if err != nil {
		fmt.Println("Reading test config failed", err)
	}

	m.Run()
}

func TestWithoutREVIEWERS(t *testing.T) {
	type args struct {
		revIp []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"No reviewers", args{nil}, nil},

		{"Single team from config", args{[]string{"backend"}},
			[]string{"r=b1@org.com", "r=b2@org.com"}},

		{"Multiple teams from config", args{[]string{"backend", "frontend"}},
			[]string{"r=b1@org.com", "r=b2@org.com", "r=f1@org.com", "r=f2@org.com"}},

		{"Reviwer email as input", args{[]string{"x@com.org"}}, []string{"r=x@com.org"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := review.ResolveReviewers(tt.args.revIp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetReviewers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithREVIEWERS(t *testing.T) {
	d1 := []byte("x1@org.com\nx2@org.com\n")
	os.WriteFile("REVIEWERS", d1, 0644)

	type args struct {
		revIp []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Specify no reviewers",
			args{nil},
			[]string{"r=x1@org.com", "r=x2@org.com"}},

		{"Specify single team from config",
			args{[]string{"backend"}},
			[]string{"r=b1@org.com", "r=b2@org.com", "r=x1@org.com", "r=x2@org.com"}},

		{"Specify multiple teams from config",
			args{[]string{"backend", "frontend"}},
			[]string{"r=b1@org.com", "r=b2@org.com", "r=f1@org.com", "r=f2@org.com", "r=x1@org.com", "r=x2@org.com"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := review.ResolveReviewers(tt.args.revIp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetReviewers() = %v, want %v", got, tt.want)
			}
		})
	}

	os.Remove("REVIEWERS")
}
