# CIT 90
## Week 1
### Course Outlines
https://goo.gl/K5GccY
https://goo.gl/

#### Things to consider when developing in go.
Concurrency vs parallelism

### Go Workspace
Go is dependent on having a 'go workspace'. You need three folders: 'bin', 'pkg', and 'src'.

### Terminal Commands
__pwd__: Print working Directory
__mkdir__: Make a folder at the current path with given name
__cd__: change directory to specified name, or up a folder

### Hash Function
Maps any arbitrary size of data to a fixed size. Regardless of how much data is provided. The same data will always result in the same hash value. If any binary bit of data should change somewhre in the file, it should result in a different hash.

#### Computing a hash on Windows
To compute a hash on windows, we can use a PowerShell Commandlet
__Get-FileHash__: 
  * __-Path__: the path to the file hash to compute
  * __-Algorithm__: the hashing algorithm to use when computing the hash. Default is SHA256.


### Environment Variables
__Using Powershell__: 
  * Get-Childiten env:
  * dir env:

### HW:
Variables, values & type
Exercises - Ninja Level 1