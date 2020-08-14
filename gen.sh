#!/bin/bash

version=$(cat version.txt)
macsha=$(cat dist/crtools_darwin_amd64.sha256sum | awk '{print $1}')
linuxsha=$(cat dist/crtools_linux_amd64.sha256sum | awk '{print $1}')

cat > crtools.rb <<EOF
class Crtools < Formula
    desc "Devops tools 运维工具"
    homepage "https://github.com/ysicing/crtools"
    version "${version}"
    bottle :unneeded

    if OS.mac?
      url "https://github.com/ysicing/crtools/releases/download/#{version}/crtools_darwin_amd64"
      sha256 "${macsha}"
    elsif OS.linux?
      if Hardware::CPU.intel?
        url "https://github.com/ysicing/crtools/releases/download/#{version}/crtools_linux_amd64"
        sha256 "${linuxsha}"
      end
    end

    def install
      bin.install "crtools_darwin_amd64" => "crtools"
    end
  end
EOF

docker build -t ysicing/taprb:crtools -f hack/brew/Dockerfile .
docker push ysicing/taprb:crtools