package container

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/sarulabs/di/v2"
	"github.com/sarulabs/dingo/v4"

	providerPkg "github.com/avptp/brain/internal/services/provider"

	slog "log/slog"

	config "github.com/avptp/brain/internal/config"
	data "github.com/avptp/brain/internal/generated/data"
	messaging "github.com/avptp/brain/internal/messaging"
	ses "github.com/aws/aws-sdk-go/service/ses"
	hcaptcha "github.com/kataras/hcaptcha"
	in "github.com/nicksnyder/go-i18n/v2/i18n"
	realclientipgo "github.com/realclientip/realclientip-go"
)

// C retrieves a Container from an interface.
// The function panics if the Container can not be retrieved.
//
// The interface can be :
//   - a *Container
//   - an *http.Request containing a *Container in its context.Context
//     for the dingo.ContainerKey("dingo") key.
//
// The function can be changed to match the needs of your application.
var C = func(i interface{}) *Container {
	if c, ok := i.(*Container); ok {
		return c
	}
	r, ok := i.(*http.Request)
	if !ok {
		panic("could not get the container with dic.C()")
	}
	c, ok := r.Context().Value(dingo.ContainerKey("dingo")).(*Container)
	if !ok {
		panic("could not get the container from the given *http.Request in dic.C()")
	}
	return c
}

type builder struct {
	builder *di.Builder
}

// NewBuilder creates a builder that can be used to create a Container.
// You probably should use NewContainer to create the container directly.
// But using NewBuilder allows you to redefine some di services.
// This can be used for testing.
// But this behavior is not safe, so be sure to know what you are doing.
func NewBuilder(scopes ...string) (*builder, error) {
	if len(scopes) == 0 {
		scopes = []string{di.App, di.Request, di.SubRequest}
	}
	b, err := di.NewBuilder(scopes...)
	if err != nil {
		return nil, fmt.Errorf("could not create di.Builder: %v", err)
	}
	provider := &providerPkg.Provider{}
	if err := provider.Load(); err != nil {
		return nil, fmt.Errorf("could not load definitions with the Provider (Provider from github.com/avptp/brain/internal/services/provider): %v", err)
	}
	for _, d := range getDiDefs(provider) {
		if err := b.Add(d); err != nil {
			return nil, fmt.Errorf("could not add di.Def in di.Builder: %v", err)
		}
	}
	return &builder{builder: b}, nil
}

// Add adds one or more definitions in the Builder.
// It returns an error if a definition can not be added.
func (b *builder) Add(defs ...di.Def) error {
	return b.builder.Add(defs...)
}

// Set is a shortcut to add a definition for an already built object.
func (b *builder) Set(name string, obj interface{}) error {
	return b.builder.Set(name, obj)
}

// Build creates a Container in the most generic scope.
func (b *builder) Build() *Container {
	return &Container{ctn: b.builder.Build()}
}

// NewContainer creates a new Container.
// If no scope is provided, di.App, di.Request and di.SubRequest are used.
// The returned Container has the most generic scope (di.App).
// The SubContainer() method should be called to get a Container in a more specific scope.
func NewContainer(scopes ...string) (*Container, error) {
	b, err := NewBuilder(scopes...)
	if err != nil {
		return nil, err
	}
	return b.Build(), nil
}

// Container represents a generated dependency injection container.
// It is a wrapper around a di.Container.
//
// A Container has a scope and may have a parent in a more generic scope
// and children in a more specific scope.
// Objects can be retrieved from the Container.
// If the requested object does not already exist in the Container,
// it is built thanks to the object definition.
// The following attempts to get this object will return the same object.
type Container struct {
	ctn di.Container
}

// Scope returns the Container scope.
func (c *Container) Scope() string {
	return c.ctn.Scope()
}

// Scopes returns the list of available scopes.
func (c *Container) Scopes() []string {
	return c.ctn.Scopes()
}

// ParentScopes returns the list of scopes wider than the Container scope.
func (c *Container) ParentScopes() []string {
	return c.ctn.ParentScopes()
}

// SubScopes returns the list of scopes that are more specific than the Container scope.
func (c *Container) SubScopes() []string {
	return c.ctn.SubScopes()
}

// Parent returns the parent Container.
func (c *Container) Parent() *Container {
	if p := c.ctn.Parent(); p != nil {
		return &Container{ctn: p}
	}
	return nil
}

