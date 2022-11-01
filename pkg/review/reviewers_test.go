package review_test

import (
	"fmt"
	"gerritr/pkg/review"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
)

func TestMain(m *testing.M) {
	fPath := filepath.Join("..", "..", "test", "config.yml")
	fProvider := file.Provider(fPath)
	err := review.Config.Load(fProvider, yaml.Parser())
	if err != nil {
		fmt.Println("Reading config failed", err)
	}

	review.SetAllowedEmailDomains()

	m.Run()
	os.Remove("REVIEWERS")
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
			[]string{"r=b1@yes.com", "r=b2@yes.com"}},

		{"Multiple teams from config", args{[]string{"backend", "frontend"}},
			[]string{"r=b1@yes.com", "r=b2@yes.com", "r=f1@yes.com", "r=f2@yes.com"}},

		{"Invalid email", args{[]string{"someone@no.com"}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := review.GetReviewers(tt.args.revIp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetReviewers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithREVIEWERS(t *testing.T) {
	d1 := []byte("x1@yes.com\nx2@yes.com\n")
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
			[]string{"r=x1@yes.com", "r=x2@yes.com"}},

		{"Specify single team from config",
			args{[]string{"backend"}},
			[]string{"r=b1@yes.com", "r=b2@yes.com", "r=x1@yes.com", "r=x2@yes.com"}},

		{"Specify multiple teams from config",
			args{[]string{"backend", "frontend"}},
			[]string{"r=b1@yes.com", "r=b2@yes.com", "r=f1@yes.com", "r=f2@yes.com", "r=x1@yes.com", "r=x2@yes.com"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := review.GetReviewers(tt.args.revIp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetReviewers() = %v, want %v", got, tt.want)
			}
		})
	}

	os.Remove("REVIEWERS")
}

func TestWithoutEmailCfg(t *testing.T) {
	os.Remove("REVIEWERS")
	review.Config.Delete("allowedEmailDomains")
	review.SetAllowedEmailDomains()

	type args struct {
		revIp []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Single team from config and a random email",
			args{[]string{"someome@no.com", "backend"}},
			[]string{"r=someome@no.com", "r=b1@yes.com", "r=b2@yes.com"}},
		{"Random email",
			args{[]string{"someone@no.com"}},
			[]string{"r=someone@no.com"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := review.GetReviewers(tt.args.revIp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetReviewers() = %v, want %v", got, tt.want)
			}
		})
	}
}
