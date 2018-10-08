package acadlib

import (
	"golang.org/x/sys/windows/registry"
)

const (
	_RegPathAcad              = `Software\Autodesk\AutoCAD`
	_RegPathEditorConfAcadR15 = `Software\Autodesk\AutoCAD\R15.0\ACAD-201:409\Profiles\<<Unnamed Profile>>\Editor Configuration`
	_RegPathEditorConfAcadR16 = `Software\Autodesk\AutoCAD\R16.0\ACAD-201:409\Profiles\<<Unnamed Profile>>\Editor Configuration`
	_RegPathEditorConfAcadR19 = `Software\Autodesk\AutoCAD\R19.0\ACAD-B001:409\Profiles\<<Unnamed Profile>>\Editor Configuration`
	_RegPathEditorConfAcadR22 = `Software\Autodesk\AutoCAD\R22.0\ACAD-2001:409\Profiles\<<Unnamed Profile>>\Editor Configuration`
	_RegPathEditorConfAcadR23 = `Software\Autodesk\AutoCAD\R23.0\ACAD-2001:409\Profiles\<<Unnamed Profile>>\Editor Configuration`
)

//TAcadVersion - struct for store version
type TAcadVersion struct {
	ver  string
	name string
	path string
}

var _AcadVersion = [...]TAcadVersion{
	{"-", "-", "-"},
	{"AutoCAD 2000", "15.0", _RegPathEditorConfAcadR15},
	{"AutoCAD 2013", "19.0", _RegPathEditorConfAcadR19},
	{"AutoCAD 2018", "22.0", _RegPathEditorConfAcadR22},
	{"AutoCAD 2019", "23.0", _RegPathEditorConfAcadR23},
}

//IsAutocadInstalled Checks installed in system any version of AutoCAD
// return true if installed,
// return false if not istalled
// use const _REG_PATH_ACAD_ == `Software\Autodesk\AutoCAD`
func IsAutocadInstalled() bool {
	_, err := registry.OpenKey(registry.CURRENT_USER, _RegPathAcad, registry.QUERY_VALUE)
	return (err == nil)
}

//ReleaseAutocadInstalled return string with release of AutoCAD installer in system
//return "-" if Autocad not installed
func ReleaseAutocadInstalled() string {
	key, err := registry.OpenKey(registry.CURRENT_USER, _RegPathAcad, registry.QUERY_VALUE)
	if err != nil {
		return "-"
	}
	defer key.Close()

	s, _, err := key.GetStringValue("CurVer")
	if err != nil {
		return "-"
	}
	return s
}

//PathAutocadLogFile - return string to path of log file
func PathAutocadLogFile() string {
	if !IsAutocadInstalled() {
		return "-"
	}

	path := _RegPathAcad
	key, err := registry.OpenKey(registry.CURRENT_USER, path, registry.QUERY_VALUE)
	if err != nil {
		return "-"
	}
	defer key.Close()

	s, _, err := key.GetStringValue("CurVer")
	if err != nil {
		return "-"
	}

	path = path + "\\" + s
	key, err = registry.OpenKey(registry.CURRENT_USER, path, registry.QUERY_VALUE)
	if err != nil {
		return "-"
	}

	s, _, err = key.GetStringValue("CurVer")
	if err != nil {
		return "-"
	}

	path = path + "\\" + s
	key, err = registry.OpenKey(registry.CURRENT_USER, path, registry.QUERY_VALUE)
	if err != nil {
		return "-"
	}

	return path + "\\Profiles\\<<Unnamed Profile>>\\Editor Configuration"
}
