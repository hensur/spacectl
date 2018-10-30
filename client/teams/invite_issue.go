package teams

import (
	"errors"
	"fmt"
)

type inviteRequestInvitee struct {
	ID    string `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
}

type inviteRequest struct {
	Invitee inviteRequestInvitee `json:"invitee"`
	Message string               `json:"message"`
	Role    string               `json:"role,omitempty"`
}

func (t *teamsClient) InviteByEmail(teamID string, email string, message string, role string) (Invite, error) {
	var invite Invite

	if email == "" {
		return invite, errors.New("email must not be empty")
	}

	if role == "" {
		return invite, errors.New("role must be set")
	}

	req := inviteRequest{
		Invitee: inviteRequestInvitee{
			Email: email,
		},
		Message: message,
		Role:    role,
	}

	url := fmt.Sprintf("/teams/%s/invites", teamID)
	err := t.client.Post(url, &req, &invite)
	if err != nil {
		return invite, fmt.Errorf("could not invite user: %s", err)
	}

	return invite, nil
}

func (t *teamsClient) InviteByUID(teamID string, uid string, message string, role string) (Invite, error) {
	var invite Invite

	if uid == "" {
		return invite, errors.New("UID must not be empty")
	}

	if role == "" {
		return invite, errors.New("role must be set")
	}

	req := inviteRequest{
		Invitee: inviteRequestInvitee{
			ID: uid,
		},
		Message: message,
		Role:    role,
	}

	url := fmt.Sprintf("/teams/%s/invites", teamID)
	err := t.client.Post(url, &req, &invite)
	if err != nil {
		return invite, fmt.Errorf("could not invite user: %s", err)
	}

	return invite, nil
}
