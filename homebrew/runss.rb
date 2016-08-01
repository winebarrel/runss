require 'formula'

class Runss < Formula
  VERSION = '0.1.0'

  homepage 'https://github.com/winebarrel/runss'
  url "https://github.com/winebarrel/runss/releases/download/v#{VERSION}/runss-v#{VERSION}-darwin-amd64.gz"
  sha256 '0be1ba60a5b6e50894e581e5f9b59d10f28aedec3e3ea610c6f024fc83e4b1a5'
  version VERSION
  head 'https://github.com/winebarrel/runss.git', :branch => 'master'

  def install
    system "mv runss-v#{VERSION}-darwin-amd64 runss"
    bin.install 'runss'
  end
end
