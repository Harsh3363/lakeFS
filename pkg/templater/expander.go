package templater

import (
	"github.com/treeverse/lakefs/pkg/auth"
	auth_model "github.com/treeverse/lakefs/pkg/auth/model"
	config "github.com/treeverse/lakefs/pkg/config"

	"context"
	"errors"
	"fmt"
	"path"
	//	html_template "html/template"
	"io"
	"io/fs"
	"strings"
	"text/template"
)

const (
	// Prefix inside templates FS to access actual contents.
	embeddedContentPrefix = "content"
)

var (
	ErrNotFound             = errors.New("template not found")
	ErrPathTraversalBlocked = errors.New("path traversal blocked")
)

type AuthService interface {
	auth.Authorizer
	auth.CredentialsCreator
}

type ControlledParams struct {
	Ctx  context.Context
	Auth AuthService
	// User is the user expanding the template.
	User *auth_model.User
}

type UncontrolledData struct {
	// UserName is the name of the executing user.
	Username string
	// Query is the (parsed) querystring of the HTTP access.
	Query map[string]string
}

// Params parametrizes a single template expansion.
type Params struct {
	// Controlled is the data visible to functions to control expansion.
	// It is _not_ directly visible to templates for expansion.
	Controlled *ControlledParams
	// Data is directly visible to templates for expansion, with no
	// authorization required.
	Data *UncontrolledData
}

// Expander is a template that may be expanded as requested by users.
type Expander interface {
	// Prepare checks whether the template expansion will succeed.  Call
	// it before Expand to ensure nothing is written when expansion fails.
	Prepare(params *Params) error

	// Expand returns an error or serves the template into w.  It may
	// write to w before returning an expansion error; call Prepare
	// first to avoid this.  (However Expand may still fail, for
	// instance if writing fails!)
	Expand(w io.Writer, params *Params) error
}

type expander struct {
	template *template.Template
	cfg      *config.Config
	auth     AuthService
}

// MakeExpander creates an expander for the text of tmpl.
func MakeExpander(name, tmpl string, cfg *config.Config, auth AuthService) (Expander, error) {
	t := template.New(name).Funcs(templateFuncs).Option("missingkey=error")
	t, err := t.Parse(tmpl)
	if err != nil {
		return nil, err
	}

	return &expander{
		template: t,
		cfg:      cfg,
		auth:     auth,
	}, nil
}

func (e *expander) Prepare(params *Params) error {
	return e.Expand(io.Discard, params)
}

func (e *expander) Expand(w io.Writer, params *Params) error {
	clone, err := e.template.Clone()
	if err != nil {
		return fmt.Errorf("Expand: %w", err)
	}
	wrappedFuncs := WrapFuncMapWithData(templateFuncs, params.Controlled)
	clone.Funcs(wrappedFuncs)
	return clone.Execute(w, params.Data)
}

// ExpanderMap reads and caches Expanders from a fs.FS.  Currently it
// provides no uncaching as it is only used with a prebuilt FS.
type ExpanderMap struct {
	fs   fs.FS
	cfg  *config.Config
	auth AuthService

	expanders map[string]Expander
}

func NewExpanderMap(fs fs.FS, cfg *config.Config, auth AuthService) *ExpanderMap {
	return &ExpanderMap{
		fs:        fs,
		cfg:       cfg,
		auth:      auth,
		expanders: make(map[string]Expander, 0),
	}
}

func (em *ExpanderMap) Get(ctx context.Context, username, name string) (Expander, error) {
	if e, ok := em.expanders[name]; ok {
		// Fast-path through the cache
		if e == nil {
			// Negative cache
			return nil, ErrNotFound
		}
		return e, nil
	}

	// Compute path
	p := path.Join(embeddedContentPrefix, name)
	if !strings.HasPrefix(p, embeddedContentPrefix+"/") {
		// Path traversal, fail
		return nil, fmt.Errorf("%s: %w", name, ErrPathTraversalBlocked)
	}

	tmpl, err := fs.ReadFile(em.fs, p)
	if errors.Is(err, fs.ErrNotExist) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	e, err := MakeExpander(name, string(tmpl), em.cfg, em.auth)
	if err != nil {
		// Store negative cache result
		e = nil
	}
	em.expanders[name] = e
	return e, err
}
