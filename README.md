golang library get AutoCAD parameters from registry

functions:
1. 
func IsAutocadInstalled() bool
//IsAutocadInstalled - checks installed in system any version of AutoCAD
// return true if installed,
// return false if not istalled
// used const _RegPathAcad === `Software\Autodesk\AutoCAD`

2. 
func ReleaseAutocadInstalled() string
//ReleaseAutocadInstalled return string with release of AutoCAD installer in system
//return "-" if Autocad not installed

3. 
func RegPathAutocadLogFile() string
//RegPathAutocadLogFile - return string registry path to key store log file folder

4.
func PathAutocadLogFile() string
//PathAutocadLogFile - return path to folder where AutoCAD store log file

SAMPLE:
