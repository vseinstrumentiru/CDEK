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
- `Pvz.IsDressingRoom`, `Pvz.HaveCashless` and `Pvz.AllowedCod` now `*bool` (was `*string`)
- `DeleteOrderReq.OrderCount` now `*int`, (was `*string`)

## [2.0.0] - 2019-07-06
### Added
- Documentation

### Changed
- renamed `ClientConf.XmlApiUrl` parameter to `ClientConf.CdekAPIURL`
