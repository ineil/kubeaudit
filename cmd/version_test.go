package cmd

import (
	"bytes"
	"regexp"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestVersionCmd(t *testing.T) {
	buf := new(bytes.Buffer)
	log.SetOutput(buf)
	log.SetFormatter(&log.TextFormatter{})
	cmd := RootCmd
	cmd.SetArgs([]string{"version", "-l"})

	if err := cmd.Execute(); err != nil {
		t.Fatal("Version command returned an error.")
	}

	type exp struct {
		str string
		msg string
	}

	exps := []exp{
		{`.*level=info\s+msg=Kubeaudit\s+BuildDate=\S+\s+Commit=[[:xdigit:]]+\s+Version=\d+\.\d+\.\d+.*`, "missing kubeaudit version"},
		{`.*level=info\s+msg=\"Kubernetes server version\"\s+Major=\d+\s+Minor=\d+\+*\s+Platform=\S+`, "missing server version"},
	}

	for _, e := range exps {
		assert.Regexp(t, regexp.MustCompile(e.str), buf, e.msg)
	}
}
