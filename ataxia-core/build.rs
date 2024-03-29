use time::OffsetDateTime;

use std::env;
use std::fs::File;
use std::io::Write;
use std::path::Path;

fn main() {
    let out_dir = env::var("OUT_DIR").unwrap();
    let dest_path = Path::new(&out_dir).join("version.rs");
    let mut f = File::create(dest_path).unwrap();

    let output: String = format!(
        "static ATAXIA_COMPILED: &str = \"{}\";",
        OffsetDateTime::now_utc()
    );

    f.write_all(output.as_bytes()).unwrap();
}
