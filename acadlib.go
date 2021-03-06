package acadlib

import (
	"golang.org/x/sys/windows/registry"
)

const (
	_RegPathAcad       = `Software\Autodesk\AutoCAD`
	_RegPathAcadEditor = `\Profiles\<<Unnamed Profile>>\Editor Configuration`
	_RegLogFilePath    = "LogFilePath"
)

//IsAutocadInstalled - checks installed in system any version of AutoCAD
// return true if installed,
// return false if not istalled
// used const _RegPathAcad === `Software\Autodesk\AutoCAD`
func IsAutocadInstalled() bool {
	_, err := registry.OpenKey(registry.CURRENT_USER, _RegPathAcad, registry.QUERY_VALUE)
	return (err == nil)
}

//ReleaseAutocadInstalled - return string with release of AutoCAD installer in system
//return "" if Autocad not installed
func ReleaseAutocadInstalled() string {
	key, err := registry.OpenKey(registry.CURRENT_USER, _RegPathAcad, registry.QUERY_VALUE)
	defer key.Close()
	if err != nil {
		return ""
	}

	s, _, err := key.GetStringValue("CurVer")
	if err != nil {
		return ""
	}
	return s
}

//RegPathAutocadLogFile - return string registry path to key store log file folder
//on error return ""
func RegPathAutocadLogFile() string {
	if !IsAutocadInstalled() {
		return ""
	}

	s := ReleaseAutocadInstalled()
	if s == "" {
		return ""
	}
	path := _RegPathAcad + "\\" + s // s like "Software\Autodesk\AutoCAD\R19.0"
	key, err := registry.OpenKey(registry.CURRENT_USER, path, registry.QUERY_VALUE)
	if err != nil {
		return ""
	}

	s, _, err = key.GetStringValue("CurVer") // s like "ACAD-B001:409"
	if err != nil {
		return ""
	}

	path = path + "\\" + s // s like "Software\Autodesk\AutoCAD\R19.0\ACAD-B001:409\Profiles\<<Unnamed Profile>>\Editor Configuration"
	key, err = registry.OpenKey(registry.CURRENT_USER, path, registry.QUERY_VALUE)
	if err != nil {
		return ""
	}

	return path + _RegPathAcadEditor // s like "Software\Autodesk\AutoCAD\R19.0\ACAD-B001:409\"
}

//PathAutocadLogFile - return path to folder where AutoCAD store log file
//return "" on error
func PathAutocadLogFile() string {
	s := RegPathAutocadLogFile()
	if s == "" {
		return ""
	}
	key, err := registry.OpenKey(registry.CURRENT_USER, s, registry.QUERY_VALUE)
	if err != nil {
		return ""
	}

	s, _, err = key.GetStringValue(_RegLogFilePath)
	if err != nil {
		return ""
	}
	return s
}
