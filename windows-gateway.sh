set -e
cd gateway/ && go build -o ../env/gateway/ && cd ..
# ln -s ./pinpub/pinpub.json ./env/pinpub/pinpub.json
cd env/gateway/ && ./gateway