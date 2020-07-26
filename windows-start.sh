set -e
cd pinpub/ && go build -o ../env/pinpub/ && cd ..
# ln -s ./pinpub/pinpub.json ./env/pinpub/pinpub.json
cd env/pinpub/ && ./pinpub