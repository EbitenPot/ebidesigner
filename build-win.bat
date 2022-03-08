del /F /S /Q ebidesigner.exe
del /F /S /Q .\win64\ebidesigner.exe
go build -buildmode=exe
powershell Copy-Item ebidesigner.exe ./win64/
cd win64
ebidesigner.exe