// SubContainer creates a new Container in the next sub-scope
// that will have this Container as parent.
func (c *Container) SubContainer() (*Container, error) {
	sub, err := c.ctn.SubContainer()
	if err != nil {
		return nil, err
	}
	return &Container{ctn: sub}, nil
}

// SafeGet retrieves an object from the Container.
// The object has to belong to this scope or a more generic one.
// If the object does not already exist, it is created and saved in the Container.
// If the object can not be created, it returns an error.
func (c *Container) SafeGet(name string) (interface{}, error) {
	return c.ctn.SafeGet(name)
}

// Get is similar to SafeGet but it does not return the error.
// Instead it panics.
func (c *Container) Get(name string) interface{} {
	return c.ctn.Get(name)
}

// Fill is similar to SafeGet but it does not return the object.
// Instead it fills the provided object with the value returned by SafeGet.
// The provided object must be a pointer to the value returned by SafeGet.
func (c *Container) Fill(name string, dst interface{}) error {
	return c.ctn.Fill(name, dst)
}

// UnscopedSafeGet retrieves an object from the Container, like SafeGet.
// The difference is that the object can be retrieved
// even if it belongs to a more specific scope.
// To do so, UnscopedSafeGet creates a sub-container.
// When the created object is no longer needed,
// it is important to use the Clean method to delete this sub-container.
func (c *Container) UnscopedSafeGet(name string) (interface{}, error) {
	return c.ctn.UnscopedSafeGet(name)
}

// UnscopedGet is similar to UnscopedSafeGet but it does not return the error.
// Instead it panics.
func (c *Container) UnscopedGet(name string) interface{} {
	return c.ctn.UnscopedGet(name)
}

// UnscopedFill is similar to UnscopedSafeGet but copies the object in dst instead of returning it.
func (c *Container) UnscopedFill(name string, dst interface{}) error {
	return c.ctn.UnscopedFill(name, dst)
}

// Clean deletes the sub-container created by UnscopedSafeGet, UnscopedGet or UnscopedFill.
func (c *Container) Clean() error {
	return c.ctn.Clean()
}

// DeleteWithSubContainers takes all the objects saved in this Container
// and calls the Close function of their Definition on them.
// It will also call DeleteWithSubContainers on each child and remove its reference in the parent Container.
// After deletion, the Container can no longer be used.
// The sub-containers are deleted even if they are still used in other goroutines.
// It can cause errors. You may want to use the Delete method instead.
func (c *Container) DeleteWithSubContainers() error {
	return c.ctn.DeleteWithSubContainers()
}

// Delete works like DeleteWithSubContainers if the Container does not have any child.
// But if the Container has sub-containers, it will not be deleted right away.
// The deletion only occurs when all the sub-containers have been deleted manually.
// So you have to call Delete or DeleteWithSubContainers on all the sub-containers.
func (c *Container) Delete() error {
	return c.ctn.Delete()
}

// IsClosed returns true if the Container has been deleted.
func (c *Container) IsClosed() bool {
	return c.ctn.IsClosed()
}

// SafeGetCaptcha retrieves the "captcha" object from the request scope.
//
// ---------------------------------------------
//
//	name: "captcha"
//	type: *hcaptcha.Client
//	scope: "request"
//	build: func
//	params:
//		- "0": Service(*config.Config) ["config"]
//	unshared: false
//	close: false
//
// ---------------------------------------------
//
// If the object can not be retrieved, it returns an error.
func (c *Container) SafeGetCaptcha() (*hcaptcha.Client, error) {
	i, err := c.ctn.SafeGet("captcha")
	if err != nil {
		var eo *hcaptcha.Client
		return eo, err
	}
	o, ok := i.(*hcaptcha.Client)
	if !ok {
		return o, errors.New("could get 'captcha' because the object could not be cast to *hcaptcha.Client")
	}
	return o, nil
}

// GetCaptcha retrieves the "captcha" object from the request scope.
//
// ---------------------------------------------
//
//	name: "captcha"
//	type: *hcaptcha.Client
//	scope: "request"
//	build: func
//	params:
//		- "0": Service(*config.Config) ["config"]
//	unshared: false
//	close: false
//
// ---------------------------------------------
//
// If the object can not be retrieved, it panics.
func (c *Container) GetCaptcha() *hcaptcha.Client {
	o, err := c.SafeGetCaptcha()
	if err != nil {
		panic(err)
	}
	return o
}

