rm -f govcl-template
rm -f ./linux64/govcl-template
go build
cp govcl-template ./linux64/
cd linux64
govcl-template