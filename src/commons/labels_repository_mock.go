// +build dev

package commons

import (
	"github.com/gofrs/uuid"
	datastructures "github.com/bbernhard/imagemonkey-core/datastructures"
	"errors"
)

type LabelsRepository struct {
	gitCheckoutDir            string
}

func NewLabelsRepository(projectOwner string, repositoryName string, gitCheckoutDir string) *LabelsRepository {
	return &LabelsRepository{
		gitCheckoutDir:            gitCheckoutDir,
	}
}

func (p *LabelsRepository) SetToken(token string) {
}


func (p *LabelsRepository) Clone() error {
	return nil
}


func (p *LabelsRepository) AddLabelAndPushToRepo(trendingLabel datastructures.TrendingLabelBotTask) (string, error) {
	branchNameUuid, err := uuid.NewV4()
	if err != nil {
		return "", errors.New("Couldn't create branch name: " + err.Error())
	}

	if trendingLabel.LabelType == "normal" {
		labelEntry, err := generateLabelEntry(trendingLabel.RenameTo, trendingLabel.Plural, trendingLabel.Description)
		if err != nil {
			return "", errors.New("Couldn't generate label entry: " + err.Error())
		}

		autoGeneratedLabelsWriter := NewAutoLabelsWriter(p.gitCheckoutDir + "/en/includes/labels/autogenerated/" + labelEntry.Uuid + ".json")
		err = autoGeneratedLabelsWriter.Add(trendingLabel.Name, labelEntry)
		if err != nil {
			return "", errors.New("Couldn't add label: " + err.Error())
		}
	} else if trendingLabel.LabelType == "meta" {
		metaLabelEntry, err := generateMetaLabelEntry(trendingLabel.RenameTo, trendingLabel.Plural, trendingLabel.Description)
		if err != nil {
			return "", errors.New("Couldn't generate label entry: " + err.Error())
		}

		autoGeneratedMetaLabelsWriter := NewMetaLabelsWriter(p.gitCheckoutDir + "/en/includes/metalabels/autogenerated/" + metaLabelEntry.Uuid + ".json")
		err = autoGeneratedMetaLabelsWriter.Add(trendingLabel.Name, metaLabelEntry)
		if err != nil {
			return "", errors.New("Couldn't add label: " + err.Error())
		}
	} else {
		return "", errors.New("Invalid label type " + trendingLabel.LabelType)
	}

	return branchNameUuid.String(), nil
}

func (p *LabelsRepository) MergeRemoteBranchIntoMaster(branchName string) error {
	return nil
}

func (p *LabelsRepository) RemoveRemoteBranch(branchName string) error {
	return nil
}

func (p *LabelsRepository) RemoveLocal() error {
	return nil
}

