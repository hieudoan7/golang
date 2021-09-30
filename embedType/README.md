# Type Embedding in Go
example:
type A struct {
    B
    field_name C
}
B: called embedded type
A: called embedding type

Some notes:
- All public properties/methods of B can be access by A via 2 syntax:
  + A.(public properties/methods) (just like B)
  + A.B.(public properties/methods) (just like existing an hidden field name which is identical with type name)
- Because of characteristic above, we can conclude the differences between embedding type and explicit field name.
  + When using explitcit field name => composite
  + when using embedding: => inherit
