[ ! -d "dist" ] && mkdir dist
cd dist
[ ! -d "windows" ] && mkdir windows
[ ! -d "linux" ] && mkdir linux
[ ! -d "mac" ] && mkdir mac
cd ../
echo "Set up folders"
GOOS=windows GOARCH=amd64 go build -o dist/windows/backgammon_windows.exe
echo "Windows"
GOOS=linux GOARCH=amd64 go build -o dist/linux/backgammon_linux
echo "Linux"
GOOS=darwin GOARCH=amd64 go build -o dist/mac/backgammon_mac
echo "Mac"