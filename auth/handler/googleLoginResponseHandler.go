package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/hellicopthecat/aisha_do_dot/jwt"
	dtos "github.com/hellicopthecat/aisha_do_dot/users/dtos/request"
	"github.com/hellicopthecat/aisha_do_dot/users/entity"
	"github.com/labstack/echo/v5"
)

type GoogleTokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIs    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
	IdToke       string `json:"id_token"`
}

type GoogleUserResponse struct {
	Email string `json:"email"`
	Sub   string `json:"sub"`
	Name  string `json:"name"`
}

func (h AuthHandler) GoogleLoginResponseHandler(e *echo.Context) error {

	stateCookie, err := e.Cookie("state")
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{
			"STATUS":  "FAILED",
			"MESSAGE": "잘못된 요청입니다.",
		})
	}
	stateParam := e.QueryParam("state")
	codeParam := e.QueryParam("code")

	clientId := os.Getenv("GOOGLE_CLIENT_ID")
	secretKey := os.Getenv("GOOGLE_SECRET_KEY")

	e.SetCookie(&http.Cookie{
		Name:     "state",
		Value:    "",
		MaxAge:   -1,
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // 배포 시 true
		SameSite: http.SameSiteLaxMode,
	})

	if stateCookie.Value != stateParam {
		return e.JSON(http.StatusBadRequest, map[string]string{
			"STATUS":  "FAILED",
			"MESSAGE": "잘못된 요청입니다.",
		})
	}

	if codeParam == "" {
		return e.JSON(http.StatusBadRequest, map[string]string{
			"STATUS":  "FAILED",
			"MESSAGE": "잘못된 요청입니다.",
		})
	}

	qp := url.Values{}
	qp.Add("client_id", clientId)
	qp.Add("client_secret", secretKey)
	qp.Add("code", codeParam)
	qp.Add("grant_type", "authorization_code")
	qp.Add("redirect_uri", "http://localhost:8080/api/auth/google/login/response")

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	ctx := e.Request().Context()

	tokenReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		"https://oauth2.googleapis.com/token",
		strings.NewReader(qp.Encode()),
	)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{
			"STATUS":  "FAILED",
			"MESSAGE": "토큰 요청에 실패했습니다.",
		})
	}

	tokenReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	tokenRes, err := client.Do(tokenReq)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{
			"STATUS":  "FAILED",
			"MESSAGE": "Google 서버에 연결하지 못했습니다.",
		})
	}

	defer tokenRes.Body.Close()

	if tokenRes.StatusCode != http.StatusOK {
		return e.JSON(http.StatusInternalServerError, map[string]string{
			"STATUS":  "FAILED",
			"MESSAGE": "Google 토큰 교환에 실패했습니다.",
		})
	}

	var tokenResponse GoogleTokenResponse

	if err := json.NewDecoder(tokenRes.Body).Decode(&tokenResponse); err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{
			"STATUS":  "FAILED",
			"MESSAGE": "Google 토큰 응답을 읽지 못 했습니다.",
		})
	}

	if tokenResponse.AccessToken == "" {
		return e.JSON(http.StatusInternalServerError, map[string]string{
			"STATUS":  "FAILED",
			"MESSAGE": "Google 엑세스 토큰이 없습니다.",
		})
	}

	userReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		"https://openidconnect.googleapis.com/v1/userinfo",
		nil,
	)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]string{
			"STATUS":  "FAILED",
			"MESSAGE": "사용자 정보 요청 생성에 실패했습니다.",
		})
	}

	userReq.Header.Set("Authorization", "Bearer "+tokenResponse.AccessToken)

	userRes, err := client.Do(userReq)

	if err != nil {
		return e.JSON(http.StatusBadGateway, map[string]string{
			"STATUS":  "FAILED",
			"MESSAGE": "Google 사용자 정보 서버에 연결하지 못했습니다..",
		})
	}
	defer userReq.Body.Close()

	if userRes.StatusCode != http.StatusOK {
		return e.JSON(http.StatusBadGateway, map[string]string{
			"STATUS":  "FAILED",
			"MESSAGE": "Google 사용자정보 조회에 실패했습니다.",
		})
	}

	var userResponse GoogleUserResponse
	if err := json.NewDecoder(userRes.Body).Decode(&userResponse); err != nil {
		return e.JSON(http.StatusBadGateway, map[string]string{
			"STATUS":  "FAILED",
			"MESSAGE": "Google 사용자정보를 읽지 못했습니다.",
		})
	}

	if userResponse.Sub == "" {
		return e.JSON(http.StatusBadGateway, map[string]string{
			"STATUS":  "FAILED",
			"MESSAGE": "Google 사용자 ID가 없습니다.",
		})
	}

	// 유저 찾기
	email, err := h.authHandler.FindUserByEmailService(ctx, userResponse.Email)

	// user 없으면 생성 로그인
	// token들 생성
	if err == nil && email == nil {
		uid, _ := uuid.NewUUID()
		rt, err := setCookies(uid, userResponse.Email, e)

		if err != nil {
			return e.JSON(http.StatusUnauthorized, map[string]string{
				"STATUS":  "FAILED",
				"MESSAGE": "토큰 등록에 실패하였습니다.",
			})
		}

		user := &entity.User{
			Id:            uid,
			Email:         userResponse.Email,
			Social:        "GOOGLE",
			Provider_id:   userResponse.Sub,
			Name:          userResponse.Name,
			Refresh_token: *rt,
		}
		h.authHandler.CreateUserSerivce(ctx, user)
	}
	// 있으면 그냥 로그인
	if err == nil && email != nil {
		uid, err := h.authHandler.GetUserIdByEmailService(ctx, *email)

		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]string{
				"STATUS":  "FAILED",
				"MESSAGE": "사용자정보를 읽어오는데 실패했습니다.",
			})
		}
		if uid.String() == "" {
			return e.JSON(http.StatusBadRequest, map[string]string{
				"STATUS":  "FAILED",
				"MESSAGE": "존재하지 않는 사용자입니다..",
			})
		}

		rt, err := setCookies(*uid, userResponse.Email, e)
		if err != nil {
			return e.JSON(http.StatusUnauthorized, map[string]string{
				"STATUS":  "FAILED",
				"MESSAGE": "토큰 등록에 실패하였습니다.",
			})
		}

		dto := dtos.UpdateRefreshToken{
			Id:            *uid,
			Refresh_token: *rt,
		}
		h.authHandler.UpdateRefreshTokenService(ctx, dto)

	}
	return e.Redirect(http.StatusFound, "http://localhost:3000")
}

func setCookies(uid uuid.UUID, email string, e *echo.Context) (*string, error) {
	at, err := jwt.CreateJWTTokens(uid, email, true)
	if err != nil {
		return nil, fmt.Errorf("토큰생성에 실패했습니다. :: %v", err)
	}

	rt, err := jwt.CreateJWTTokens(uid, email, false)
	if err != nil {
		return nil, fmt.Errorf("토큰생성에 실패했습니다. :: %v", err)
	}

	e.SetCookie(&http.Cookie{
		Name:     "access_token",
		Value:    *at,
		Path:     "/",
		MaxAge:   60 * 60 * 24 * 3,
		HttpOnly: true,
		Secure:   false, // 배포시 true
		SameSite: http.SameSiteLaxMode,
	})
	e.SetCookie(&http.Cookie{
		Name:     "refresh_token",
		Value:    *rt,
		Path:     "/",
		MaxAge:   60 * 60 * 24 * 7,
		HttpOnly: true,
		Secure:   false, // 배포시 true
		SameSite: http.SameSiteLaxMode,
	})
	return rt, nil
}
