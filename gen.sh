#!/bin/bash

set -e

VERSION=1.5.1

if [ ! -d assets/k8s.io/kubernetes ]; then
    mkdir -p assets/k8s.io
    wget https://github.com/kubernetes/kubernetes/archive/v${VERSION}.zip
    unzip v${VERSION}.zip
    mv kubernetes-${VERSION} assets/k8s.io/kubernetes
    rm v${VERSION}.zip
fi

rm -rf api apis runtime util types.go

PKG=$PWD

cd assets

protobuf=$( find k8s.io/kubernetes/pkg/{api,apis,util,runtime} -name '*.proto' )
for file in $protobuf; do
    echo $file
    protoc --gogofast_out=$PKG $file
done

cd -

mv k8s.io/kubernetes/pkg/* .
rm -rf k8s.io
sed -i '' 's|"k8s.io/kubernetes/pkg|"github.com/ericchiang/k8s|g' $(find {api,apis,util,runtime} -name '*.go')
sed -i '' 's|"k8s.io.kubernetes.pkg.|"github.com/ericchiang.k8s.|g' $(find {api,apis,util,runtime} -name '*.go')

rm -rf assets

cd .gobuild/src/github.com/ericchiang/k8s && go run gen.go
cd -

