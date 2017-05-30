use std::io;

use toml::de::Error as TomlError;

error_chain! {
    // The type defined for this error. These are the conventional
    // and recommended names, but they can be arbitrarily chosen.
    //
    // It is also possible to leave this section out entirely, or
    // leave it empty, and these names will be used automatically.
    types {
        Error, ErrorKind, ResultExt, Result;
    }

    // Automatic conversions between this error chain and other
    // error chains. In this case, it will e.g. generate an
    // `ErrorKind` variant called `Another` which in turn contains
    // the `other_error::ErrorKind`, with conversions from
    // `other_error::Error`.
    //
    // Optionally, some attributes can be added to a variant.
    //
    // This section can be empty.
    links {}

    // Automatic conversions between this error chain and other
    // error types not defined by the `error_chain!`. These will be
    // wrapped in a new error with, in the first case, the
    // `ErrorKind::Fmt` variant. The description and cause will
    // forward to the description and cause of the original error.
    //
    // Optionally, some attributes can be added to a variant.
    //
    // This section can be empty.
    foreign_links {
        Io(io::Error) #[cfg(unix)];
        Parser(TomlError);
    }

    // Define additional `ErrorKind` variants. The syntax here is
    // the same as `quick_error!`, but the `from()` and `cause()`
    // syntax is not supported.
    errors {
      VecParserError(vec: Vec<TomlError>) {
        description("parser errors")
        display("parser errors: {:?}", vec)
      }
    }
}
