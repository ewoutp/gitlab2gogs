package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ewoutp/go-gitlab-client"
	"github.com/gogits/go-gogs-client"
)

var (
	gitlabHost     string
	gitlabApiPath  string
	gitlabUser     string
	gitlabPassword string
	gitlabToken    string
	gogsUrl        string
	gogsToken      string
	gogsUser       string
)

func init() {
	flag.StringVar(&gitlabHost, "gitlab-host", "", "")
	flag.StringVar(&gitlabApiPath, "gitlab-api-path", "", "")
	flag.StringVar(&gitlabUser, "gitlab-user", "", "")
	flag.StringVar(&gitlabPassword, "gitlab-password", "", "")
	flag.StringVar(&gitlabToken, "gitlab-token", "", "")
	flag.StringVar(&gogsUrl, "gogs-url", "", "")
	flag.StringVar(&gogsToken, "gogs-token", "", "")
	flag.StringVar(&gogsUser, "gogs-user", "", "")
}

func main() {
	flag.Parse()

	gc := gogs.NewClient(gogsUrl, gogsToken)
	orgMap := make(map[string]*gogs.Organization)

	getOrg := func(name string) *gogs.Organization {
		org, ok := orgMap[name]
		if ok {
			return org
		}
		org, err := gc.GetOrg(name)
		if err == nil {
			orgMap[name] = org
			return org
		}
		createOpt := gogs.CreateOrgOption{
			UserName: name,
		}
		org, err = gc.AdminCreateOrg(gogsUser, createOpt)
		if err != nil {
			exitf("Failed to create organization '%s': %v\n", name, err)
		}
		orgMap[name] = org
		return org
	}

	migrate := func(p *gogitlab.Project) {
		_, err := gc.GetRepo(p.Namespace.Name, p.Name)
		if err == nil {
			fmt.Printf("%s | %s already exists\n", p.Namespace.Name, p.Name)
		} else {
			org := getOrg(p.Namespace.Name)
			name := fixName(p.Name)
			fmt.Printf("%s | %s migrating as '%s'...\n", p.Namespace.Name, p.Name, name)
			opts := gogs.MigrateRepoOption{
				CloneAddr:    p.HttpRepoUrl,
				AuthUsername: gitlabUser,
				AuthPassword: gitlabPassword,
				UID:          int(org.ID),
				RepoName:     name,
				Private:      !p.Public,
				Description:  p.Description,
			}
			_, err := gc.MigrateRepo(opts)
			if err != nil {
				exitf("Failed to migrate '%s | %s': %v\n", p.Namespace.Name, p.Name, err)
			}
		}
	}

	gitlab := gogitlab.NewGitlab(gitlabHost, gitlabApiPath, gitlabToken)
	projects, err := gitlab.AllProjects()
	if err != nil {
		exitf("Cannot get gitlab projects: %v\n", err)
	}
	for _, p := range projects {
		if p.Archived {
			continue
		}
		migrate(p)
	}
}

func fixName(name string) string {
	switch name {
	case "api": // reserved
		return "theapi"
	default:
		return name
	}
}

func exitf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
	os.Exit(1)
}