// UnscopedSafeGetCaptcha retrieves the "captcha" object from the request scope.
//
// ---------------------------------------------
//
//	name: "captcha"
//	type: *hcaptcha.Client
//	scope: "request"
//	build: func
//	params:
//		- "0": Service(*config.Config) ["config"]
//	unshared: false
//	close: false
//
// ---------------------------------------------
//
// This method can be called even if request is a sub-scope of the container.
// If the object can not be retrieved, it returns an error.
func (c *Container) UnscopedSafeGetCaptcha() (*hcaptcha.Client, error) {
	i, err := c.ctn.UnscopedSafeGet("captcha")
	if err != nil {
		var eo *hcaptcha.Client
		return eo, err
	}
	o, ok := i.(*hcaptcha.Client)
	if !ok {
		return o, errors.New("could get 'captcha' because the object could not be cast to *hcaptcha.Client")
	}
	return o, nil
}

// UnscopedGetCaptcha retrieves the "captcha" object from the request scope.
//
// ---------------------------------------------
//
//	name: "captcha"
//	type: *hcaptcha.Client
//	scope: "request"
//	build: func
//	params:
//		- "0": Service(*config.Config) ["config"]
//	unshared: false
//	close: false
//
// ---------------------------------------------
//
// This method can be called even if request is a sub-scope of the container.
// If the object can not be retrieved, it panics.
func (c *Container) UnscopedGetCaptcha() *hcaptcha.Client {
	o, err := c.UnscopedSafeGetCaptcha()
	if err != nil {
		panic(err)
	}
	return o
}

// Captcha retrieves the "captcha" object from the request scope.
//
// ---------------------------------------------
//
//	name: "captcha"
//	type: *hcaptcha.Client
//	scope: "request"
//	build: func
//	params:
//		- "0": Service(*config.Config) ["config"]
//	unshared: false
//	close: false
//
// ---------------------------------------------
//
// It tries to find the container with the C method and the given interface.
// If the container can be retrieved, it calls the GetCaptcha method.
// If the container can not be retrieved, it panics.
func Captcha(i interface{}) *hcaptcha.Client {
	return C(i).GetCaptcha()
}

// SafeGetConfig retrieves the "config" object from the main scope.
//
// ---------------------------------------------
//
//	name: "config"
//	type: *config.Config
//	scope: "main"
//	build: func
//	params: nil
//	unshared: false
//	close: false
//
// ---------------------------------------------
//
// If the object can not be retrieved, it returns an error.
func (c *Container) SafeGetConfig() (*config.Config, error) {
	i, err := c.ctn.SafeGet("config")
	if err != nil {
		var eo *config.Config
		return eo, err
	}
	o, ok := i.(*config.Config)
	if !ok {
		return o, errors.New("could get 'config' because the object could not be cast to *config.Config")
	}
	return o, nil
}

// GetConfig retrieves the "config" object from the main scope.
//
// ---------------------------------------------
//
//	name: "config"
//	type: *config.Config
//	scope: "main"
//	build: func
//	params: nil
//	unshared: false
//	close: false
//
// ---------------------------------------------
//
// If the object can not be retrieved, it panics.
func (c *Container) GetConfig() *config.Config {
	o, err := c.SafeGetConfig()
	if err != nil {
		panic(err)
	}
	return o
}

// UnscopedSafeGetConfig retrieves the "config" object from the main scope.
//
// ---------------------------------------------
//
//	name: "config"
//	type: *config.Config
//	scope: "main"
//	build: func
//	params: nil
//	unshared: false
//	close: false
//
// ---------------------------------------------
//
// This method can be called even if main is a sub-scope of the container.
// If the object can not be retrieved, it returns an error.
func (c *Container) UnscopedSafeGetConfig() (*config.Config, error) {
	i, err := c.ctn.UnscopedSafeGet("config")
	if err != nil {
		var eo *config.Config
		return eo, err
	}
	o, ok := i.(*config.Config)
	if !ok {
		return o, errors.New("could get 'config' because the object could not be cast to *config.Config")
	}
	return o, nil
}

