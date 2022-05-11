# ------------------------------------------------------------------------------
# go project
# ------------------------------------------------------------------------------

GO      = go build
LD      = go build -o
EX      =
TEST    = go test
FMT     = go fmt

# ------------------------------------------------------------------------------
# configuraiton
# ------------------------------------------------------------------------------
REPO      = alpha
LIBS      = github.com/xiejw/${REPO}/nn/...
CMDS      = github.com/xiejw/${REPO}/cmd/...
BUILD_DIR = .build

# ------------------------------------------------------------------------------
# actions
# ------------------------------------------------------------------------------

.PHONY: compile

compile: compile_libs ${BUILD_DIR}/alpha

compile_libs:
	${GO} ${LIBS}

fmt:
	${FMT} ${LIBS}
	${FMT} ${CMDS}

test:
	${TEST} ${LIBS}

clean:
	go mod tidy

# ------------------------------------------------------------------------------
# binaries
# ------------------------------------------------------------------------------
${BUILD_DIR}:
	@mkdir -p ${BUILD_DIR}

.PHONY: a alpha ${BUILD_DIR}/alpha

a: alpha

alpha: ${BUILD_DIR}/alpha
	@${BUILD_DIR}/alpha

${BUILD_DIR}/alpha: ${BUILD_DIR}
	${LD} $@ cmd/alpha/main.go

