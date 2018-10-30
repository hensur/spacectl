package teams

import "github.com/mittwald/spacectl/client/lowlevel"

type TeamsClient interface {
	List() ([]Team, error)
	Get(teamIDOrName string) (*Team, error)
	Delete(teamIDOrName string) error
	Create(string, string) (Team, error)
	InviteByEmail(teamID string, email string, message string, role string) (Invite, error)
	InviteByUID(teamID string, uid string, message string, role string) (Invite, error)
	ListMembers(teamID string) ([]Membership, error)
	ListRoles(teamID string) ([]TeamRole, error)
}

func NewTeamsClient(client *lowlevel.SpacesLowlevelClient) TeamsClient {
	return &teamsClient{client}
}

type teamsClient struct {
	client *lowlevel.SpacesLowlevelClient
}
