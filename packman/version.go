// Copyright (c) 2023, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package packman

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/Masterminds/semver/v3"
	"goki.dev/goki/config"
	"goki.dev/grease"
	"goki.dev/xe"
)

// GetVersion prints the version of the project.
//
//gti:add
func GetVersion(c *config.Config) error {
	fmt.Println(c.Version)
	return nil
}

// SetVersion updates the config and version file of the config project based
// on the config version and commits and pushes the changes.
//
//gti:add
func SetVersion(c *config.Config) error {
	// we need to update the config file with the new version
	// TODO: determine correct config file instead of just first one
	err := grease.Save(c, grease.ConfigFiles[0])
	if err != nil {
		return fmt.Errorf("error saving new version to config file: %w", err)
	}
	str, err := VersionFileString(c)
	if err != nil {
		return fmt.Errorf("error generating version file string: %w", err)
	}
	err = os.WriteFile(c.Release.VersionFile, []byte(str), 0666)
	if err != nil {
		return fmt.Errorf("error writing version string to version file: %w", err)
	}
	err = PushVersionFileGit(c)
	if err != nil {
		return fmt.Errorf("error pushing version file to Git: %w", err)
	}
	return nil
}

// UpdateVersion updates the version of the project by one patch version.
//
//gti:add
func UpdateVersion(c *config.Config) error {
	ver, err := semver.NewVersion(c.Version)
	if err != nil {
		return fmt.Errorf("error getting semver version from version %q: %w", c.Version, err)
	}

	if !strings.HasPrefix(ver.Prerelease(), "dev") { // if no dev pre-release, we can just do standard increment
		*ver = ver.IncPatch()
	} else { // otherwise, we have to increment pre-release version instead
		pvn := strings.TrimPrefix(ver.Prerelease(), "dev")
		pver, err := semver.NewVersion(pvn)
		if err != nil {
			return fmt.Errorf("error parsing dev version %q from version %q: %w", pvn, c.Version, err)
		}
		*pver = pver.IncPatch()
		// apply incremented pre-release version to main version
		nv, err := ver.SetPrerelease("dev" + pver.String())
		if err != nil {
			return fmt.Errorf("error setting pre-release of new version to %q from repository version %q: %w", "dev"+pver.String(), c.Version, err)
		}
		*ver = nv
	}

	c.Version = "v" + ver.String()
	return SetVersion(c) // now we can set to newly calculated version
}

// VersionFileString returns the version file string
// for a project with the given config info.
func VersionFileString(c *config.Config) (string, error) {
	var b strings.Builder
	b.WriteString("// Code generated by \"goki version\"; DO NOT EDIT.\n\n")
	b.WriteString("package " + c.Release.Package + "\n\n")
	b.WriteString("const (\n")
	b.WriteString("\t// Version is the version of this package being used\n")
	b.WriteString("\tVersion = \"" + c.Version + "\"\n")

	gc := exec.Command("git", "rev-parse", "--short", "HEAD")
	res, err := gc.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error getting previous git commit: %w (%s)", err, res)
	}
	b.WriteString("\t// GitCommit is the commit just before the latest version commit\n")
	b.WriteString("\tGitCommit = \"" + strings.TrimSuffix(string(res), "\n") + "\"\n")

	date := time.Now().UTC().Format("2006-01-02 15:04")
	b.WriteString("\t// VersionDate is the date-time of the latest version commit in UTC (in the format 'YYYY-MM-DD HH:MM', which is the Go format '2006-01-02 15:04')\n")
	b.WriteString("\tVersionDate = \"" + date + "\"\n")
	b.WriteString(")\n")
	return b.String(), nil
}

// PushVersionFileGit makes and pushes a Git commit
// updating the version file based on the given
// config info. It does not actually update the
// version file; it only commits and pushes the
// changes that should have already been made by
// [UpdateVersion].
func PushVersionFileGit(c *config.Config) error {
	vc := xe.VerboseConfig()
	err := xe.Run(vc, "git", "add", c.Release.VersionFile)
	if err != nil {
		return fmt.Errorf("error adding version file: %w", err)
	}
	err = xe.Run(vc, "git", "commit", "-am", "updated version to "+c.Version)
	if err != nil {
		return fmt.Errorf("error commiting release commit: %w", err)
	}
	err = xe.Run(vc, "git", "push")
	if err != nil {
		return fmt.Errorf("error pushing commit: %w", err)
	}
	return nil
}
