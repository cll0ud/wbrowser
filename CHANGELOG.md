# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.0.2] - 2021-05-21
### Added
- Changelog file.

### Changed
- Improve "README" file with more information about the config file.
- Change `log` to `logger` to avoid package name collision.
- Change the behavior for when there's no `domain` or `target`: the 
  program will now just open the default browser instead of throwing an error.
  
### Fixed
- Fix string check for when there's no rule for that specific domain.

[Unreleased]: https://github.com/cll0ud/wbrowser/compare/1.0.2...HEAD
[1.0.2]: https://github.com/cll0ud/wbrowser/compare/1.0.1...1.0.2
