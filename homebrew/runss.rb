require 'formula'

class Runss < Formula
  VERSION = '0.1.4'

  homepage 'https://github.com/winebarrel/runss'
  url "https://github.com/winebarrel/runss/releases/download/v#{VERSION}/runss-v#{VERSION}-darwin-amd64.gz"
  sha256 'a168d41feadfd20bff196f10517384c88b8eb050ff94df0289336530ad871054'
  version VERSION
  head 'https://github.com/winebarrel/runss.git', :branch => 'master'

  def install
    system "mv runss-v#{VERSION}-darwin-amd64 runss"
    bin.install 'runss'
  end
end
