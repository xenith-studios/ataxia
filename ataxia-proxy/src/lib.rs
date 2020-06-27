//! Ataxia game engine library code
#![deny(
    missing_debug_implementations,
    missing_copy_implementations,
    trivial_casts,
    trivial_numeric_casts,
    unsafe_code,
    unused_import_braces,
    unused_qualifications,
    clippy::all
)]
#![warn(missing_docs, clippy::pedantic)]

pub mod proxy;

pub use crate::proxy::Proxy;
