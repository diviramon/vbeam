set -e
go build -o ./env/pinpub/ github.com/diviramon/vbeam/pinpub
# ln -s ./pinpub/pinpub.json ./env/pinpub/pinpub.json
cd env/pinpub/ && ./pinpub