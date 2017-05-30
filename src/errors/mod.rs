mod config_error;

pub use self::config_error::Error as ConfigError;
pub use self::config_error::ErrorKind as ConfigErrorKind;
pub use self::config_error::ResultExt as ConfigResultExt;
pub type ConfigResult<T> = Result<T, ConfigError>;