// UnscopedGetConfig retrieves the "config" object from the main scope.
//
// ---------------------------------------------
//
//	name: "config"
//	type: *config.Config
//	scope: "main"
//	build: func
//	params: nil
//	unshared: false
//	close: false
//
// ---------------------------------------------
//
// This method can be called even if main is a sub-scope of the container.
// If the object can not be retrieved, it panics.
func (c *Container) UnscopedGetConfig() *config.Config {
	o, err := c.UnscopedSafeGetConfig()
	if err != nil {
		panic(err)
	}
	return o
}

// Config retrieves the "config" object from the main scope.
//
// ---------------------------------------------
//
//	name: "config"
//	type: *config.Config
//	scope: "main"
//	build: func
//	params: nil
//	unshared: false
//	close: false
//
// ---------------------------------------------
//
// It tries to find the container with the C method and the given interface.
// If the container can be retrieved, it calls the GetConfig method.
// If the container can not be retrieved, it panics.
func Config(i interface{}) *config.Config {
	return C(i).GetConfig()
}

// SafeGetData retrieves the "data" object from the main scope.
//
// ---------------------------------------------
//
//	name: "data"
//	type: *data.Client
//	scope: "main"
//	build: func
//	params:
//		- "0": Service(*config.Config) ["config"]
//	unshared: false
//	close: true
//
// ---------------------------------------------
//
// If the object can not be retrieved, it returns an error.
func (c *Container) SafeGetData() (*data.Client, error) {
	i, err := c.ctn.SafeGet("data")
	if err != nil {
		var eo *data.Client
		return eo, err
	}
	o, ok := i.(*data.Client)
	if !ok {
		return o, errors.New("could get 'data' because the object could not be cast to *data.Client")
	}
	return o, nil
}

// GetData retrieves the "data" object from the main scope.
//
// ---------------------------------------------
//
//	name: "data"
//	type: *data.Client
//	scope: "main"
//	build: func
//	params:
//		- "0": Service(*config.Config) ["config"]
//	unshared: false
//	close: true
//
// ---------------------------------------------
//
// If the object can not be retrieved, it panics.
func (c *Container) GetData() *data.Client {
	o, err := c.SafeGetData()
	if err != nil {
		panic(err)
	}
	return o
}

// UnscopedSafeGetData retrieves the "data" object from the main scope.
//
// ---------------------------------------------
//
//	name: "data"
//	type: *data.Client
//	scope: "main"
//	build: func
//	params:
//		- "0": Service(*config.Config) ["config"]
//	unshared: false
//	close: true
//
// ---------------------------------------------
//
// This method can be called even if main is a sub-scope of the container.
// If the object can not be retrieved, it returns an error.
func (c *Container) UnscopedSafeGetData() (*data.Client, error) {
	i, err := c.ctn.UnscopedSafeGet("data")
	if err != nil {
		var eo *data.Client
		return eo, err
	}
	o, ok := i.(*data.Client)
	if !ok {
		return o, errors.New("could get 'data' because the object could not be cast to *data.Client")
	}
	return o, nil
}

// UnscopedGetData retrieves the "data" object from the main scope.
//
// ---------------------------------------------
//
//	name: "data"
//	type: *data.Client
//	scope: "main"
//	build: func
//	params:
//		- "0": Service(*config.Config) ["config"]
//	unshared: false
//	close: true
//
// ---------------------------------------------
//
// This method can be called even if main is a sub-scope of the container.
// If the object can not be retrieved, it panics.
func (c *Container) UnscopedGetData() *data.Client {
	o, err := c.UnscopedSafeGetData()
	if err != nil {
		panic(err)
	}
	return o
}

// Data retrieves the "data" object from the main scope.
//
// ---------------------------------------------
//
//	name: "data"
//	type: *data.Client
//	scope: "main"
//	build: func
//	params:
//		- "0": Service(*config.Config) ["config"]
//	unshared: false
//	close: true
//
// ---------------------------------------------
//
// It tries to find the container with the C method and the given interface.
// If the container can be retrieved, it calls the GetData method.
// If the container can not be retrieved, it panics.
func Data(i interface{}) *data.Client {
	return C(i).GetData()
}

