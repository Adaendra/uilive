# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.2.0] - 2023-07-22
### Changed
- Up Golang version to **1.20**
- Code reorganisation

### Fixed
- Make "getTermSize()" use local variables instead of package-level ones [fork #39](https://github.com/gosuri/uilive/pull/39)
- Writer.Flush: properly compute line count for buffers longer than 2 lines [fork #33](https://github.com/gosuri/uilive/pull/33)
- Fix wrong screen drawing by not counting ASCII color codes [fork #30](https://github.com/gosuri/uilive/pull/30)

## [0.1.0] - 2023-07-22
### Added
- Fork of the project [gosuri/uilive](https://github.com/Adaendra/uilive) v0.0.4
- `.gitignore` file

### Removed
- `.trvis.yml` file