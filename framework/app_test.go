package framework

import (
	"os"
	"testing"

	"github.com/redhat-appstudio/helmet/api"
	"github.com/redhat-appstudio/helmet/internal/chartfs"
)

func testAppContext() *api.AppContext {
	return api.NewAppContext(
		"helmet-test",
		api.WithNamespace("test-ns"),
	)
}

func testChartFS(t *testing.T) *chartfs.ChartFS {
	t.Helper()
	return chartfs.New(os.DirFS("../test/charts"))
}

func TestNewApp_WithoutImage_Succeeds(t *testing.T) {
	t.Parallel()
	app, err := NewApp(testAppContext(), testChartFS(t))
	if err != nil {
		t.Fatalf("NewApp without WithImage should succeed, got: %v", err)
	}
	if app == nil {
		t.Fatal("expected non-nil App")
	}
}

func TestWithImage_SetsImage(t *testing.T) {
	t.Parallel()
	app, err := NewApp(testAppContext(), testChartFS(t),
		WithImage("test:latest"),
	)
	if err != nil {
		t.Fatalf("NewApp with WithImage should succeed, got: %v", err)
	}
	if app == nil {
		t.Fatal("expected non-nil App")
	}
	if app.image != "test:latest" {
		t.Fatalf("expected image %q, got %q", "test:latest", app.image)
	}
}

func TestWithMCPImage_DeprecatedAlias(t *testing.T) {
	t.Parallel()
	app, err := NewApp(testAppContext(), testChartFS(t),
		WithMCPImage("deprecated:latest"),
	)
	if err != nil {
		t.Fatalf("NewApp with WithMCPImage should succeed, got: %v", err)
	}
	if app == nil {
		t.Fatal("expected non-nil App")
	}
	if app.image != "deprecated:latest" {
		t.Fatalf("expected image %q, got %q", "deprecated:latest", app.image)
	}
}
