package subcmd

import (
	"testing"

	"github.com/onsi/gomega"
	"github.com/redhat-appstudio/helmet/internal/flags"
	"github.com/redhat-appstudio/helmet/internal/mcptools"
)

func noopMCPToolsBuilder(_ mcptools.MCPToolsContext) ([]mcptools.Interface, error) {
	return nil, nil
}

func TestMCPServer_Complete_EmptyImage_ReturnsError(t *testing.T) {
	t.Parallel()
	g := gomega.NewWithT(t)

	appCtx := testAppContext()
	runCtx := testRunContext(t)
	manager := testManager(t, runCtx)

	m := NewMCPServer(appCtx, runCtx, flags.NewFlags(), manager, noopMCPToolsBuilder, "")
	err := m.Complete(nil)
	g.Expect(err).To(gomega.HaveOccurred())
	g.Expect(err.Error()).To(gomega.ContainSubstring("WithImage()"))
}

func TestMCPServer_Complete_WithImage_Succeeds(t *testing.T) {
	t.Parallel()
	g := gomega.NewWithT(t)

	appCtx := testAppContext()
	runCtx := testRunContext(t)
	manager := testManager(t, runCtx)

	m := NewMCPServer(appCtx, runCtx, flags.NewFlags(), manager, noopMCPToolsBuilder, "test:latest")
	g.Expect(m.Complete(nil)).ToNot(gomega.HaveOccurred())
}

func TestMCPServer_Complete_ImageOverriddenByFlag(t *testing.T) {
	t.Parallel()
	g := gomega.NewWithT(t)

	appCtx := testAppContext()
	runCtx := testRunContext(t)
	manager := testManager(t, runCtx)

	m := NewMCPServer(appCtx, runCtx, flags.NewFlags(), manager, noopMCPToolsBuilder, "")
	err := m.cmd.PersistentFlags().Set("image", "override:latest")
	g.Expect(err).ToNot(gomega.HaveOccurred())
	g.Expect(m.Complete(nil)).ToNot(gomega.HaveOccurred())
}

func TestMCPServer_Validate_AlwaysSucceeds(t *testing.T) {
	t.Parallel()
	g := gomega.NewWithT(t)

	appCtx := testAppContext()
	runCtx := testRunContext(t)
	manager := testManager(t, runCtx)

	m := NewMCPServer(appCtx, runCtx, flags.NewFlags(), manager, noopMCPToolsBuilder, "test:latest")
	g.Expect(m.Validate()).ToNot(gomega.HaveOccurred())
}
