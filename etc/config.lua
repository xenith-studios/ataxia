-- Ataxia Configuration file
--
-- All of the configuration is written in plain Lua. Lua must be able to
-- parse this configuration script without any syntax errors. Basic options
-- are in the format of:
--      option = value
--
-- Since this is a regular Lua script, you can write any valid Lua code to
-- perform the configuration you wish. Most options will have defaults, so if
-- you leave them out, they will fall back to the hard-coded values.


-- Port to run the main engine on
main_port = 9000

-- Set this option to true to run in the background as a server daemon upon
-- startup. This is useful for certain startup scripts.
-- This option currently does not work due to Go lacking a proper fork utility.
-- Defaults to disabled.
-- daemonize = false

-- Location of the PID file
pid_file = "data/ataxia.pid"

-- Set this option to a full directory path to have Ataxia chroot on startup.
-- This option requires starting Ataxia with root priviledges.
chroot = ""

-- User to run ataxia as.
-- This option requires starting Ataxia with root priviledges.
user = ""

-- Group to run ataxia as.
-- This option requires starting Ataxia with root priviledges.
group = ""

-- Logging facility to use
-- Options are: standard, file, syslog
log_facility = "standard"

-- Location of the default log file for the file logger
-- log_file = "log/ataxia.log"

-- Email facility to use
-- Options are: none, smtp, sendmail
email_facility = "none"

-- Location of the sendmail binary
-- sendmail = "/usr/sbin/sendmail"

-- Administrator's email address for notifications
-- admin_email = ""

-- Email address to send abuse notifications
-- abuse_email = ""

-- Storage facility to use
-- Options are: file, database
storage_facility = "file"

-- Location of the world data
world_data = "data/world/"

-- Location of account data
account_data = "data/accounts/"

-- Location of player data
player_data = "data/players/"

-- Max simultaenous connections total
max_clients = 1000

-- Max simultaneous connections per IP
max_clients_per_host = 5

-- Allow creation of new accounts
account_creation = true

-- Maximum number of characters per account
chars_per_account = 3

-- Maximum number of simultaneous characters online per account
active_per_account = 1

-- Autosave interval in minutes
autosave = 15

-- Backup players
backup_players = true
