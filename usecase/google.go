package usecase

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/SSSBoOm/SE_PROJECT_BACKEND/domain"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type googleUsecase struct {
	cfg *domain.ConfigEnv
}

func NewGoogleUsecase(cfg *domain.ConfigEnv) domain.GoogleUsecase {
	return &googleUsecase{cfg: cfg}
}

func (u *googleUsecase) GoogleConfig() *oauth2.Config {
	config := &oauth2.Config{
		ClientID:     u.cfg.GOOGLE_CLIENT_ID,
		ClientSecret: u.cfg.GOOGLE_CLIENT_SECRET,
		RedirectURL:  u.cfg.GOOGLE_REDIRECT,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}

	return config
}

func (u *googleUsecase) GetToken(c *fiber.Ctx) (*oauth2.Token, error) {
	token, err := u.GoogleConfig().Exchange(c.Context(), c.FormValue("code"))
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (u *googleUsecase) GetProfile(token string) (*domain.GoogleResponse, error) {
	reqUrl, err := url.Parse("https://www.googleapis.com/oauth2/v1/userinfo")
	if err != nil {
		panic(err)
	}
	ptoken := fmt.Sprintf("Bearer %s", token)

	res := &http.Request{
		Method: "GET",
		URL:    reqUrl,
		Header: map[string][]string{"Authorization": {ptoken}},
	}

	req, err := http.DefaultClient.Do(res)
	if err != nil {
		panic(err)
	}

	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	var data domain.GoogleResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}
	return &data, nil
}
