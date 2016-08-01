require 'formula'

class Runss < Formula
  VERSION = '0.1.2'

  homepage 'https://github.com/winebarrel/runss'
  url "https://github.com/winebarrel/runss/releases/download/v#{VERSION}/runss-v#{VERSION}-darwin-amd64.gz"
  sha256 '2562f8aca1eb576f9286f7ec0a59d27012ff230824f72e273e21a8614dbd4e72'
  version VERSION
  head 'https://github.com/winebarrel/runss.git', :branch => 'master'

  def install
    system "mv runss-v#{VERSION}-darwin-amd64 runss"
    bin.install 'runss'
  end
end
