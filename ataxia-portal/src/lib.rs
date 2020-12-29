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
#![forbid(unsafe_code)]
#![warn(missing_docs, clippy::pedantic)]
// Disable this lint due to Tokio using the binding name _task_context
#![recursion_limit = "256"]

pub mod portal;

pub use crate::portal::Portal;
