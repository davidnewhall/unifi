// +build !windows,!darwin

package poller

// DefaultConfFile is where to find config if --config is not prvided.
const DefaultConfFile = "/etc/unifi-poller/up.conf"

// DefaultObjPath is the path to look for shared object libraries (plugins).
const DefaultObjPath = "/usr/lib/unifi-poller"
