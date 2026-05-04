package installer

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"os"
	"testing"

	"github.com/redhat-appstudio/helmet/internal/flags"
	"github.com/redhat-appstudio/helmet/internal/resolver"

	o "github.com/onsi/gomega"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chartutil"
)

func newTestInstaller() *Installer {
	dep := resolver.NewDependencyWithNamespace(
		&chart.Chart{Metadata: &chart.Metadata{Name: "test-chart"}},
		"test-ns",
	)
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	return NewInstaller(logger, flags.NewFlags(), nil, dep, nil)
}

func TestSetRenderedValues(t *testing.T) {
	g := o.NewWithT(t)

	i := newTestInstaller()

	inputBytes := []byte("key: value\n")
	inputValues := chartutil.Values{"key": "value"}

	i.SetRenderedValues(inputBytes, inputValues)

	g.Expect(i.valuesBytes).To(o.Equal(inputBytes))
	g.Expect(i.values).To(o.Equal(inputValues))
}

func TestSetRenderedValues_PrintRawValues(t *testing.T) {
	g := o.NewWithT(t)

	i := newTestInstaller()

	inputBytes := []byte("key: value\n")
	inputValues := chartutil.Values{"key": "value"}
	i.SetRenderedValues(inputBytes, inputValues)

	old := os.Stdout
	r, w, err := os.Pipe()
	g.Expect(err).To(o.Succeed())
	t.Cleanup(func() {
		os.Stdout = old
		_ = r.Close()
		_ = w.Close()
	})

	os.Stdout = w
	i.PrintRawValues()
	os.Stdout = old
	g.Expect(w.Close()).To(o.Succeed())

	var buf bytes.Buffer
	_, err = io.Copy(&buf, r)
	g.Expect(err).To(o.Succeed())

	g.Expect(buf.String()).To(o.ContainSubstring("key: value"))
}

func TestRenderValues_RequiresSetValues(t *testing.T) {
	g := o.NewWithT(t)

	i := newTestInstaller()

	err := i.RenderValues()
	g.Expect(err).To(o.HaveOccurred())
	g.Expect(err).To(o.MatchError(fmt.Errorf("values not set")))
}
