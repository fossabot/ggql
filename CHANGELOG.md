# Changelog

Changes to the GraphQL package are listed here. Releases follow semantic
versioning.

## [1.0.0] - 2020-06-05

First release

### Added

- License file.

## [0.9.37] - 2020-06-04

### Added

- ggqlgen can now take multiple files for input.

### Changed

- Arguments for ggqlgen have changed to support multiple files.

## [0.9.36] - 2020-06-03

### Changed

- Updated golang version to 1.14. No other changes.

## [0.9.35] - 2020-05-04

### Fixed

- nil is now returned with coerce errors instead of the original value.

## [0.9.34] - 2020-05-04

### Fixed

- null is now allowed for scalar types.

## [0.9.33] - 2020-04-23

### Fixed

- Negative numbers were marked as invalid.

## [0.9.32] - 2020-04-12

### Fixed

- Float Scalar correctly converts int64 to Float which occurs when using JSON input.

## [0.9.31] - 2020-02-27

### Fixed

- scalar Time docstring updated.

- Allow null input Time.

## [0.9.30] - 2020-02-13

### Fixed

- Use of Unions and the @go directive fixed.

## [0.9.29] - 2020-02-11

### Fixed

- Non-Null Enum values were not checked correct but are now.

## [0.9.28] - 2020-02-07

### Fixed

- Fixed empty value bug that could cause infinite looping on invalid queries.

## [0.9.27] - 2020-02-03

### Fixed

- Executable parse bug that caused a hang on extra characters is fixed.

## [0.9.26] - 2020-01-30

### Changed

- ggqlgen can now use .go files as input, given the right conditions (see
  ggqlgen usage)

## [0.9.25] - 2020-01-30

### Fixed

- Check for null values in non-null list arguments.

### Added

- CORS headers added to example to allow use with GraphiQL.

## [0.9.24] - 2020-01-29

### Fixed

- Check enum argument values.

## [0.9.23] - 2020-01-20

Input checks

### Fixed

- Additional input validation is now performed.

- Input argument coercion performed in all cases.

## [0.9.22] - 2020-01-17

List nil check

### Fixed

- Empty lists like argList now check for a nil dictionary.

## [0.9.21] - 2020-01-16

Nested errors

### Fixed

- A `ggql.Errors` returned from a Resolve call is now flattened so it appears correctly in results.

## [0.9.20] - 2020-01-14

Validation error locations

### Fixed

- Add line and column to all schema validation errors.

## [0.9.19] - 2020-01-14

Stub generation

### Added

- Code generation for stubs and embedded schema SDL.

### Fixed

- `ofType` for NON_NULL and LIST now returns a Type instead of `null`.

- Maybe not a bug fix but the @deprecated directive now has a double quoted
  default value to match expectations from GraphiQL. This is a special case
  and no other string values are double quoted.

## [0.9.18] - 2020-01-07

Error handling

### Added

- Allow multiple errors in return as well as partial data.

- Include location and path in errors.

## [0.9.17] - 2020-01-02

Time output should be a string.

### Fixed

- TimeScalar CoerceOut now outputs an RFC3339 string.

## [0.9.16] - 2020-01-02

Better input validation.

### Added

- Additional type and coercion checks on input variables.

### Fixed

- Empty field names in a query now cause an error.

- Descriptions don't carry over to subsequent types.

## [0.9.15] - 2019-12-31

Exclude default types.

### Added

- When creating a new root the `Time` and `Int64` types can be excluded thereby
  allowing those type names to be implemented differently than the default.

## [0.9.14] - 2019-12-20

Time handled better.

### Fixed

- TimeScalar now accepts time.Time.

## [0.9.13] - 2019-12-12

Bug fixes.

### Fixed

- Empty optional bug fixed.

- Trim multi line documentation strings when parsing and format properly in
  output. Issue #15

- Parsing HTTP POST requests no longer misses the last character which is read
  but also returns and EOF error on HTTP Body Readers. Issue #14

- Variable substitution now is correct replaced in list as object instead of
  just scalars.