// SafeGetI18n retrieves the "i18n" object from the app scope.
//
// ---------------------------------------------
//
//	name: "i18n"
//	type: *in.Bundle
//	scope: "app"
//	build: func
//	params: nil
//	unshared: false
//	close: false
//
// ---------------------------------------------
//
// If the object can not be retrieved, it returns an error.
func (c *Container) SafeGetI18n() (*in.Bundle, error) {
	i, err := c.ctn.SafeGet("i18n")
	if err != nil {
		var eo *in.Bundle
		return eo, err
	}
	o, ok := i.(*in.Bundle)
	if !ok {
		return o, errors.New("could get 'i18n' because the object could not be cast to *in.Bundle")
	}
	return o, nil
}

// GetI18n retrieves the "i18n" object from the app scope.
//
// ---------------------------------------------
//
//	name: "i18n"
//	type: *in.Bundle
//	scope: "app"
//	build: func
//	params: nil
//	unshared: false
//	close: false
//
// ---------------------------------------------
//
// If the object can not be retrieved, it panics.
func (c *Container) GetI18n() *in.Bundle {
	o, err := c.SafeGetI18n()
	if err != nil {
		panic(err)
	}
	return o
}

// UnscopedSafeGetI18n retrieves the "i18n" object from the app scope.
//
// ---------------------------------------------
//
//	name: "i18n"
//	type: *in.Bundle
//	scope: "app"
//	build: func
//	params: nil
//	unshared: false
//	close: false
//
// ---------------------------------------------
//
// This method can be called even if app is a sub-scope of the container.
// If the object can not be retrieved, it returns an error.
func (c *Container) UnscopedSafeGetI18n() (*in.Bundle, error) {
	i, err := c.ctn.UnscopedSafeGet("i18n")
	if err != nil {
		var eo *in.Bundle
		return eo, err
	}
	o, ok := i.(*in.Bundle)
	if !ok {
		return o, errors.New("could get 'i18n' because the object could not be cast to *in.Bundle")
	}
	return o, nil
}

// UnscopedGetI18n retrieves the "i18n" object from the app scope.
//
// ---------------------------------------------
//
//	name: "i18n"
//	type: *in.Bundle
//	scope: "app"
//	build: func
//	params: nil
//	unshared: false
//	close: false
//
// ---------------------------------------------
//
// This method can be called even if app is a sub-scope of the container.
// If the object can not be retrieved, it panics.
func (c *Container) UnscopedGetI18n() *in.Bundle {
	o, err := c.UnscopedSafeGetI18n()
	if err != nil {
		panic(err)
	}
	return o
}

// I18n retrieves the "i18n" object from the app scope.
//
// ---------------------------------------------
//
//	name: "i18n"
//	type: *in.Bundle
//	scope: "app"
//	build: func
//	params: nil
//	unshared: false
//	close: false
//
// ---------------------------------------------
//
// It tries to find the container with the C method and the given interface.
// If the container can be retrieved, it calls the GetI18n method.
// If the container can not be retrieved, it panics.
func I18n(i interface{}) *in.Bundle {
	return C(i).GetI18n()
}

// SafeGetIpStrategy retrieves the "ipStrategy" object from the main scope.
//
// ---------------------------------------------
//
//	name: "ipStrategy"
//	type: realclientipgo.Strategy
//	scope: "main"
//	build: func
//	params:
//		- "0": Service(*config.Config) ["config"]
//	unshared: false
//	close: false
//
// ---------------------------------------------
//
// If the object can not be retrieved, it returns an error.
func (c *Container) SafeGetIpStrategy() (realclientipgo.Strategy, error) {
	i, err := c.ctn.SafeGet("ipStrategy")
	if err != nil {
		var eo realclientipgo.Strategy
		return eo, err
	}
	o, ok := i.(realclientipgo.Strategy)
	if !ok {
		return o, errors.New("could get 'ipStrategy' because the object could not be cast to realclientipgo.Strategy")
	}
	return o, nil
}

// GetIpStrategy retrieves the "ipStrategy" object from the main scope.
//
// ---------------------------------------------
//
//	name: "ipStrategy"
//	type: realclientipgo.Strategy
//	scope: "main"
//	build: func
//	params:
//		- "0": Service(*config.Config) ["config"]
//	unshared: false
//	close: false
//
// ---------------------------------------------
//
// If the object can not be retrieved, it panics.
func (c *Container) GetIpStrategy() realclientipgo.Strategy {
	o, err := c.SafeGetIpStrategy()
	if err != nil {
		panic(err)
	}
	return o
}

