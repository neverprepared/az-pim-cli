class AzPimCli < Formula
  desc "CLI utility for managing Azure PIM role assignments"
  homepage "https://github.com/mindmorass/az-pim-cli"
  license "MIT"
  head "https://github.com/mindmorass/az-pim-cli.git", branch: "main"

  depends_on "go" => :build

  def install
    ldflags = %W[-s -w -X main.version=HEAD -X main.tag=HEAD]
    system "go", "build", *std_go_args(ldflags: ldflags), "."
  end

  test do
    assert_match "az-pim-cli", shell_output("#{bin}/az-pim-cli --help")
  end
end
