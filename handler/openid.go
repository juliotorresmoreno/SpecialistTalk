package handler

import (
	"bytes"
	"context"
	"crypto/sha256"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/juliotorresmoreno/SpecialistTalk/configs"
	"github.com/juliotorresmoreno/SpecialistTalk/helper"
	"github.com/juliotorresmoreno/SpecialistTalk/openid/storage"
	"github.com/labstack/echo/v4"
	"github.com/zitadel/oidc/pkg/op"
	"golang.org/x/text/language"
)

type OpenIDHandler struct {
	authenticate authenticate
	callback     func(string) string
}

func AttachOpenID(e *echo.Echo) {
	// this will allow us to use an issuer with http:// instead of https://
	os.Setenv(op.OidcDevMode, "true")

	conf := configs.GetConfig()
	ctx := context.Background()
	baseUrl := conf.BaseUrl

	openIDKey := conf.OpenIDKey
	key := sha256.Sum256([]byte(openIDKey))
	storage := storage.NewStorage(storage.NewUserStore())

	provider, err := newOP(ctx, storage, baseUrl, key)
	if err != nil {
		log.Fatal(err)
	}

	l := OpenIDHandler{
		authenticate: storage,
		callback:     op.AuthCallbackURL(provider),
	}

	g := e.Group("login")
	g.GET("/username", l.loginHandler)
	g.POST("/username", l.checkLoginHandler)

	providerHandler := provider.HttpHandler()
	handlerFn := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		providerHandler.ServeHTTP(w, r)
	})
	handler := echo.WrapMiddleware(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w2 := helper.NewRegisterHTTPResponseWriter(w)
			handlerFn(w2, r)

			if w2.StatusCode != http.StatusNotFound {
				w2.Push()
			} else {
				h.ServeHTTP(w, r)
			}

		})
	})
	e.Use(handler)
}

const (
	pathLoggedOut = "/logged-out"
)

// newOP will create an OpenID Provider for localhost on a specified port with a given encryption key
// and a predefined default logout uri
// it will enable all options (see descriptions)
func newOP(ctx context.Context, storage op.Storage, issuer string, key [32]byte) (op.OpenIDProvider, error) {
	config := &op.Config{
		Issuer:    issuer,
		CryptoKey: key,

		// will be used if the end_session endpoint is called without a post_logout_redirect_uri
		DefaultLogoutRedirectURI: pathLoggedOut,

		// enables code_challenge_method S256 for PKCE (and therefore PKCE in general)
		CodeMethodS256: true,

		// enables additional client_id/client_secret authentication by form post (not only HTTP Basic Auth)
		AuthMethodPost: true,

		// enables additional authentication by using private_key_jwt
		AuthMethodPrivateKeyJWT: true,

		// enables refresh_token grant use
		GrantTypeRefreshToken: true,

		// enables use of the `request` Object parameter
		RequestObjectSupported: true,

		// this example has only static texts (in English), so we'll set the here accordingly
		SupportedUILocales: []language.Tag{language.English},
	}
	handler, err := op.NewOpenIDProvider(ctx, config, storage,
		// as an example on how to customize an endpoint this will change the authorization_endpoint from /authorize to /auth
		op.WithCustomAuthEndpoint(op.NewEndpoint("auth")),
	)
	if err != nil {
		return nil, err
	}
	return handler, nil
}

const (
	queryAuthRequestID = "authRequestID"
)

var loginTmpl, _ = template.New("login").Parse(`
	<!DOCTYPE html>
	<html>
		<head>
			<meta charset="UTF-8">
			<title>Login</title>
		</head>
		<body style="display: flex; align-items: center; justify-content: center; height: 100vh;">
			<form method="POST" action="/openid/username" style="height: 200px; width: 200px;">

				<input type="hidden" name="id" value="{{.ID}}">

				<div>
					<label for="username">Username:</label>
					<input id="username" name="username" style="width: 100%">
				</div>

				<div>
					<label for="password">Password:</label>
					<input id="password" name="password" style="width: 100%">
				</div>

				<p style="color:red; min-height: 1rem;">{{.Error}}</p>

				<button type="submit">Login</button>
			</form>
		</body>
	</html>`)

type authenticate interface {
	CheckUsernamePassword(username, password, id string) error
}

func (l *OpenIDHandler) openidConfiguration(c echo.Context) error {
	return c.String(200, "hola mundo")
}

func (l *OpenIDHandler) loginHandler(c echo.Context) error {
	r := c.Request()
	w := c.Response()
	err := r.ParseForm()
	if err != nil {
		return helper.MakeHTTPError(
			http.StatusInternalServerError,
			fmt.Sprintf("cannot parse form:%s", err),
		)
	}

	renderLogin(w, r.FormValue(queryAuthRequestID), nil)
	return nil
}

func renderLogin(w io.Writer, id string, err error) error {
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}
	data := &struct {
		ID    string
		Error string
	}{
		ID:    id,
		Error: errMsg,
	}
	err = loginTmpl.Execute(w, data)
	if err != nil {
		return err
	}
	return nil
}

type checkLoginPayload struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (l *OpenIDHandler) checkLoginHandler(c echo.Context) error {
	w := c.Response()
	payload := &checkLoginPayload{}
	err := c.Bind(payload)
	if err != nil {
		return helper.MakeHTTPError(
			http.StatusInternalServerError,
			fmt.Sprintf("cannot parse form:%s", err),
		)
	}
	username := payload.Username
	password := payload.Password
	id := payload.ID
	err = l.authenticate.CheckUsernamePassword(username, password, id)

	if err != nil {
		buff := bytes.NewBufferString("")
		err = renderLogin(buff, id, err)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return nil
		}
		return c.HTMLBlob(200, buff.Bytes())
	}

	return c.Redirect(http.StatusFound, l.callback(id))
}
