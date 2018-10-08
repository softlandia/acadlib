golang library get AutoCAD parameters

functions:
1. func IsAutocadInstalled() bool
IsAutocadInstalled Checks installed in system any version of AutoCAD
return true if installed,
return false if not istalled
use const _REG_PATH_ACAD_ == `Software\Autodesk\AutoCAD`

2. func ReleaseAutocadInstalled() string
ReleaseAutocadInstalled return string with release of AutoCAD installer in system
return "-" if Autocad not installed

3. func PathAutocadLogFile() string
PathAutocadLogFile - return string to path of log file