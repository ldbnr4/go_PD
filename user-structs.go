package main

import "gopkg.in/mgo.v2/bson"

// User ...
// TODO: replace Joined and Nickname with UserProfile
type User struct {
	bson.ObjectId                       "_id"
	Username                            string
	Email                               string
	Password                            string
	Nickname                            string
	Joined                              string
	Friends, Albums, FriendReqs, Tagged []bson.ObjectId
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

// DelUserMsg ...
type DelUserMsg struct {
	ID       string
	Password string
}

// LoginMsg ...
type LoginMsg struct {
	Username string
	Password string
}

// GetFriendsResponse ...
type GetFriendsResponse struct {
	FriendReqs []SimpleUser
	Friends    []SimpleUser
}

// SimpleUser ...
type SimpleUser struct {
	Name string
	ID   string
}

// UserProfile ...
type UserProfile struct {
	Nickname string
	Joined   string
}
