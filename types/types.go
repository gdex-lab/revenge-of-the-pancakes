package types

// Pancake represents a happy-face pancake with choclate chips only the on the "true" side
type Pancake bool

// PancakeStack can be any length slice of the custom pancake type (bool). It is read from left (top) to right (bottom).
type PancakeStack []Pancake
