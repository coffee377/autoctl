#!/bin/sh

set -e

usage() {
	echo "Usage: $0 <--dynamic|--static> [--system]">&2
	exit 1
}

if [ "$#" -eq "0" ]; then
	usage
fi

# 导入环境变量
. "$(dirname $0)/env.sh"

BUILD_SYSTEM=OFF
BUILD_SHARED_LIBS=

while [ $# -gt 0 ]; do
	case "$1" in
		--static)
			BUILD_PATH="${ROOT}/static-build"
			BUILD_SHARED_LIBS=OFF
			;;

		--dynamic)
			BUILD_PATH="${ROOT}/dynamic-build"
			BUILD_SHARED_LIBS=ON
			;;

		--system)
			BUILD_SYSTEM=ON
			;;

		*)
			usage
			;;
	esac
	shift
done

if [ -z "${BUILD_SHARED_LIBS}" ]; then
	usage
fi


if [ ! -d $LIBGIT2_PATH ]; then
  echo 'code not exist'
#  mkdir -p $LIBGIT2_PATH
#  git submodule update --init --recursive
#  trap "git submodule update --init" EXIT
  exit 0
fi

if [ -n "${BUILD_LIBGIT_REF}" ]; then
	git -C "${LIBGIT2_PATH}" checkout "${BUILD_LIBGIT_REF}"
	trap "git submodule update --init" EXIT
fi

BUILD_DEPRECATED_HARD="ON"
if [ "${BUILD_SYSTEM}" = "ON" ]; then
	BUILD_INSTALL_PREFIX=${SYSTEM_INSTALL_PREFIX-"/usr"}
	# Most system-wide installations won't intentionally omit deprecated symbols.
	BUILD_DEPRECATED_HARD="OFF"
else
	BUILD_INSTALL_PREFIX="${BUILD_PATH}/install"
	mkdir -p "${BUILD_PATH}/install/lib"
fi

USE_BUNDLED_ZLIB="ON"
if [ "${USE_CHROMIUM_ZLIB}" = "ON" ]; then
	USE_BUNDLED_ZLIB="Chromium"
fi

mkdir -p "${BUILD_PATH}/build"
cd "${BUILD_PATH}/build"
cmake -DTHREADSAFE=ON \
      -DBUILD_CLAR=OFF \
      -DBUILD_SHARED_LIBS"=${BUILD_SHARED_LIBS}" \
      -DREGEX_BACKEND=builtin \
      -DUSE_BUNDLED_ZLIB="${USE_BUNDLED_ZLIB}" \
      -DUSE_HTTPS=OFF \
      -DUSE_SSH=OFF \
      -DCURL=OFF \
      -DCMAKE_C_FLAGS=-fPIC \
      -DCMAKE_BUILD_TYPE="RelWithDebInfo" \
      -DCMAKE_INSTALL_PREFIX="${BUILD_INSTALL_PREFIX}" \
      -DCMAKE_INSTALL_LIBDIR="lib" \
      -DDEPRECATE_HARD="${BUILD_DEPRECATE_HARD}" \
      "${LIBGIT2_PATH}"

if which make nproc >/dev/null && [ -f Makefile ]; then
	# Make the build parallel if make is available and cmake used Makefiles.
	exec make "-j$(nproc --all)" install
else
	exec cmake --build . --target install
fi