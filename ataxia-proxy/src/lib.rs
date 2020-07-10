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
// Disable this lint due to Tokio using the binding name _task_context
#![allow(clippy::used_underscore_binding)]
#![recursion_limit = "256"]

pub mod proxy;

pub use crate::proxy::Proxy;
