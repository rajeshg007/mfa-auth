ARCH=(linux/amd64 linux/arm darwin/amd64 darwin/arm64 windows/amd64 windows/arm)
for i in "${ARCH[@]}"
do
	arrIN=(${i//\// })
	echo ${arrIN[@]}
	env GOOS="${arrIN[0]}" GOARCH="${arrIN[1]}" go build -o mfa-auth-${arrIN[0]}-${arrIN[1]}
done
