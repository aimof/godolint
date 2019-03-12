package rules

import (
	"testing"
)

func TestDL3015Check(t *testing.T) {
	cases := []struct {
		dockerfileStr string
		file          string
		expectedRst   []string
		expectedErr   error
	}{
		{
			dockerfileStr: `FROM busybox
RUN apt-get install -y python=2.7
`,
			file:        "DL3015Check_Dockerfile",
			expectedRst: []string{"DL3015Check_Dockerfile:2 DL3015 Avoid additional packages by specifying `--no-install-recommends`\n"},
			expectedErr: nil,
		},
		{
			dockerfileStr: `FROM busybox
RUN apt-get install -y python=2.7 && RUN apt-get install -y ruby
`,
			file:        "DL3015Check_Dockerfile_2",
			expectedRst: []string{"DL3015Check_Dockerfile_2:2 DL3015 Avoid additional packages by specifying `--no-install-recommends`\n"},
			expectedErr: nil,
		},
	}

	for i, tc := range cases {
		rst, err := dockerFileParse(tc.dockerfileStr)
		if err != nil {
			t.Errorf("#%d dl3015Check parse error %s", i, tc.dockerfileStr)
		}

		gotRst, gotErr := dl3015Check(rst.AST, tc.file)
		if !sliceEq(gotRst, tc.expectedRst) {
			t.Errorf("#%d dl3015Check results deep equal has returned: want %s, got %s", i, tc.expectedRst, gotRst)
		}

		if gotErr != tc.expectedErr {
			t.Errorf("#%d dl3015Check error has returned: want %s, got %s", i, tc.expectedErr, gotErr)
		}
		cleanup(t)
	}
}