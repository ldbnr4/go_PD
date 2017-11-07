package main

import "gopkg.in/mgo.v2/bson"

// ServerUser ...
type ServerUser struct {
	bson.ObjectId "_id"
	Username      string
	Email         string
	Password      string
	UserProfile
	Friends, Albums, FriendReqs, Tagged []bson.ObjectId
}

//ClientUser
type ClientUser struct {
	ObjectID bson.ObjectId
	Email    string
	UserProfile
	GetFriendsResponse
	GetAlbumsResp
}

// CreateUserError ...
// TODO: convert these to strings
type CreateUserError struct {
	Username bool
	Email    bool
}

// AddUserMsg ...
type AddUserMsg struct {
	Username string
	Password string
	Email    string
	Nickname string
}

// AddUserResp ...
type AddUserResp struct {
	Error CreateUserError
	ID    string
}

// LoginMsg ...
type LoginMsg struct {
	Username string
	Password string
}

// GetFriendsResponse ...
type GetFriendsResponse struct {
	FriendReqs []string
	Friends    []string
}

// UserProfile ...
type UserProfile struct {
	Nickname string
	Joined   string
}
