### golang library get AutoCAD parameters from registry ###

	download: go get -u github.com/softlandia/acadlib

functions:  
1. IsAutocadInstalled() bool
2. ReleaseAutocadInstalled() string
3. RegPathAutocadLogFile() string
4. PathAutocadLogFile() string

_________________________________________________________________________
	func IsAutocadInstalled() bool

IsAutocadInstalled - checks installed in system any version of AutoCAD 
return true if installed,  
return false if not istalled  
used const _RegPathAcad === "Software\Autodesk\AutoCAD"

	func ReleaseAutocadInstalled() string  

ReleaseAutocadInstalled return string with release of AutoCAD installer in system  
return "-" if Autocad not installed

	func RegPathAutocadLogFile() string  

RegPathAutocadLogFile - return string registry path to key store log file folder


	func PathAutocadLogFile() string  

PathAutocadLogFile - return path to folder where AutoCAD store log file