// UnscopedSafeGetIpStrategy retrieves the "ipStrategy" object from the main scope.
//
// ---------------------------------------------
//
//	name: "ipStrategy"
//	type: realclientipgo.Strategy
//	scope: "main"
//	build: func
//	params:
//		- "0": Service(*config.Config) ["config"]
//	unshared: false
//	close: false
//
// ---------------------------------------------
//
// This method can be called even if main is a sub-scope of the container.
// If the object can not be retrieved, it returns an error.
func (c *Container) UnscopedSafeGetIpStrategy() (realclientipgo.Strategy, error) {
	i, err := c.ctn.UnscopedSafeGet("ipStrategy")
	if err != nil {
		var eo realclientipgo.Strategy
		return eo, err
	}
	o, ok := i.(realclientipgo.Strategy)
	if !ok {
		return o, errors.New("could get 'ipStrategy' because the object could not be cast to realclientipgo.Strategy")
	}
	return o, nil
}

// UnscopedGetIpStrategy retrieves the "ipStrategy" object from the main scope.
//
// ---------------------------------------------
//
//	name: "ipStrategy"
//	type: realclientipgo.Strategy
//	scope: "main"
//	build: func
//	params:
//		- "0": Service(*config.Config) ["config"]
//	unshared: false
//	close: false
//
// ---------------------------------------------
//
// This method can be called even if main is a sub-scope of the container.
// If the object can not be retrieved, it panics.
func (c *Container) UnscopedGetIpStrategy() realclientipgo.Strategy {
	o, err := c.UnscopedSafeGetIpStrategy()
	if err != nil {
		panic(err)
	}
	return o
}

// IpStrategy retrieves the "ipStrategy" object from the main scope.
//
// ---------------------------------------------
//
//	name: "ipStrategy"
//	type: realclientipgo.Strategy
//	scope: "main"
//	build: func
//	params:
//		- "0": Service(*config.Config) ["config"]
//	unshared: false
//	close: false
//
// ---------------------------------------------
//
// It tries to find the container with the C method and the given interface.
// If the container can be retrieved, it calls the GetIpStrategy method.
// If the container can not be retrieved, it panics.
func IpStrategy(i interface{}) realclientipgo.Strategy {
	return C(i).GetIpStrategy()
}

// SafeGetLogger retrieves the "logger" object from the main scope.
//
// ---------------------------------------------
//
//	name: "logger"
//	type: *slog.Logger
//	scope: "main"
//	build: func
//	params:
//		- "0": Service(*config.Config) ["config"]
//	unshared: false
//	close: false
//
// ---------------------------------------------
//
// If the object can not be retrieved, it returns an error.
func (c *Container) SafeGetLogger() (*slog.Logger, error) {
	i, err := c.ctn.SafeGet("logger")
	if err != nil {
		var eo *slog.Logger
		return eo, err
	}
	o, ok := i.(*slog.Logger)
	if !ok {
		return o, errors.New("could get 'logger' because the object could not be cast to *slog.Logger")
	}
	return o, nil
}

// GetLogger retrieves the "logger" object from the main scope.
//
// ---------------------------------------------
//
//	name: "logger"
//	type: *slog.Logger
//	scope: "main"
//	build: func
//	params:
//		- "0": Service(*config.Config) ["config"]
//	unshared: false
//	close: false
//
// ---------------------------------------------
//
// If the object can not be retrieved, it panics.
func (c *Container) GetLogger() *slog.Logger {
	o, err := c.SafeGetLogger()
	if err != nil {
		panic(err)
	}
	return o
}

// UnscopedSafeGetLogger retrieves the "logger" object from the main scope.
//
// ---------------------------------------------
//
//	name: "logger"
//	type: *slog.Logger
//	scope: "main"
//	build: func
//	params:
//		- "0": Service(*config.Config) ["config"]
//	unshared: false
//	close: false
//
// ---------------------------------------------
//
// This method can be called even if main is a sub-scope of the container.
// If the object can not be retrieved, it returns an error.
func (c *Container) UnscopedSafeGetLogger() (*slog.Logger, error) {
	i, err := c.ctn.UnscopedSafeGet("logger")
	if err != nil {
		var eo *slog.Logger
		return eo, err
	}
	o, ok := i.(*slog.Logger)
	if !ok {
		return o, errors.New("could get 'logger' because the object could not be cast to *slog.Logger")
	}
	return o, nil
}

// UnscopedGetLogger retrieves the "logger" object from the main scope.
//
// ---------------------------------------------
//
//	name: "logger"
//	type: *slog.Logger
//	scope: "main"
//	build: func
//	params:
//		- "0": Service(*config.Config) ["config"]
//	unshared: false
//	close: false
//
// ---------------------------------------------
//
// This method can be called even if main is a sub-scope of the container.
// If the object can not be retrieved, it panics.
func (c *Container) UnscopedGetLogger() *slog.Logger {
	o, err := c.UnscopedSafeGetLogger()
	if err != nil {
		panic(err)
	}
	return o
}

// Logger retrieves the "logger" object from the main scope.
//
// ---------------------------------------------
//
//	name: "logger"
//	type: *slog.Logger
//	scope: "main"
//	build: func
//	params:
//		- "0": Service(*config.Config) ["config"]
//	unshared: false
//	close: false
//
// ---------------------------------------------
//
// It tries to find the container with the C method and the given interface.
// If the container can be retrieved, it calls the GetLogger method.
// If the container can not be retrieved, it panics.
func Logger(i interface{}) *slog.Logger {
	return C(i).GetLogger()
}

// SafeGetMessenger retrieves the "messenger" object from the app scope.
//
// ---------------------------------------------
//
//	name: "messenger"
//	type: messaging.Messenger
//	scope: "app"
//	build: func
//	params:
//		- "0": Service(*config.Config) ["config"]
//		- "1": Service(*ses.SES) ["ses"]
//		- "2": Service(*in.Bundle) ["i18n"]
//	unshared: false
//	close: false
//
// ---------------------------------------------
//
// If the object can not be retrieved, it returns an error.
func (c *Container) SafeGetMessenger() (messaging.Messenger, error) {
	i, err := c.ctn.SafeGet("messenger")
	if err != nil {
		var eo messaging.Messenger
		return eo, err
	}
	o, ok := i.(messaging.Messenger)
	if !ok {
		return o, errors.New("could get 'messenger' because the object could not be cast to messaging.Messenger")
	}
	return o, nil
}

// GetMessenger retrieves the "messenger" object from the app scope.
//
// ---------------------------------------------
//
//	name: "messenger"
//	type: messaging.Messenger
//	scope: "app"
//	build: func
//	params:
//		- "0": Service(*config.Config) ["config"]
//		- "1": Service(*ses.SES) ["ses"]
//		- "2": Service(*in.Bundle) ["i18n"]
//	unshared: false
//	close: false
//
// ---------------------------------------------
//
// If the object can not be retrieved, it panics.
func (c *Container) GetMessenger() messaging.Messenger {
	o, err := c.SafeGetMessenger()
	if err != nil {
		panic(err)
	}
	return o
}

// UnscopedSafeGetMessenger retrieves the "messenger" object from the app scope.
//
// ---------------------------------------------
//
//	name: "messenger"
//	type: messaging.Messenger
//	scope: "app"
//	build: func
//	params:
//		- "0": Service(*config.Config) ["config"]
//		- "1": Service(*ses.SES) ["ses"]
//		- "2": Service(*in.Bundle) ["i18n"]
//	unshared: false
//	close: false
//
// ---------------------------------------------
//
// This method can be called even if app is a sub-scope of the container.
// If the object can not be retrieved, it returns an error.
func (c *Container) UnscopedSafeGetMessenger() (messaging.Messenger, error) {
	i, err := c.ctn.UnscopedSafeGet("messenger")
	if err != nil {
		var eo messaging.Messenger
		return eo, err
	}
	o, ok := i.(messaging.Messenger)
	if !ok {
		return o, errors.New("could get 'messenger' because the object could not be cast to messaging.Messenger")
	}
	return o, nil
}

// UnscopedGetMessenger retrieves the "messenger" object from the app scope.
//
// ---------------------------------------------
//
//	name: "messenger"
//	type: messaging.Messenger
//	scope: "app"
//	build: func
//	params:
//		- "0": Service(*config.Config) ["config"]
//		- "1": Service(*ses.SES) ["ses"]
//		- "2": Service(*in.Bundle) ["i18n"]
//	unshared: false
//	close: false
//
// ---------------------------------------------
//
// This method can be called even if app is a sub-scope of the container.
// If the object can not be retrieved, it panics.
func (c *Container) UnscopedGetMessenger() messaging.Messenger {
	o, err := c.UnscopedSafeGetMessenger()
	if err != nil {
		panic(err)
	}
	return o
}

