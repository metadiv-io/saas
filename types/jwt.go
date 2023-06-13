package types

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/metadiv-io/saas/constant"
)

type Jwt struct {
	UUID       string   `json:"uuid"`
	UserUUID   string   `json:"user_uuid"`
	Username   string   `json:"username"`
	Type       string   `json:"type"`       // admin, user, workspace_user, api_key
	UserAgent  string   `json:"user_agent"` // * means all
	IPs        []string `json:"ips"`        // * means all
	Workspaces []string `json:"workspaces"` // admin has no workspace; workspace_user/api_key has only one workspace; user has many workspaces
}

func (j *Jwt) IssueAdminToken(d time.Duration, privPEM, userUUID, username, ip, userAgent string) (string, error) {
	j.UUID = uuid.NewString()
	j.UserUUID = userUUID
	j.Username = username
	j.Type = constant.JWT_TYPE_ADMIN
	j.UserAgent = userAgent
	j.IPs = []string{ip}
	return j.issueToken(d, privPEM)
}

func (j *Jwt) IssueUserToken(d time.Duration, privPEM, userUUID, username, ip, userAgent string, workspaces []string) (string, error) {
	j.UUID = uuid.NewString()
	j.UserUUID = userUUID
	j.Username = username
	j.Type = constant.JWT_TYPE_USER
	j.UserAgent = userAgent
	j.IPs = []string{ip}
	j.Workspaces = workspaces
	return j.issueToken(d, privPEM)
}

func (j *Jwt) IssueWorkspaceUserToken(d time.Duration, privPEM, userUUID, username, ip, userAgent, workspaceUUID string) (string, error) {
	j.UUID = uuid.NewString()
	j.UserUUID = userUUID
	j.Username = username
	j.Type = constant.JWT_TYPE_WORKSPACE_USER
	j.UserAgent = userAgent
	j.IPs = []string{ip}
	j.Workspaces = []string{workspaceUUID}
	return j.issueToken(d, privPEM)
}

func (j *Jwt) IssueAPIKeyToken(d time.Duration, privPEM, userUUID, username string, ips []string, workspaceUUID string) (string, error) {
	j.UUID = uuid.NewString()
	j.UserUUID = userUUID
	j.Username = username
	j.Type = constant.JWT_TYPE_API_KEY
	j.IPs = ips
	j.Workspaces = []string{workspaceUUID}
	return j.issueToken(d, privPEM)
}

func (j *Jwt) ParseToken(pubPEM, token string) error {
	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(pubPEM))
	if err != nil {
		return fmt.Errorf("jwt error - parsing public pem: %w", err)
	}
	claims := make(jwt.MapClaims)
	_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return fmt.Errorf("jwt error - parsing token: %w", err)
	}
	j.UUID = claims["uuid"].(string)
	j.UserUUID = claims["user_uuid"].(string)
	j.Username = claims["username"].(string)
	j.Type = claims["type"].(string)
	j.UserAgent = claims["user_agent"].(string)
	j.IPs = make([]string, 0)
	for _, v := range claims["ips"].([]interface{}) {
		j.IPs = append(j.IPs, v.(string))
	}
	j.Workspaces = make([]string, 0)
	for _, v := range claims["workspaces"].([]interface{}) {
		j.Workspaces = append(j.Workspaces, v.(string))
	}
	return nil
}

func (j *Jwt) IsAdmin() bool {
	return j.Type == "admin"
}

func (j *Jwt) IsUser() bool {
	return j.Type == "user"
}

func (j *Jwt) IsWorkspaceUser() bool {
	return j.Type == "workspace_user"
}

func (j *Jwt) IsAPIKey() bool {
	return j.Type == "api_key"
}

func (j *Jwt) IsIPAllowed(ip string) bool {
	if j.IPs == nil {
		return false
	}
	for _, v := range j.IPs {
		if v == "*" || v == ip {
			return true
		}
	}
	return false
}

func (j *Jwt) IsUserAgentAllowed(userAgent string) bool {
	if j.UserAgent == "*" {
		return true
	}
	return j.UserAgent == userAgent
}

func (j *Jwt) IsWorkspaceAllowed(workspaceUUID string) bool {
	if j.Workspaces == nil {
		return false
	}
	for _, v := range j.Workspaces {
		if v == workspaceUUID {
			return true
		}
	}
	return false
}

func (j *Jwt) issueToken(d time.Duration, privPEM string) (string, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privPEM))
	if err != nil {
		return "", fmt.Errorf("jwt error - parsing private pem: %w", err)
	}
	now := time.Now()
	claims := make(jwt.MapClaims)
	claims["uuid"] = j.UUID
	claims["user_uuid"] = j.UserUUID
	claims["username"] = j.Username
	claims["type"] = j.Type
	claims["ips"] = j.IPs
	claims["user_agent"] = j.UserAgent
	claims["workspaces"] = j.Workspaces
	claims["iat"] = now.Unix()
	claims["exp"] = now.Add(d).Unix()

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		return "", fmt.Errorf("jwt error - signing token: %w", err)
	}
	return token, nil
}
