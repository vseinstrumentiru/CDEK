# Changelog
All notable changes to this project will be documented in this file.

## [Unreleased]
### Added
- Changelog
- Tests
- `ClientConf` now has constructor (`NewClientConf`) with setters

### Changed
- `Auth` is not necessary for client now
- Now you can path calculator url to SDK (`ClientConf.CalculatorURL`), default value is present while you use constructor

## [2.0.0] - 2019-07-09
### Added
- Documentation

### Changed
- renamed `ClientConf.XmlApiUrl` parameter to `ClientConf.CdekAPIURL`