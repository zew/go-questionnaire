DIRECTORY="build"

if [ -d "$DIRECTORY" ]; then
  echo "starting tests..."
  go test -v  -race  ./...
  sleep 3
  read -p "Press [Enter] to close..."
else
  echo "run this from app root"
fi


