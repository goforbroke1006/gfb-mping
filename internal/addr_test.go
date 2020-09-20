package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAddressesList(t *testing.T) {
	type args struct {
		r string
	}
	tests := []struct {
		name                 string
		args                 args
		wantContainsInResult []string
		wantErr              bool
	}{
		{
			name: "positive, addresses C-type",
			args: args{r: "192.168.1.1/24"},
			wantContainsInResult: []string{
				"192.168.1.0",
				"192.168.1.1",
				"192.168.1.2",
				"192.168.1.3",
				"192.168.1.4",
				"192.168.1.255",
			},
			wantErr: false,
		},
		{
			name: "positive, addresses B-type",
			args: args{r: "172.20.0.1/16"},
			wantContainsInResult: []string{
				"172.20.0.1",
				"172.20.0.2",
				"172.20.0.3",
				"172.20.1.1",
				"172.20.2.1",
				"172.20.3.1",
				"172.20.128.128",
				"172.20.255.255",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := GetAddressesList(tt.args.r)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetAddressesList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			for _, wr := range tt.wantContainsInResult {
				if !assert.Contains(t, gotResult, wr) {
					t.Errorf("GetAddressesList() gotResult does not contains %v", wr)
				}
			}
		})
	}
}
