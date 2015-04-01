// Copyright 2015 Canonical Ltd.
// Copyright 2015 Cloudbase Solutions SRL
// Licensed under the AGPLv3, see LICENCE file for details.

package commands

const (
	// AptConfFilePath is the full file path for the proxy settings that are
	// written by cloud-init and the machine environ worker.
	AptConfFilePath = "/etc/apt/apt.conf.d/42-juju-proxy-settings"

	// the basic command for all dpkg calls:
	dpkg = "dpkg"

	// the basic command for all apt-get calls:
	//		--assume-yes to never prompt for confirmation
	//		--force-confold is passed to dpkg to never overwrite config files
	aptget = "apt-get --assume-yes --option Dpkg::Options::=--force-confold"

	// the basic command for all apt-cache calls:
	aptcache = "apt-cache"

	// the basic command for all add-apt-repository calls:
	//		--yes to never prompt for confirmation
	addaptrepo = "add-apt-repository --yes"

	// the basic command for all apt-config calls:
	aptconfig = "apt-config dump"

	// the basic format for specifying a proxy option for apt:
	aptProxySettingFormat = "Acquire::%s::Proxy %q;"
)

// aptCmder is the packageCommander instantiation for apt-based systems.
var aptCmder = packageCommander{
	prereq:              buildCommand(aptget, "install python-software-properties"),
	update:              buildCommand(aptget, "update"),
	upgrade:             buildCommand(aptget, "upgrade"),
	install:             buildCommand(aptget, "install"),
	remove:              buildCommand(aptget, "remove"),
	purge:               buildCommand(aptget, "purge"),
	search:              buildCommand(aptcache, "search --names-only ^%s$"),
	isInstalled:         buildCommand(dpkg, "-s %s"),
	listAvailable:       buildCommand(aptcache, "pkgnames"),
	listInstalled:       buildCommand(dpkg, "--get-selections"),
	addRepository:       buildCommand(addaptrepo, "%s"),
	listRepositories:    buildCommand(`sed -r -n "s|^deb(-src)? (.*)|\2|p"`, "/etc/apt/sources.list"),
	removeRepository:    buildCommand(addaptrepo, "--remove ppa:%s"),
	cleanup:             buildCommand(aptget, "autoremove"),
	getProxy:            buildCommand(aptconfig, "Acquire::http::Proxy Acquire::https::Proxy Acquire::ftp::Proxy"),
	proxySettingsFormat: aptProxySettingFormat,
	setProxy:            buildCommand("echo %s >> ", AptConfFilePath),
}
