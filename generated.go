package exitcode_go

const (

	// Success - Generic success code. (from libc)
	Success = 0

	// Failure - Generic failure or unspecified error. (from libc)
	Failure = 1

	// InvalidArgument - Invalid or excess arguments. (from LSB)
	InvalidArgument = 2

	// NotImplemented - Unimplemented feature. (from LSB)
	NotImplemented = 3

	// NoPermission - The user has insufficient privileges. (from LSB)
	NoPermission = 4

	// NotInstalled - The program is not installed. (from LSB)
	NotInstalled = 5

	// NotConfigured - The program is not configured. (from LSB)
	NotConfigured = 6

	// NotRunning - The program is not running. (from LSB)
	NotRunning = 7

	// Usage - Command line usage error (from BSD)
	Usage = 64

	// DataErr - Data format error (from BSD)
	DataErr = 65

	// NoInput - Cannot open input (from BSD)
	NoInput = 66

	// NoUser - Addressee unknown (from BSD)
	NoUser = 67

	// NoHost - Host name unknown (from BSD)
	NoHost = 68

	// Unavailable - Service unavailable (from BSD)
	Unavailable = 69

	// Software - internal software error (from BSD)
	Software = 70

	// OSErr - System error (e.g., can't fork) (from BSD)
	OSErr = 71

	// OSFile - Critical OS file missing (from BSD)
	OSFile = 72

	// CantCreat - Can't create (user) output file (from BSD)
	CantCreat = 73

	// IOErr - Input/output error (from BSD)
	IOErr = 74

	// TempFail - Temporary failure; user is invited to retry (from BSD)
	TempFail = 75

	// Protocol - Remote error in protocol (from BSD)
	Protocol = 76

	// NoPermBSD - Permission denied (from BSD)
	NoPermBSD = 77

	// ConfigBSD - Configuration error (from BSD)
	ConfigBSD = 78

	// Chdir - Changing to the requested working directory failed. (from systemd)
	Chdir = 200

	// Nice - Failed to set up process scheduling priority (nice level). (from systemd)
	Nice = 201

	// Fds - Failed to close unwanted file descriptors, or to adjust passed file descriptors. (from systemd)
	Fds = 202

	// Exec - The actual process execution failed (specifically, the execve(2) system call). Most likely this is caused by a missing or non-accessible executable file. (from systemd)
	Exec = 203

	// Memory - Failed to perform an action due to memory shortage. (from systemd)
	Memory = 204

	// Limits - Failed to adjust resource limits. (from systemd)
	Limits = 205

	// OomAdjust - Failed to adjust the OOM setting. (from systemd)
	OomAdjust = 206

	// SignalMask - Failed to set process signal mask. (from systemd)
	SignalMask = 207

	// Stdin - Failed to set up standard input. (from systemd)
	Stdin = 208

	// Stdout - Failed to set up standard output. (from systemd)
	Stdout = 209

	// Chroot - Failed to change root directory (chroot(2)). (from systemd)
	Chroot = 210

	// Ioprio - Failed to set up IO scheduling priority. (from systemd)
	Ioprio = 211

	// TimerSlack - Failed to set up timer slack. (from systemd)
	TimerSlack = 212

	// SecureBits - Failed to set process secure bits. (from systemd)
	SecureBits = 213

	// SetScheduler - Failed to set up CPU scheduling. (from systemd)
	SetScheduler = 214

	// CPUAffinity - Failed to set up CPU affinity. (from systemd)
	CPUAffinity = 215

	// Group - Failed to determine or change group credentials. (from systemd)
	Group = 216

	// User - Failed to determine or change user credentials, or to set up user namespacing. (from systemd)
	User = 217

	// Capabilities - Failed to drop capabilities, or apply ambient capabilities. (from systemd)
	Capabilities = 218

	// CGroup - Setting up the service control group failed. (from systemd)
	CGroup = 219

	// SetSID - Failed to create new process session. (from systemd)
	SetSID = 220

	// Confirm - Execution has been cancelled by the user. See the systemd.confirm_spawn= kernel command line setting on kernel-command-line(7) for details. (from systemd)
	Confirm = 221

	// StdErr - Failed to set up standard error output. (from systemd)
	StdErr = 222

	// PAM - Failed to set up PAM session. (from systemd)
	PAM = 224

	// Network - Failed to set up network namespacing. (from systemd)
	Network = 225

	// Namespace - Failed to set up mount, UTS, or IPC namespacing. (from systemd)
	Namespace = 226

	// NoNewPrivileges - Failed to disable new privileges. (from systemd)
	NoNewPrivileges = 227

	// Seccomp - Failed to apply system call filters. (from systemd)
	Seccomp = 228

	// SELinuxContext - Determining or changing SELinux context failed. (from systemd)
	SELinuxContext = 229

	// Personality - Failed to set up an execution domain (personality). (from systemd)
	Personality = 230

	// AppArmorProfile - Failed to prepare changing AppArmor profile. (from systemd)
	AppArmorProfile = 231

	// AddressFamilies - Failed to restrict address families. (from systemd)
	AddressFamilies = 232

	// RuntimeDirectory - Setting up runtime directory failed. (from systemd)
	RuntimeDirectory = 233

	// Chown - Failed to adjust socket ownership. Used for socket units only. (from systemd)
	Chown = 235

	// SmackProcessLabel - Failed to set SMACK label. (from systemd)
	SmackProcessLabel = 236

	// Keyring - Failed to set up kernel keyring. (from systemd)
	Keyring = 237

	// StateDirectory - Failed to set up unit's state directory. (from systemd)
	StateDirectory = 238

	// CacheDirectory - Failed to set up unit's cache directory. (from systemd)
	CacheDirectory = 239

	// LogsDirectory - Failed to set up unit's logging directory. (from systemd)
	LogsDirectory = 240

	// ConfigurationDirectory - Failed to set up unit's configuration directory. (from systemd)
	ConfigurationDirectory = 241

	// NUMAPolicy - Failed to set up unit's NUMA memory policy. (from systemd)
	NUMAPolicy = 242

	// Credentials - Failed to set up unit's credentials. (from systemd)
	Credentials = 243

	// BPF - Failed to apply BPF restrictions. (from systemd)
	BPF = 245
)
