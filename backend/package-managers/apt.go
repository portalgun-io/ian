// Copyright © 2016 Theotime LEVEQUE theotime@protonmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package packagemanagers

import (
	"errors"
	"fmt"
	"os"
	"runtime"

	"github.com/thylong/ian/backend/command"
)

// Apt immutable instance.
var Apt = AptPackageManager{Path: "/usr/bin/apt-get", Name: "apt"}

// ErrAptMissingFeature is returned when triggering an unsupported feature.
var ErrAptMissingFeature = errors.New("apt is not designed to support this feature")

// AptPackageManager is the official Debian (and associated distributions) package manager.
// (more: https://wiki.debian.org/Apt)
type AptPackageManager struct {
	Path string
	Name string
}

// Install given Apt package.
func (b AptPackageManager) Install(packageName string) (err error) {
	err = command.ExecuteCommand(execCommand(b.Path, "install", packageName))
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
	}
	return err
}

// Uninstall given Apt package.
func (b AptPackageManager) Uninstall(packageName string) (err error) {
	err = command.ExecuteCommand(execCommand(b.Path, "remove", packageName))
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
	}
	return err
}

// Cleanup all the local archives and previous versions.
func (b AptPackageManager) Cleanup() (err error) {
	err = command.ExecuteCommand(execCommand(b.Path, "autoremove"))
	return err
}

// UpdateOne pulls last versions infos from related repositories.
// This is not performing any updates and should be coupled
// with upgradeAll command.
func (b AptPackageManager) UpdateOne(packageName string) (err error) {
	return ErrAptMissingFeature
}

// UpgradeOne Npm packages to the last known versions.
func (b AptPackageManager) UpgradeOne(packageName string) error {
	return command.ExecuteCommand(execCommand(b.Path, "upgrade", packageName))
}

// UpdateAll pulls last versions infos from realted repositories.
// This is not performing any updates and should be coupled
// with upgradeAll command.
func (b AptPackageManager) UpdateAll() error {
	return command.ExecuteCommand(execCommand(b.Path, "update"))
}

// UpgradeAll Apt packages to the last known versions.
func (b AptPackageManager) UpgradeAll() error {
	return command.ExecuteCommand(execCommand(b.Path, "full-upgrade"))
}

// IsInstalled returns true if Apt executable is found.
func (b AptPackageManager) IsInstalled() bool {
	if _, err := os.Stat(b.Path); err != nil {
		return false
	}
	return true
}

// IsOSPackageManager returns true for Mac OS.
func (b AptPackageManager) IsOSPackageManager() bool {
	return b.IsInstalled() && runtime.GOOS == "linux`"
}

// GetExecPath return immutable path to Apt executable.
func (b AptPackageManager) GetExecPath() string {
	return b.Path
}

// GetName return the name of the package manager.
func (b AptPackageManager) GetName() string {
	return b.Name
}

// Setup does nothing (apt comes by default in Linux distributions)
func (b AptPackageManager) Setup() (err error) {
	return nil
}
