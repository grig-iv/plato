{
  description = "My prompt flake";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-24.05";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = {
    self,
    nixpkgs,
    flake-utils,
  }:
    flake-utils.lib.eachDefaultSystem (
      system: let
        pkgs = import nixpkgs {inherit system;};
      in {
        packages.default = pkgs.buildGoModule {
          name = "plato";
          src = ./.;
          vendorHash = "sha256-qf787Dg6Xy3HGkgroATQ+Xd15muLpkVDuzFskBktEXo=";
          meta = {
            mainProgram = "plato";
          };
        };
      }
    );
}
