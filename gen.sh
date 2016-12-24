#!/bin/bash

set -ex

VERSION=1.4.6

if [ ! -d assets/k8s.io/kubernetes ]; then
    mkdir -p assets/k8s.io
    wget https://github.com/kubernetes/kubernetes/archive/v${VERSION}.zip
    unzip v${VERSION}.zip
    mv kubernetes-${VERSION} assets/k8s.io/kubernetes
    rm v1.4.6.zip
fi

rm -rf api apis config.go runtime util types.go

PKG=$PWD

cd assets

protobuf=$( find k8s.io/kubernetes/pkg/{api,apis,util,runtime} -name '*.proto' )
for file in $protobuf; do
    echo $file
    protoc --gofast_out=$PKG $file
done

cd -

mv k8s.io/kubernetes/pkg/* .
rm -rf k8s.io

client_dir="client/unversioned/clientcmd/api/v1"
cp assets/k8s.io/kubernetes/pkg/${client_dir}/types.go config.go
sed -i 's|package v1|package k8s|g' config.go

sed -i 's|"k8s.io/kubernetes/pkg|"github.com/ericchiang/k8s|g' $( find {api,apis,config.go,util,runtime} -name '*.go' )
sed -i 's|"k8s.io.kubernetes.pkg.|"github.com/ericchiang.k8s.|g' $( find {api,apis,config.go,util,runtime} -name '*.go' )

rm -rf assets

go run gen.go

cat << EOF >> api/unversioned/time.go
package unversioned

import (
    "encoding/json"
    "time"
)

func (t Time) MarshalJSON() ([]byte, error) {
    var seconds, nanos int64
    if t.Seconds != nil {
        seconds = *t.Seconds
    }
    if t.Nanos != nil {
        nanos = int64(*t.Nanos)
    }
    return json.Marshal(time.Unix(seconds, nanos))
}

func (t *Time) UnmarshalJSON(p []byte) error {
    var t1 time.Time
    if err := json.Unmarshal(p, &t1); err != nil {
        return err
    }
    seconds := t1.Unix()
    nanos := int32(t1.UnixNano())
    t.Seconds = &seconds
    t.Nanos = &nanos
    return nil
}
EOF
gofmt -w api/unversioned/time.go
