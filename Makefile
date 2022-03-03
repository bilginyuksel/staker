up:
	export ALGOD_URL=http://localhost:8080
	export ALGOD_TOKEN=$(cat /var/lib/algorand/algod.token)
	export KMD_URL=http://localhost:7833
	export KMD_TOKEN=$(cat /var/lib/algorand/kmd.token)
	go run .