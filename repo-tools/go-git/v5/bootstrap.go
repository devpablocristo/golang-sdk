package sdkgogit

import (
	ports "github.com/devpablocristo/golang-sdk/repo-tools/go-git/v5/ports"
)

func Bootstrap(repoRemoteUrl, repoLocalPath, repoBranch string) (ports.Client, error) {
	// config := newConfig(
	// 	viper.GetString("GIT_REPO_URL"),
	// 	viper.GetString("GIT_REPO_PATH"),
	// 	viper.GetString("GIT_REPO_BRANCH"),
	// )

	config := newConfig(repoRemoteUrl, repoLocalPath, repoBranch)

	if err := config.Validate(); err != nil {
		return nil, err
	}

	return newClient(config)
}
