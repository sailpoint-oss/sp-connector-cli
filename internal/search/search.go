package search

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/fatih/color"
	"github.com/mitchellh/mapstructure"
	sailpoint "github.com/sailpoint-oss/golang-sdk/sdk-output"
	sailpointsdk "github.com/sailpoint-oss/golang-sdk/sdk-output/v3"
	"github.com/sailpoint-oss/sailpoint-cli/internal/output"
)

func ParseIndicie(indicie string) (sailpointsdk.Index, error) {
	switch indicie {
	case "accessprofiles":
		return sailpointsdk.INDEX_ACCESSPROFILES, nil
	case "accountactivities":
		return sailpointsdk.INDEX_ACCOUNTACTIVITIES, nil
	case "entitlements":
		return sailpointsdk.INDEX_ENTITLEMENTS, nil
	case "events":
		return sailpointsdk.INDEX_EVENTS, nil
	case "identities":
		return sailpointsdk.INDEX_IDENTITIES, nil
	case "roles":
		return sailpointsdk.INDEX_ROLES, nil
	}
	return "*", fmt.Errorf("indicie provided is invalid")
}

func BuildSearch(searchQuery string, sort []string, indicies []string) (sailpointsdk.Search1, error) {

	search := sailpointsdk.NewSearch1()
	search.Query = sailpointsdk.NewQuery()
	search.Query.Query = &searchQuery
	search.Sort = sort
	search.Indices = []sailpointsdk.Index{}

	for i := 0; i < len(indicies); i++ {
		tempIndicie, err := ParseIndicie(indicies[i])

		if err != nil {
			return *search, err
		}

		search.Indices = append(search.Indices, tempIndicie)
	}

	return *search, nil
}

func PerformSearch(apiClient sailpoint.APIClient, search sailpointsdk.Search1) (SearchResults, error) {
	var SearchResults SearchResults

	ctx := context.TODO()
	resp, r, err := sailpoint.PaginateWithDefaults[map[string]interface{}](apiClient.V3.SearchApi.SearchPost(ctx).Search1(search))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	color.Green("Search complete, saving results")

	for i := 0; i < len(resp); i++ {
		entry := resp[i]
		switch entry["_type"] {
		case "accountactivity":
			var AccountActivity AccountActivity
			err := mapstructure.Decode(entry, &AccountActivity)
			if err != nil {
				return SearchResults, err
			}
			SearchResults.AccountActivities = append(SearchResults.AccountActivities, AccountActivity)

		case "accessprofile":
			var AccessProfile AccessProfile
			err := mapstructure.Decode(entry, &AccessProfile)
			if err != nil {
				return SearchResults, err
			}
			SearchResults.AccessProfiles = append(SearchResults.AccessProfiles, AccessProfile)

		case "entitlement":
			var Entitlement Entitlement
			err := mapstructure.Decode(entry, &Entitlement)
			if err != nil {
				return SearchResults, err
			}
			SearchResults.Entitlements = append(SearchResults.Entitlements, Entitlement)

		case "event":
			var Event Event
			err := mapstructure.Decode(entry, &Event)
			if err != nil {
				return SearchResults, err
			}
			SearchResults.Events = append(SearchResults.Events, Event)

		case "identity":
			var Identity Identity
			err := mapstructure.Decode(entry, &Identity)
			if err != nil {
				return SearchResults, err
			}
			SearchResults.Identities = append(SearchResults.Identities, Identity)

		case "role":
			var Role Role
			err := mapstructure.Decode(entry, &Role)
			if err != nil {
				return SearchResults, err
			}
			SearchResults.Roles = append(SearchResults.Roles, Role)
		}
	}

	return SearchResults, nil
}

func IterateIndicies(SearchResults SearchResults, searchQuery string, folderPath string, outputTypes []string) error {
	var err error
	if len(SearchResults.AccountActivities) > 0 {
		fileName := fmt.Sprintf("query=%s&indicie=%s", searchQuery, "AccountActivities")
		err = SaveResults(SearchResults.AccountActivities, fileName, folderPath, outputTypes)
		if err != nil {
			return err
		}
	}
	if len(SearchResults.AccessProfiles) > 0 {
		fileName := fmt.Sprintf("query=%s&indicie=%s", searchQuery, "AccessProfiles")
		err = SaveResults(SearchResults.AccessProfiles, fileName, folderPath, outputTypes)
		if err != nil {
			return err
		}
	}
	if len(SearchResults.Entitlements) > 0 {
		fileName := fmt.Sprintf("query=%s&indicie=%s", searchQuery, "Entitlements")
		err = SaveResults(SearchResults.Entitlements, fileName, folderPath, outputTypes)
		if err != nil {
			return err
		}
	}
	if len(SearchResults.Events) > 0 {
		fileName := fmt.Sprintf("query=%s&indicie=%s", searchQuery, "Events")
		err = SaveResults(SearchResults.Events, fileName, folderPath, outputTypes)
		if err != nil {
			return err
		}
	}
	if len(SearchResults.Identities) > 0 {
		fileName := fmt.Sprintf("query=%s&indicie=%s", searchQuery, "Identities")
		err = SaveResults(SearchResults.Identities, fileName, folderPath, outputTypes)
		if err != nil {
			return err
		}
	}
	if len(SearchResults.Roles) > 0 {
		fileName := fmt.Sprintf("query=%s&indicie=%s", searchQuery, "Roles")
		err = SaveResults(SearchResults.Roles, fileName, folderPath, outputTypes)
		if err != nil {
			return err
		}
	}
	return nil
}

func SaveResults[T any](formattedResponse []T, fileName string, filePath string, outputTypes []string) error {
	for i := 0; i < len(outputTypes); i++ {
		outputType := outputTypes[i]
		switch outputType {
		case "json":
			fileName = fileName + ".json"
			savePath := path.Join(filePath, fileName)
			err := output.SaveJSONFile(formattedResponse, fileName, filePath)
			if err != nil {
				return err
			}
			color.Green("Saving file: %s", savePath)
		case "csv":
			fileName = fileName + ".csv"
			savePath := path.Join(filePath, fileName)
			err := output.SaveCSVFile(formattedResponse, fileName, filePath)
			if err != nil {
				return err
			}
			color.Green("Saving file: %s", savePath)
		default:
			return fmt.Errorf("invalid output type provided %s", outputType)
		}
	}

	return nil
}
