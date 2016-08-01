require 'formula'

class Runss < Formula
  VERSION = '0.1.3'

  homepage 'https://github.com/winebarrel/runss'
  url "https://github.com/winebarrel/runss/releases/download/v#{VERSION}/runss-v#{VERSION}-darwin-amd64.gz"
  sha256 '7687a68b775cbca81b3f5ba4917bb70d4279eb494d5b2ae3a9f6da7be05bad82'
  version VERSION
  head 'https://github.com/winebarrel/runss.git', :branch => 'master'

  def install
    system "mv runss-v#{VERSION}-darwin-amd64 runss"
    bin.install 'runss'
  end
end
