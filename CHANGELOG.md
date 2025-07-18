v2.0.0 (compared to v1.5.3)

New:

- **Gonfique configs**: Collect all customizations to one place.
  - **Path directives**:
    - **Export**: create named declarations with automatically generated meaningful-stable names
    - **Replace**: overwrite resolved types for better integration
    - **Declare**: create named declarations with the provided name
    - **Mapify**: make it a map instead of a struct
  - **Type directives**:
    - **Parent refs**: add parent refs to structs
    - **Embed**: make the hierarchy between generated types explicit
    - **Iterator**: make your structs iterable with a for-range loop
    - **Accessors**: implement getters and setters on fields of your choice

Removed flags

- `-organize` flag is removed. `export` and `iterate` directives should cover the usecases.
- `-use` flag is removed. `replace` directive should cover the usecase.
- `-mappings` flag is removed. `declare` directive should cover the usecase.

v2.0.0 (compared to v1.6.5)

- Using `parent` or `accessors` is now not possible without `declare`. Previously using either of those without `declare` would result implicit type declaration with autogenerated name.
