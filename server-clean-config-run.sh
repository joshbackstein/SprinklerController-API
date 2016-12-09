if [ -f ./config.json.bak ]; then
  echo "File \"config.json.bak\" was found. Copying it to \"config.json\" file."
  echo ""
  cp config.json.bak config.json && go run *.go
else
  echo "File \"config.json.bak\" not found. Creating it from \"config.json\" file."
  echo ""
  cp config.json config.json.bak && go run *.go
fi
