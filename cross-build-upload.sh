printenv

ARCH=(linux/amd64 linux/arm darwin/amd64 windows/amd64 windows/arm)
TAG=${1##*/}
echo "${TAG} being generated"
for i in "${ARCH[@]}"
do
	arrIN=(${i//\// })
	echo ${arrIN[@]}
	env GOOS="${arrIN[0]}" GOARCH="${arrIN[1]}" go build -o mfa-auth-${arrIN[0]}-${arrIN[1]}
	gh release upload ${TAG} mfa-auth-${arrIN[0]}-${arrIN[1]}
done