// Package Openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.1 DO NOT EDIT.
package Openapi

const (
	BearerScopes = "Bearer.Scopes"
)

// Error defines model for Error.
type Error struct {
	Message string `json:"message"`
}

// Meh defines model for Meh.
type Meh struct {

	// ユーザーID
	Id *uint64 `json:"id,omitempty"`

	// 投稿したテキスト
	Text string `json:"text"`
	User *User  `json:"user,omitempty"`
}

// Pagination defines model for Pagination.
type Pagination struct {

	// 1ページの最大数
	Count uint64 `json:"count"`

	// ページの最後の要素のID
	LastId *uint64 `json:"last_id,omitempty"`
}

// TimelineResponse defines model for TimelineResponse.
type TimelineResponse struct {
	Mehs       []Meh      `json:"mehs"`
	Pagination Pagination `json:"pagination"`
}

// User defines model for User.
type User struct {

	// ユーザーID
	Id uint64 `json:"id"`

	// 名前
	ScreenName string `json:"screen_name"`
}

// FollowUserJSONBody defines parameters for FollowUser.
type FollowUserJSONBody struct {

	// フォローされるユーザーのID
	FolloweeId uint64 `json:"followee_id"`

	// フォローするユーザーのID
	UserId uint64 `json:"user_id"`
}

// RemoveUserJSONBody defines parameters for RemoveUser.
type RemoveUserJSONBody struct {

	// フォローされているユーザーのID
	FolloweeId uint64 `json:"followee_id"`

	// フォローしているユーザーのID
	UserId uint64 `json:"user_id"`
}

// ShowTimelineJSONBody defines parameters for ShowTimeline.
type ShowTimelineJSONBody struct {
	Pagination Pagination `json:"pagination"`
}

// CreateMehJSONBody defines parameters for CreateMeh.
type CreateMehJSONBody struct {

	// テキスト
	Text string `json:"text"`

	// 投稿するユーザーのID
	UserId uint64 `json:"user_id"`
}

// CreateUserJSONBody defines parameters for CreateUser.
type CreateUserJSONBody struct {

	// スクリーンネーム
	ScreenName string `json:"screen_name"`
}

// FollowUserJSONRequestBody defines body for FollowUser for application/json ContentType.
type FollowUserJSONRequestBody FollowUserJSONBody

// RemoveUserJSONRequestBody defines body for RemoveUser for application/json ContentType.
type RemoveUserJSONRequestBody RemoveUserJSONBody

// ShowTimelineJSONRequestBody defines body for ShowTimeline for application/json ContentType.
type ShowTimelineJSONRequestBody ShowTimelineJSONBody

// CreateMehJSONRequestBody defines body for CreateMeh for application/json ContentType.
type CreateMehJSONRequestBody CreateMehJSONBody

// CreateUserJSONRequestBody defines body for CreateUser for application/json ContentType.
type CreateUserJSONRequestBody CreateUserJSONBody