//! Ataxia game engine library code
#![deny(
    missing_debug_implementations,
    missing_copy_implementations,
    trivial_casts,
    trivial_numeric_casts,
    unsafe_code,
    unused_import_braces,
    unused_qualifications
)]
#![warn(missing_docs)]

pub mod config;
pub mod server;

pub use crate::config::Config;
pub use crate::server::Server;
