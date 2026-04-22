# Windows integrity control

The **Mandatory Integrity Control** security feature on Windows devices allocates an integrity level (IL) to all applications and processes. The device’s operating system or database constrains a user’s or initiator’s ability to access or perform other operations on an *object/target* (such as files, memory, or directories). Both the initiator and the object are allocated an IL, with low as the most restricted access. When the initiator attempts to access the object, their ILs are compared. The initiator can’t access the object if it has a lower IL than the object.

For more information about integrity levels, see Microsoft’s [Mandatory Integrity Control](https://docs.microsoft.com/en-us/windows/win32/secauthz/mandatory-integrity-control) documentation.

Windows Standalone player can detect if it’s running at low integrity level. For more information, see Microsoft’s documentation on [Designing Applications to Run at a Low Integrity Level](https://msdn.microsoft.com/en-us/library/bb625960.aspx). In this case, one of the following things might happen:

* The log file is written to `%USERPROFILE%\AppData\LocalLow\CompanyName\ProductName`
* PlayerPrefs is saved to `HKCU\Software\AppDataLow\Software\CompanyName\ProductName`