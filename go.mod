module goki.dev/mat32/v2

go 1.21.0

toolchain go1.21.4

retract (
	v2.0.1 // Retracts v2.0.0
	v2.0.0 // Published accidentally
)

require (
	goki.dev/enums v0.9.47
	golang.org/x/image v0.14.0
)
