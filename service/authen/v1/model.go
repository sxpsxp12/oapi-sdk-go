package v1

type AuthenAccessTokenReqBody struct {
	GrantType string `json:"grant_type,omitempty"`
	Code      string `json:"code,omitempty"`
}

type AuthenRefreshAccessTokenReqBody struct {
	GrantType    string `json:"grant_type,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

type UserAccessTokenInfo struct {
	AccessToken      string `json:"access_token,omitempty"`
	AvatarUrl        string `json:"avatar_url,omitempty"`
	AvatarThumb      string `json:"avatar_thumb,omitempty"`
	AvatarMiddle     string `json:"avatar_middle,omitempty"`
	AvatarBig        string `json:"avatar_big,omitempty"`
	ExpiresIn        int    `json:"expires_in,omitempty"`
	Name             string `json:"name,omitempty"`
	EnName           string `json:"en_name,omitempty"`
	OpenId           string `json:"open_id,omitempty"`
	UnionId          string `json:"union_id,omitempty"`
	UserId           string `json:"user_id,omitempty"`
	TenantKey        string `json:"tenant_key,omitempty"`
	RefreshExpiresIn int    `json:"refresh_expires_in,omitempty"`
	RefreshToken     string `json:"refresh_token,omitempty"`
	TokenType        string `json:"token_type,omitempty"`
}

type UserInfo struct {
	Name         string `json:"name,omitempty"`
	AvatarUrl    string `json:"avatar_url,omitempty"`
	AvatarThumb  string `json:"avatar_thumb,omitempty"`
	AvatarMiddle string `json:"avatar_middle,omitempty"`
	AvatarBig    string `json:"avatar_big,omitempty"`
	Email        string `json:"email,omitempty"`
	OpenId       string `json:"open_id,omitempty"`
	UnionId      string `json:"union_id,omitempty"`
	UserId       string `json:"user_id,omitempty"`
	Mobile       string `json:"mobile,omitempty"`
}
