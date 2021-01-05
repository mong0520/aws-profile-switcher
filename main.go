package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/go-ini/ini"
	"gopkg.in/AlecAivazis/survey.v1"
)

type AwsProfile struct {
	ProfileName string
	Region      string
}

var gAwsProfiles []AwsProfile

func readAwsProfiles() (awsProfiles []AwsProfile, err error) {
	usr, _ := user.Current()
	dir := usr.HomeDir
	path := "~/.aws/credentials"
	path = filepath.Join(dir, path[2:])
	cfg, err := ini.Load(path)
	if err != nil {
		return awsProfiles, err
	}
	for _, sess := range cfg.Sections() {
		if sess.Name() == "DEFAULT" {
			continue
		}
		tmpAwsProfile := AwsProfile{
			ProfileName: sess.Name(),
			Region:      sess.Key("region").String(),
		}
		awsProfiles = append(awsProfiles, tmpAwsProfile)

	}
	return awsProfiles, nil
}

func parseEnv(selectedAwsProfileName string) string {
	if idx := strings.Index(selectedAwsProfileName, "@"); idx > 0 {
		selectedAwsProfileName = selectedAwsProfileName[idx+1 : len(selectedAwsProfileName)]
	}
	return selectedAwsProfileName
}

func writeAwsProfile(selectedAwsProfileName string) error {
	usr, _ := user.Current()
	dir := usr.HomeDir
	path := "~/.aws_exports"
	path = filepath.Join(dir, path[2:])
	f, err := os.Create(path)

	if err != nil {
		return err
	}

	defer f.Close()

	for _, profile := range gAwsProfiles {
		if profile.ProfileName == selectedAwsProfileName {
			_, err := f.WriteString(fmt.Sprintf("export AWS_PROFILE=%s\n", selectedAwsProfileName))
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func main() {
	awsProfiles, err := readAwsProfiles()
	gAwsProfiles = awsProfiles
	profileNames := []string{}
	for _, profile := range awsProfiles {
		profileNames = append(profileNames, profile.ProfileName)
	}
	qs := []*survey.Question{
		{
			Name: "profile",
			Prompt: &survey.Select{
				Message: "Choose AWS Profile:",
				Options: profileNames,
			},
		},
	}
	answers := struct {
		AwsProfile string `survey:"profile"`
	}{}

	//perform the questions
	fmt.Println("AWS profile switcher")
	err = survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Switch to %s.\n", answers.AwsProfile)
	err = writeAwsProfile(answers.AwsProfile)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Done")
	}
}
