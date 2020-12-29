//! Ataxia game engine library code
#![deny(
    missing_debug_implementations,
    missing_copy_implementations,
    trivial_casts,
    trivial_numeric_casts,
    unsafe_code,
    unused_import_braces,
    unused_qualifications,
    clippy::all,
    clippy::pedantic,
    clippy::perf,
    clippy::style
)]
#![forbid(unsafe_code)]
#![warn(missing_docs)]

pub mod config;

pub use crate::config::Config;
