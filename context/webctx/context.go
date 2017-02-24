package webctx

import (
	"fmt"

	"github.com/sclevine/agouti"

	"github.com/runabove/venom"
)

// Context Type name
const Name = "web"

// Key of context element in testsuite file
const (
	Width      = "width"
	Height     = "height"
	Screenshot = "screenshotOnFailure"
)

// Key of element in the testcase context
const (
	ContextDriverKey           = "driver"
	ContextPageKey             = "page"
	ContextScreenshotOnFailure = "screenshotOnFailure"
)

// New returns a new TestCaseContext
func New() venom.TestCaseContext {
	return &TestCaseContext{}
}

// TestCaseContex represents the context of a testcase
type TestCaseContext struct {
	venom.TestCaseContextStruct
	wd   *agouti.WebDriver
	Page *agouti.Page
}

// BuildContext build context of type web.
// It creates a new browser
func (tcc *TestCaseContext) Init() error {
	// Init web driver
	tcc.wd = agouti.PhantomJS()
	if err := tcc.wd.Start(); err != nil {
		return fmt.Errorf("Cannot start web driver %s", err)
	}

	// Init Page
	var errP error
	tcc.Page, errP = tcc.wd.NewPage()
	if errP != nil {
		return fmt.Errorf("Cannot create new page %s", errP)
	}

	resizePage := false
	if _, ok := tcc.TestCase.Context[Width]; ok {
		if _, ok := tcc.TestCase.Context[Height]; ok {
			resizePage = true
		}
	}

	// Resize Page
	if resizePage {
		var width, height int
		switch tcc.TestCase.Context[Width].(type) {
		case int:
			width = tcc.TestCase.Context[Width].(int)
		default:
			return fmt.Errorf("%s is not an integer: %s", Width, fmt.Sprintf("%s", tcc.TestCase.Context[Width]))
		}
		switch tcc.TestCase.Context[Height].(type) {
		case int:
			height = tcc.TestCase.Context[Height].(int)
		default:
			return fmt.Errorf("%s is not an integer: %s", Height, fmt.Sprintf("%s", tcc.TestCase.Context[Height]))
		}

		if err := tcc.Page.Size(width, height); err != nil {
			return fmt.Errorf("Cannot resize page: %s", err)
		}
	}
	return nil
}

// Close web driver
func (tcc *TestCaseContext) Close() error {
	return tcc.wd.Stop()
}
