require 'formula'

class Runss < Formula
  VERSION = '0.1.5'

  homepage 'https://github.com/winebarrel/runss'
  url "https://github.com/winebarrel/runss/releases/download/v#{VERSION}/runss-v#{VERSION}-darwin-amd64.gz"
  sha256 '3934981a72ff1dd20be6cd034c3860c9c0fd733ccbdc3ea755ceb08e52724c59'
  version VERSION
  head 'https://github.com/winebarrel/runss.git', :branch => 'master'

  def install
    system "mv runss-v#{VERSION}-darwin-amd64 runss"
    bin.install 'runss'
  end
end
