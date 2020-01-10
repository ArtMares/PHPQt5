package main

type Comparison interface {
    EQ(Comparison) bool
    LS(Comparison) bool
    GR(Comparison) bool
    LE(Comparison) bool
    GE(Comparison) bool
}