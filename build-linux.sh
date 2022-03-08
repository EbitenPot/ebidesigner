rm -f ebidesigner
rm -f ./linux64/ebidesigner
go build
cp ebidesigner ./linux64/
cd linux64
ebidesigner