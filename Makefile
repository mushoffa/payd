#SHELL:=/bin/sh.
REGEX_1=^v(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?$$
REGEX_2=^([0-9]{1,}\.){2}[0-9]{1,}$$
REGEX_SEMVER=[^0-9.]*\([0-9.]*\).*/\1/

STDV=standard-version
STDV_PREVIEW=${STDV} --dry-run --release-as
STDV_RELEASE=${STDV} --release-as
PREVIEW=standard-version --dry-run --release-as
GET_VERSION=sed -n '/release/ s/${REGEX_SEMVER}p'
GIT_BRANCH=$$(git branch --show-current)
VERSION_FILE=confir/version.go
VERSION_PACKAGE=config
VERSION=v0.4.1

format:
	gofmt -s -w .

report:
	goreportcard-cli

preview-%: VERSION=$$(echo "v$$(${STDV_PREVIEW} $* | ${GET_VERSION})")
preview-%:
	${STDV_PREVIEW} $*

release-%:
	VERSION=$$(echo "v$$(${STDV_PREVIEW} $* | ${GET_VERSION})"); \
	git checkout -b "release/$${VERSION}"; \
	git commit -am "release: $${VERSION}"; \
	git push -u origin ${GIT_BRANCH}; \
	${STDV_RELEASE} $*; \
	git push --follow-tags origin ${GIT_BRANCH}