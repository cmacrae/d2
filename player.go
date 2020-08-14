package d2

import (
	"encoding/json"
	"fmt"
)

type PlayerService struct {
	client *Client
}

// GetDestinyManifest Returns the current version of the manifest.
func (p *PlayerService) GetDestinyManifest() (Manifest, error) {
	u := fmt.Sprintf("Destiny2/Manifest/")

	retData := Manifest{}
	response, err := p.client.Platform.PlatformRequest("GET", u)
	if err != nil {
		return retData, err
	}
	err = json.Unmarshal(response, &retData)
	if err != nil {
		return retData, err
	}
	return retData, nil
}

// SearchDestinyPlayer Returns a list of Destiny memberships given a player name.
func (p *PlayerService) SearchDestinyPlayer(membershipType string, playerName string) (ByPlayerNameResponse, error) {
	pfmType := BungieMembershipType[membershipType]
	u := fmt.Sprintf("Destiny2/SearchDestinyPlayer/%d/%s/", pfmType, playerName)

	retData := ByPlayerNameResponse{}
	response, err := p.client.Platform.PlatformRequest("GET", u)
	if err != nil {
		return retData, err
	}
	err = json.Unmarshal(response, &retData)
	if err != nil {
		return retData, err
	}
	return retData, nil
}

// GetProfile Returns Destiny profile compnent information for the supplied membership & component type.
func (p *PlayerService) GetProfile(membershipType, membershipID, componentType string, i interface{}) error {
	pfmType := BungieMembershipType[membershipType]
	u := fmt.Sprintf("Destiny2/%d/Profile/%s?components=%s", pfmType, membershipID, componentType)

	response, err := p.client.Platform.PlatformRequest("GET", u)
	if err != nil {
		return err

	}

	if err := json.Unmarshal(response, i); err != nil {
		return err
	}

	return nil
}

func (p *PlayerService) GetProfileCharacters(membershipType, membershipID string) (GetProfileCharactersResponse, error) {
	retData := &GetProfileCharactersResponse{}
	if err := p.GetProfile(membershipType, membershipID, "Characters", retData); err != nil {
		return *retData, err
	}
	return *retData, nil
}

func (p *PlayerService) GetPublicMilestoneContent(milestoneHash string) (MilestoneContent, error) {
	u := fmt.Sprintf("Destiny2/Milestones/%s/Content/", milestoneHash)
	retData := MilestoneContent{}

	response, err := p.client.Platform.PlatformRequest("GET", u)
	if err != nil {
		return retData, err
	}
	err = json.Unmarshal(response, &retData)
	if err != nil {
		return retData, err
	}
	return retData, nil
}

//GetPublicMilestones Gets public information about currently available Milestones.
func (p *PlayerService) GetPublicMilestones() (Milestones, error) {
	u := fmt.Sprintf("Destiny2/Milestones")

	var incomingData map[string]*json.RawMessage
	var milestone Milestone
	var milestonesList Milestones

	response, err := p.client.Platform.PlatformRequest("GET", u)
	if err != nil {
		return milestonesList, err
	}
	err = json.Unmarshal(response, &incomingData)
	if err != nil {
		return milestonesList, err
	}
	if len(incomingData) > 0 {
		for _, v := range incomingData {
			err = json.Unmarshal(*v, &milestone)
			if err != nil {
				return milestonesList, err
			}
			milestonesList = append(milestonesList, milestone)
		}
	}
	return milestonesList, nil
}

//GetClanWeeklyRewardState Returns information on the weekly clan rewards and if the clan has earned them or not.
//Note that this will always report rewards as not redeemed.
func (p *PlayerService) GetClanWeeklyRewardState(clanID int) (ClanWeeklyRewardState, error) {
	u := fmt.Sprintf("Destiny2/Clan/%d/WeeklyRewardState/", clanID)

	retData := ClanWeeklyRewardState{}
	response, err := p.client.Platform.PlatformRequest("GET", u)
	if err != nil {
		return retData, err
	}
	err = json.Unmarshal(response, &retData)
	if err != nil {
		return retData, err
	}
	return retData, nil
}

//GetHistoricalStatsDefinition Returns historical stats definitions.
func (p *PlayerService) GetHistoricalStatsDefinition() (HistoricalStatsDefinition, error) {
	u := fmt.Sprintf("Destiny2/Stats/Definition/")

	retData := HistoricalStatsDefinition{}
	response, err := p.client.Platform.PlatformRequest("GET", u)
	if err != nil {
		return retData, err
	}
	err = json.Unmarshal(response, &retData)
	if err != nil {
		return retData, err
	}
	return retData, nil
}

//GetHistoricalStatsForAccount Gets aggregate historical stats organized around each character for a given account.
//PREVIEW: This endpoint is still in beta, and may experience rough edges.
//The schema is in final form, but there may be bugs that prevent desirable operation.
func (p *PlayerService) GetHistoricalStatsForAccount(membershipType string, membershipID string) (CommonAccountStats, error) {
	pfmType := BungieMembershipType[membershipType]
	u := fmt.Sprintf("Destiny2/%d/Account/%s/Stats/", pfmType, membershipID)

	retData := CommonAccountStats{}
	response, err := p.client.Platform.PlatformRequest("GET", u)
	if err != nil {
		return retData, err
	}
	err = json.Unmarshal(response, &retData)
	if err != nil {
		return retData, err
	}
	return retData, nil
}
