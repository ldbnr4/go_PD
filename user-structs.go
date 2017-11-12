package main

import "gopkg.in/mgo.v2/bson"

// ServerUser ...
type ServerUser struct {
	bson.ObjectId                       "_id"
	Username                            string "username"
	Email                               string "email"
	Password                            string "password"
	UserProfile                         "userprofile"
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
	Error      CreateUserError
	ClientUser ClientUser
}

// LoginMsg ...
type LoginMsg struct {
	Username string
	Password string
}

// GetFriendsResponse ...
type GetFriendsResponse struct {
	FriendReqs []UserProfile
	Friends    []UserProfile
}

// UserProfile ...
type UserProfile struct {
	Nickname string
	Joined   string
}
