#!/bin/sh

#
# Use this utility to import libgit2 sources into the directory.
#

set -e

. ./env.sh

if [ -d $LIBGIT2_PATH ]; then
  rm -rf $LIBGIT2_PATH
fi

git clone https://github.com/libgit2/libgit2.git $LIBGIT2_PATH

cd $LIBGIT2_PATH
git checkout $LIBGIT2_VERSION

cd $ROOT
rm -rf $LIBGIT2_PATH/.git