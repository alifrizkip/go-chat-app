package main

import (
	"errors"
)

// ErrNoAvatarURL is the error that is returned when the
// Avatar instance is unable to provide an avatar URL.
var ErrNoAvatarURL = errors.New("chat: Unable to get an avatar URL")

// Avatar represents types capable of representing
// user profile pictures.
type Avatar interface {
	// GetAvatarURL gets the avatar URL for the specified client,
	// or returns an error if something goes wrong.
	// ErrNoAvatarURL is returned if the object is unable to get
	// a URL for the specified client.
	GetAvatarURL(c *client) (string, error)
}

// AuthAvatar struct
type AuthAvatar struct{}

// UseAuthAvatar type is AuthAvatar
var UseAuthAvatar AuthAvatar

// GetAvatarURL function in AuthAvatar
func (AuthAvatar) GetAvatarURL(c *client) (string, error) {
	if url, ok := c.userData["avatar_url"]; ok {
		if urlStr, ok := url.(string); ok {
			return urlStr, nil
		}
	}
	return "", ErrNoAvatarURL
}

// GravatarAvatar struct
type GravatarAvatar struct{}

// UseGravatar types is GravatarAvatar
var UseGravatar GravatarAvatar

// GetAvatarURL function in GravatarAvatar
func (GravatarAvatar) GetAvatarURL(c *client) (string, error) {
	if userID, ok := c.userData["userID"]; ok {
		if useridStr, ok := userID.(string); ok {
			return "//www.gravatar.com/avatar/" + useridStr, nil
		}
	}
	return "", ErrNoAvatarURL
}