// Messenger retrieves the "messenger" object from the app scope.
//
// ---------------------------------------------
//
//	name: "messenger"
//	type: messaging.Messenger
//	scope: "app"
//	build: func
//	params:
//		- "0": Service(*config.Config) ["config"]
//		- "1": Service(*ses.SES) ["ses"]
//		- "2": Service(*in.Bundle) ["i18n"]
//	unshared: false
//	close: false
//
// ---------------------------------------------
//
// It tries to find the container with the C method and the given interface.
// If the container can be retrieved, it calls the GetMessenger method.
// If the container can not be retrieved, it panics.
func Messenger(i interface{}) messaging.Messenger {
	return C(i).GetMessenger()
}

// SafeGetSes retrieves the "ses" object from the app scope.
//
// ---------------------------------------------
//
//	name: "ses"
//	type: *ses.SES
//	scope: "app"
//	build: func
//	params:
//		- "0": Service(*config.Config) ["config"]
//	unshared: false
//	close: false
//
// ---------------------------------------------
//
// If the object can not be retrieved, it returns an error.
func (c *Container) SafeGetSes() (*ses.SES, error) {
	i, err := c.ctn.SafeGet("ses")
	if err != nil {
		var eo *ses.SES
		return eo, err
	}
	o, ok := i.(*ses.SES)
	if !ok {
		return o, errors.New("could get 'ses' because the object could not be cast to *ses.SES")
	}
	return o, nil
}

// GetSes retrieves the "ses" object from the app scope.
//
// ---------------------------------------------
//
//	name: "ses"
//	type: *ses.SES
//	scope: "app"
//	build: func
//	params:
//		- "0": Service(*config.Config) ["config"]
//	unshared: false
//	close: false
//
// ---------------------------------------------
//
// If the object can not be retrieved, it panics.
func (c *Container) GetSes() *ses.SES {
	o, err := c.SafeGetSes()
	if err != nil {
		panic(err)
	}
	return o
}

// UnscopedSafeGetSes retrieves the "ses" object from the app scope.
//
// ---------------------------------------------
//
//	name: "ses"
//	type: *ses.SES
//	scope: "app"
//	build: func
//	params:
//		- "0": Service(*config.Config) ["config"]
//	unshared: false
//	close: false
//
// ---------------------------------------------
//
// This method can be called even if app is a sub-scope of the container.
// If the object can not be retrieved, it returns an error.
func (c *Container) UnscopedSafeGetSes() (*ses.SES, error) {
	i, err := c.ctn.UnscopedSafeGet("ses")
	if err != nil {
		var eo *ses.SES
		return eo, err
	}
	o, ok := i.(*ses.SES)
	if !ok {
		return o, errors.New("could get 'ses' because the object could not be cast to *ses.SES")
	}
	return o, nil
}

// UnscopedGetSes retrieves the "ses" object from the app scope.
//
// ---------------------------------------------
//
//	name: "ses"
//	type: *ses.SES
//	scope: "app"
//	build: func
//	params:
//		- "0": Service(*config.Config) ["config"]
//	unshared: false
//	close: false
//
// ---------------------------------------------
//
// This method can be called even if app is a sub-scope of the container.
// If the object can not be retrieved, it panics.
func (c *Container) UnscopedGetSes() *ses.SES {
	o, err := c.UnscopedSafeGetSes()
	if err != nil {
		panic(err)
	}
	return o
}

// Ses retrieves the "ses" object from the app scope.
//
// ---------------------------------------------
//
//	name: "ses"
//	type: *ses.SES
//	scope: "app"
//	build: func
//	params:
//		- "0": Service(*config.Config) ["config"]
//	unshared: false
//	close: false
//
// ---------------------------------------------
//
// It tries to find the container with the C method and the given interface.
// If the container can be retrieved, it calls the GetSes method.
// If the container can not be retrieved, it panics.
func Ses(i interface{}) *ses.SES {
	return C(i).GetSes()
}
