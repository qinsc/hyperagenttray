DIM objShell   
set objShell = wscript.createObject("wscript.shell")   
currPath = wscript.createobject("Scripting.FileSystemObject").GetFile(Wscript.ScriptFullName).ParentFolder.Path
exePath = "cmd.exe /C" & chr(34) & currPath & "\hyperagenttray.exe" & chr(34)
iReturn = objShell.Run(exePath, 0, TRUE)   
set  objShell = Nothing