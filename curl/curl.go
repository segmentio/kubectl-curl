package curl

import (
	"sort"
	"time"
)

func AbstractUnixSocket(path string) Option {
	return fileOption("--abstract-unix-socket", "", path, "")
}

func AltSvc(path string) Option {
	return fileOption("--alt-svc", "", path, "")
}

func Anyauth(on bool) Option {
	return boolOption("--anyauth", "", on, "")
}

func Append(on bool) Option {
	return boolOption("--append", "-a", on, "")
}

func Basic(on bool) Option {
	return boolOption("--basic", "", on, "")
}

func CACert(path string) Option {
	return fileOption("--cacert", "", path, "")
}

func CAPath(path string) Option {
	return dirOption("--capath", "", path, "")
}

func CertStatus(on bool) Option {
	return boolOption("--cert-status", "", on, "")
}

func CertType(typ string) Option {
	return typeOption("--cert-type", "", typ, "")
}

func Cert(cert string) Option {
	return certificateOption("--cert", "-E", cert, "")
}

func Ciphers(ciphers ...string) Option {
	return cipherListOption("--ciphers", "", ciphers, "")
}

func CompressedSSH(on bool) Option {
	return boolOption("--compressed-ssh", "", on, "")
}

func Compressed(on bool) Option {
	return boolOption("--compressed", "", on, "")
}

func Config(path string) Option {
	return fileOption("--config", "-K", path, "")
}

func ConnectTimeout(timeout time.Duration) Option {
	return secondsOption("--connect-timeout", "", timeout, "")
}

func ConnectTo(addr string) Option {
	return connectOption("--connect-to", "", addr, "")
}

func ContinueAt(offset int64) Option {
	return offsetOption("--continue-at", "-C", offset, "")
}

func CookieJar(path string) Option {
	return fileOption("--cookie-jar", "-c", path, "")
}

func Cookie(dataOrFile string) Option {
	return dataOption("--cookie", "-b", []byte(dataOrFile), "")
}

func CreateDirs(on bool) Option {
	return boolOption("--create-dirs", "", on, "")
}

func CRLF(on bool) Option {
	return boolOption("--crlf", "", on, "")
}

func CRLFile(path string) Option {
	return fileOption("--crlfile", "", path, "")
}

func DataASCII(data []byte) Option {
	return dataOption("--data-ascii", "", data, "")
}

func DataBinary(data []byte) Option {
	return dataOption("--data-binary", "", data, "")
}

func DataRaw(data []byte) Option {
	return dataOption("--data-raw", "", data, "")
}

func DataUrlencode(data []byte) Option {
	return dataOption("--data-urlencode", "", data, "")
}

func Data(data []byte) Option {
	return dataOption("--data", "-d", data, "")
}

func Delegation(level string) Option {
	return stringOption("--delegation", "", level, "")
}

func Digest(on bool) Option {
	return boolOption("--digest", "", on, "")
}

func DisableEPSV(on bool) Option {
	return boolOption("--disable-epsv", "", on, "")
}

func Disable(on bool) Option {
	return boolOption("--disable", "-q", on, "")
}

func DisallowUsernameInURL(on bool) Option {
	return boolOption("--disallow-username-in-url", "", on, "")
}

func DNSInterface(iface string) Option {
	return interfaceOption("--dns-interface", "", iface, "")
}

func DNSIPv4Addr(addr string) Option {
	return addressOption("--dns-ipv4-addr", "", addr, "")
}

func DNSIPv6Addr(addr string) Option {
	return addressOption("--dns-ipv6-addr", "", addr, "")
}

func DNSServers(addrs ...string) Option {
	return addressListOption("--dns-servers", "", addrs, "")
}

func DoHURL(url string) Option {
	return urlOption("--doh-url", "", url, "")
}

func DumpHeader(path string) Option {
	return fileOption("--dump-header", "-D", path, "")
}

func EGDFile(path string) Option {
	return fileOption("--egd-file", "", path, "")
}

func Engine(name string) Option {
	return nameOption("--engine", "", name, "")
}

func Expect100Timeout(timeout time.Duration) Option {
	return secondsOption("--expect100-timeout", "", timeout, "")
}

func FailEarly(on bool) Option {
	return boolOption("--fail-early", "", on, "")
}

func Fail(on bool) Option {
	return boolOption("--fail", "-f", on, "")
}

func FormString(form string) Option {
	return stringOption("--form-string", "", form, "")
}

func Form(form string) Option {
	return stringOption("--form", "-F", form, "")
}

func FTPAccount(data []byte) Option {
	return dataOption("--ftp-account", "", data, "")
}

func FTPAlternativeToUser(command string) Option {
	return stringOption("--ftp-alternative-to-user", "", command, "")
}

func FTPCreateDirs(on bool) Option {
	return boolOption("--ftp-create-dirs", "", on, "")
}

func FTPMethod(method string) Option {
	return methodOption("--ftp-method", "", method, "")
}

func FTPPasv(on bool) Option {
	return boolOption("--ftp-pasv", "", on, "")
}

func FTPPort(addr string) Option {
	return addressOption("--ftp-port", "-P", addr, "")
}

func FTPPret(on bool) Option {
	return boolOption("--ftp-pret", "", on, "")
}

func FTPSkipPasvIP(on bool) Option {
	return boolOption("--ftp-skip-pasv-ip", "", on, "")
}

func FTPSSLCCCMode(mode string) Option {
	return stringOption("--ftp-ssl-ccc-mode", "", mode, "")
}

func FTPSSLCCC(on bool) Option {
	return boolOption("--ftp-ssl-ccc", "", on, "")
}

func Get(on bool) Option {
	return boolOption("--get", "-G", on, "")
}

func Globoff(on bool) Option {
	return boolOption("--globoff", "-g", on, "")
}

func HappyEyeballsTimeout(timeout time.Duration) Option {
	return millisecondsOption("--happy-eyeballs-timeout-ms", "", timeout, "")
}

func HAProxyProtocol(on bool) Option {
	return boolOption("--haproxy-protocol", "", on, "")
}

func Head(on bool) Option {
	return boolOption("--head", "-I", on, "")
}

func Header(header string) Option {
	return headerOption("--header", "-H", header, "")
}

func Hostpubmd5(md5 string) Option {
	return stringOption("--hostpubmd5", "", md5, "")
}

func HTTP09(on bool) Option {
	return boolOption("--http0.9", "", on, "")
}

func HTTP11(on bool) Option {
	return boolOption("--http1.1", "", on, "")
}

func HTTP2PriorKnowledge(on bool) Option {
	return boolOption("--http2-prior-knowledge", "", on, "")
}

func HTTP2(on bool) Option {
	return boolOption("--http2", "", on, "")
}

func IgnoreContentLength(on bool) Option {
	return boolOption("--ignore-content-length", "", on, "")
}

func Include(on bool) Option {
	return boolOption("--include", "-i", on, "")
}

func Insecure(on bool) Option {
	return boolOption("--insecure", "-k", on, "")
}

func Interface(iface string) Option {
	return interfaceOption("--interface", "", iface, "")
}

func IPv4(on bool) Option {
	return boolOption("--ipv4", "-4", on, "")
}

func IPv6(on bool) Option {
	return boolOption("--ipv6", "-6", on, "")
}

func JunkSessionCookies(on bool) Option {
	return boolOption("--junk-session-cookies", "-j", on, "")
}

func KeepaliveTime(keepalive time.Duration) Option {
	return secondsOption("--keepalive-time", "", keepalive, "")
}

func KeyType(typ string) Option {
	return typeOption("--key-type", "", typ, "")
}

func Key(key string) Option {
	return keyOption("--key", "", key, "")
}

func KRB(level string) Option {
	return stringOption("--krb", "", level, "")
}

func Libcurl(path string) Option {
	return fileOption("--libcurl", "", path, "")
}

func LimitRate(speed string) Option {
	return speedOption("--limit-rate", "", speed, "")
}

func ListOnly(on bool) Option {
	return boolOption("--list-only", "-l", on, "")
}

func LocalPort(numberOrRange string) Option {
	return portOption("--local-port", "", numberOrRange, "")
}

func Location(on bool) Option {
	return boolOption("--location", "-L", on, "")
}

func LoginOptions(options string) Option {
	return stringOption("--login-options", "", options, "")
}

func MailAuth(addr string) Option {
	return addressOption("--mail-auth", "", addr, "")
}

func MailFrom(addr string) Option {
	return addressOption("--mail-from", "", addr, "")
}

func MailRcpt(addr string) Option {
	return addressOption("--mail-rcpt", "", addr, "")
}

func Manual(on bool) Option {
	return boolOption("--manual", "-M", on, "")
}

func MaxFilesize(bytes int64) Option {
	return bytesOption("--max-filesize", "", bytes, "")
}

func MaxRedirs(num int) Option {
	return numberOption("--max-redirs", "", num, "")
}

func MaxTime(max time.Duration) Option {
	return secondsOption("--max-time", "-m", max, "")
}

func Metalink(on bool) Option {
	return boolOption("--metalink", "", on, "")
}

func Negotiate(on bool) Option {
	return boolOption("--negotiate", "", on, "")
}

func NetrcFile(path string) Option {
	return fileOption("--netrc-file", "", path, "")
}

func NetrcOptional(on bool) Option {
	return boolOption("--netrc-optional", "", on, "")
}

func Netrc(on bool) Option {
	return boolOption("--netrc", "-n", on, "")
}

func Next(on bool) Option {
	return boolOption("--next", "-:", on, "")
}

func NoALPN(on bool) Option {
	return boolOption("--no-alpn", "", on, "")
}

func NoBuffer(on bool) Option {
	return boolOption("--no-buffer", "-N", on, "")
}

func NoKeepalive(on bool) Option {
	return boolOption("--no-keepalive", "", on, "")
}

func NoNPN(on bool) Option {
	return boolOption("--no-npn", "", on, "")
}

func NoSessionID(on bool) Option {
	return boolOption("--no-sessionid", "", on, "")
}

func NoProxy(addrs ...string) Option {
	return addressListOption("--noproxy", "", addrs, "")
}

func NTLMWB(on bool) Option {
	return boolOption("--ntlm-wb", "", on, "")
}

func NTLM(on bool) Option {
	return boolOption("--ntlm", "", on, "")
}

func OAuth2Bearer(token string) Option {
	return tokenOption("--oauth2-bearer", "", token, "")
}

func Output(path string) Option {
	return fileOption("--output", "-o", path, "")
}

func Pass(phrase string) Option {
	return stringOption("--pass", "", phrase, "")
}

func PathAsIs(on bool) Option {
	return boolOption("--path-as-is", "", on, "")
}

func PinnedPubKey(path string) Option {
	return fileOption("--pinnedpubkey", "", path, "")
}

func Post301(on bool) Option {
	return boolOption("--post301", "", on, "")
}

func Post302(on bool) Option {
	return boolOption("--post302", "", on, "")
}

func Post303(on bool) Option {
	return boolOption("--post303", "", on, "")
}

func Preproxy(url string) Option {
	return proxyAddrOption("--preproxy", "", url, "")
}

func ProgressBar(on bool) Option {
	return boolOption("--progress-bar", "-#", on, "")
}

func ProtoDefault(protocol string) Option {
	return protocolOption("--proto-default", "", protocol, "")
}

func ProtoRedir(protocols ...string) Option {
	return protocolListOption("--proto-redir", "", protocols, "")
}

func Proto(protocols ...string) Option {
	return protocolListOption("--proto", "", protocols, "")
}

func ProxyAnyauth(on bool) Option {
	return boolOption("--proxy-anyauth", "", on, "")
}

func ProxyBasic(on bool) Option {
	return boolOption("--proxy-basic", "", on, "")
}

func ProxyCACert(path string) Option {
	return fileOption("--proxy-cacert", "", path, "")
}

func ProxyCertType(typ string) Option {
	return typeOption("--proxy-cert-type", "", typ, "")
}

func ProxyCert(cert string) Option {
	return certificateOption("--proxy-cert", "", cert, "")
}

func ProxyCiphers(ciphers ...string) Option {
	return cipherListOption("--proxy-ciphers", "", ciphers, "")
}

func ProxyCRLFile(path string) Option {
	return fileOption("--proxy-crlfile", "", path, "")
}

func ProxyDiget(on bool) Option {
	return boolOption("--proxy-digest", "", on, "")
}

func ProxyHeader(header string) Option {
	return headerOption("--proxy-header", "", header, "")
}

func ProxyInsecure(on bool) Option {
	return boolOption("--proxy-insecure", "", on, "")
}

func ProxyKeyType(typ string) Option {
	return typeOption("--proxy-key-type", "", typ, "")
}

func ProxyKey(key string) Option {
	return keyOption("--proxy-key", "", key, "")
}

func ProxyNegotiate(on bool) Option {
	return boolOption("--proxy-negotiate", "", on, "")
}

func ProxyNTLM(on bool) Option {
	return boolOption("--proxy-ntlm", "", on, "")
}

func ProxyPass(phrase string) Option {
	return stringOption("--proxy-pass", "", phrase, "")
}

func ProxyPinnedpubkey(path string) Option {
	return fileOption("--proxy-pinnedpubkey", "", path, "")
}

func ProxyServiceName(name string) Option {
	return nameOption("--proxy-service-name", "", name, "")
}

func ProxySSLAllowBeast(on bool) Option {
	return boolOption("--proxy-ssl-allow-beast", "", on, "")
}

func ProxyTLS13Ciphers(ciphers ...string) Option {
	return cipherListOption("--proxy-tls13-ciphers", "", ciphers, "")
}

func ProxyTLSAuthType(typ string) Option {
	return typeOption("--proxy-tlsauthtype", "", typ, "")
}

func ProxyTLSPassword(password string) Option {
	return stringOption("--proxy-tlspassword", "", password, "")
}

func ProxyTLSUser(user string) Option {
	return nameOption("--proxy-tlsuser", "", user, "")
}

func ProxyTLSv1(on bool) Option {
	return boolOption("--proxy-tlsv1", "", on, "")
}

func ProxyUser(user string) Option {
	return userOption("--proxy-user", "-U", user, "")
}

func Proxy(addr string) Option {
	return proxyAddrOption("--proxy", "-x", addr, "")
}

func Proxy10(hostPort string) Option {
	return hostPortOption("--proxy1.0", "", hostPort, "")
}

func ProxyTunnel(on bool) Option {
	return boolOption("--proxytunnel", "-p", on, "")
}

func Pubkey(key string) Option {
	return keyOption("--pubkey", "", key, "")
}

func Quote(on bool) Option {
	return boolOption("--quote", "-Q", on, "")
}

func RandomFile(path string) Option {
	return fileOption("--random-file", "", path, "")
}

func Range(r string) Option {
	return rangeOption("--range", "-r", r, "")
}

func Raw(on bool) Option {
	return boolOption("--raw", "", on, "")
}

func Referer(url string) Option {
	return urlOption("--referer", "-e", url, "")
}

func RemoteHeaderName(on bool) Option {
	return boolOption("--remote-header-name", "-J", on, "")
}

func RemoteNameAll(on bool) Option {
	return boolOption("--remote-name-all", "", on, "")
}

func RemoteName(on bool) Option {
	return boolOption("--remote-name", "-O", on, "")
}

func RemoteTime(on bool) Option {
	return boolOption("--remote-time", "-R", on, "")
}

func RequestTarget(on bool) Option {
	return boolOption("--remote-target", "", on, "")
}

func Request(method string) Option {
	return methodOption("--request", "-X", method, "")
}

func Resolve(resolve ...string) Option {
	return addressListOption("--resolve", "", resolve, "")
}

func RetryConnrefused(on bool) Option {
	return boolOption("--retry-connrefused", "", on, "")
}

func RetryDelay(delay time.Duration) Option {
	return secondsOption("--retry-delay", "", delay, "")
}

func RetryMaxTime(limit time.Duration) Option {
	return secondsOption("--retry-max-time", "", limit, "")
}

func Retry(limit int) Option {
	return numberOption("--retry", "", limit, "")
}

func SASLIR(on bool) Option {
	return boolOption("--sasl-ir", "", on, "")
}

func ServiceName(name string) Option {
	return nameOption("--service-name", "", name, "")
}

func ShowError(on bool) Option {
	return boolOption("--show-error", "-S", on, "")
}

func Silent(on bool) Option {
	return boolOption("--silent", "-s", on, "")
}

func SOCKS4(hostPort string) Option {
	return hostPortOption("--socks4", "", hostPort, "")
}

func SOCKS4a(hostPort string) Option {
	return hostPortOption("--socks4a", "", hostPort, "")
}

func SOCKS5Basic(on bool) Option {
	return boolOption("--socks5-basic", "", on, "")
}

func SOCKS5GssAPIService(name string) Option {
	return nameOption("--socks5-gssapi-service", "", name, "")
}

func SOCKS5GssAPI(on bool) Option {
	return boolOption("--socks5-gssapi", "", on, "")
}

func SOCKS5Hostname(hostPort string) Option {
	return hostPortOption("--socks5-hostname", "", hostPort, "")
}

func SOCKS5(hostPort string) Option {
	return hostPortOption("--socks5", "", hostPort, "")
}

func SpeedLimit(speed string) Option {
	return speedOption("--speed-limit", "-Y", speed, "")
}

func SpeedTime(speed time.Duration) Option {
	return secondsOption("--speed-time", "-y", speed, "")
}

func SSLAllowBeast(on bool) Option {
	return boolOption("--ssl-allow-beast", "", on, "")
}

func SSLNoRevoke(on bool) Option {
	return boolOption("--ssl-no-revoke", "", on, "")
}

func SSLReqd(on bool) Option {
	return boolOption("--ssl-reqd", "", on, "")
}

func SSL(on bool) Option {
	return boolOption("--ssl", "", on, "")
}

func SSLv2(on bool) Option {
	return boolOption("--sslv2", "-2", on, "")
}

func SSLv3(on bool) Option {
	return boolOption("--sslv3", "-3", on, "")
}

func Stderr(on bool) Option {
	return boolOption("--stderr", "", on, "")
}

func StyledOutput(on bool) Option {
	return boolOption("--styled-output", "", on, "")
}

func SuppressConnectHeaders(on bool) Option {
	return boolOption("--suppress-connect-headers", "", on, "")
}

func TCPFastOpen(on bool) Option {
	return boolOption("--tcp-fastopen", "", on, "")
}

func TCPNoDelay(on bool) Option {
	return boolOption("--tcp-nodelay", "", on, "")
}

func TelnetOption(opt string) Option {
	return stringOption("--telnet-option", "-t", opt, "")
}

func TFTPBlkSize(size int) Option {
	return numberOption("--tftp-blksize", "", size, "")
}

func TFTPNoOptions(on bool) Option {
	return boolOption("--tftp-no-options", "", on, "")
}

func TimeCond(date string) Option {
	return stringOption("--time-cond", "-z", date, "")
}

func TLSMax(version string) Option {
	return stringOption("--tls-max", "", version, "")
}

func TLS13Ciphers(ciphers ...string) Option {
	return cipherListOption("--tls13-ciphers", "", ciphers, "")
}

func TLSAuthType(typ string) Option {
	return typeOption("--tlsauthtype", "", typ, "")
}

func TLSPassword(on bool) Option {
	return boolOption("--tslpassword", "", on, "")
}

func TLSUser(user string) Option {
	return nameOption("--tlsuser", "", user, "")
}

func TLSv10(on bool) Option {
	return boolOption("--tlsv1.0", "", on, "")
}

func TLSv11(on bool) Option {
	return boolOption("--tlsv1.1", "", on, "")
}

func TLSv12(on bool) Option {
	return boolOption("--tlsv1.2", "", on, "")
}

func TLSv13(on bool) Option {
	return boolOption("--tlsv1.3", "", on, "")
}

func TSLv1(on bool) Option {
	return boolOption("--tlsv1", "-1", on, "")
}

func TrEncoding(on bool) Option {
	return boolOption("--tr-encoding", "", on, "")
}

func TraceASCII(path string) Option {
	return fileOption("--trace-ascii", "", path, "")
}

func TraceTime(on bool) Option {
	return boolOption("--trace-time", "", on, "")
}

func Trace(path string) Option {
	return fileOption("--trace", "", path, "")
}

func UnixSocket(path string) Option {
	return fileOption("--unix-socket", "", path, "")
}

func UploadFile(path string) Option {
	return fileOption("--upload-file", "-T", path, "")
}

func Url(url string) Option {
	return urlOption("--url", "", url, "")
}

func UseASCII(on bool) Option {
	return boolOption("--use-ascii", "-B", on, "")
}

func UserAgent(userAgent string) Option {
	return nameOption("--user-agent", "", userAgent, "")
}

func User(userPassword string) Option {
	return userOption("--user", "-u", userPassword, "")
}

func Verbose(on bool) Option {
	return boolOption("--verbose", "-v", on, "")
}

func Version(on bool) Option {
	return boolOption("--version", "-V", on, "")
}

func WriteOut(format string) Option {
	return stringOption("--write-out", "-w", format, "")
}

func XAttr(on bool) Option {
	return boolOption("--xattr", "", on, "")
}

func NewOptionSet() OptionSet {
	options := OptionSet{
		AbstractUnixSocket(""),
		AltSvc(""),
		Anyauth(false),
		Append(false),
		Basic(false),
		CACert(""),
		CAPath(""),
		CertStatus(false),
		CertType(""),
		Cert(""),
		Ciphers(),
		CompressedSSH(false),
		Compressed(false),
		Config(""),
		ConnectTimeout(0),
		ConnectTo(""),
		ContinueAt(0),
		CookieJar(""),
		Cookie(""),
		CreateDirs(false),
		CRLF(false),
		CRLFile(""),
		DataASCII(nil),
		DataBinary(nil),
		DataRaw(nil),
		DataUrlencode(nil),
		Data(nil),
		Delegation(""),
		Digest(false),
		DisableEPSV(false),
		Disable(false),
		DisallowUsernameInURL(false),
		DNSInterface(""),
		DNSIPv4Addr(""),
		DNSIPv6Addr(""),
		DNSServers(),
		DoHURL(""),
		DumpHeader(""),
		EGDFile(""),
		Engine(""),
		Expect100Timeout(0),
		FailEarly(false),
		Fail(false),
		FormString(""),
		Form(""),
		FTPAccount(nil),
		FTPAlternativeToUser(""),
		FTPCreateDirs(false),
		FTPMethod(""),
		FTPPasv(false),
		FTPPort(""),
		FTPPret(false),
		FTPSkipPasvIP(false),
		FTPSSLCCCMode(""),
		FTPSSLCCC(false),
		Get(false),
		Globoff(false),
		HappyEyeballsTimeout(0),
		HAProxyProtocol(false),
		Head(false),
		Header(""),
		Hostpubmd5(""),
		HTTP09(false),
		HTTP11(false),
		HTTP2PriorKnowledge(false),
		HTTP2(false),
		IgnoreContentLength(false),
		Include(false),
		Insecure(false),
		Interface(""),
		IPv4(false),
		IPv6(false),
		JunkSessionCookies(false),
		KeepaliveTime(0),
		KeyType(""),
		Key(""),
		KRB(""),
		Libcurl(""),
		LimitRate(""),
		ListOnly(false),
		LocalPort(""),
		Location(false),
		LoginOptions(""),
		MailAuth(""),
		MailFrom(""),
		MailRcpt(""),
		Manual(false),
		MaxFilesize(0),
		MaxRedirs(0),
		MaxTime(0),
		Metalink(false),
		Negotiate(false),
		NetrcFile(""),
		NetrcOptional(false),
		Netrc(false),
		Next(false),
		NoALPN(false),
		NoBuffer(false),
		NoKeepalive(false),
		NoNPN(false),
		NoSessionID(false),
		NoProxy(),
		NTLMWB(false),
		NTLM(false),
		OAuth2Bearer(""),
		Output(""),
		Pass(""),
		PathAsIs(false),
		PinnedPubKey(""),
		Post301(false),
		Post302(false),
		Post303(false),
		Preproxy(""),
		ProgressBar(false),
		ProtoDefault(""),
		ProtoRedir(),
		Proto(),
		ProxyAnyauth(false),
		ProxyBasic(false),
		ProxyCACert(""),
		ProxyCertType(""),
		ProxyCert(""),
		ProxyCiphers(),
		ProxyCRLFile(""),
		ProxyDiget(false),
		ProxyHeader(""),
		ProxyInsecure(false),
		ProxyKeyType(""),
		ProxyKey(""),
		ProxyNegotiate(false),
		ProxyNTLM(false),
		ProxyPass(""),
		ProxyPinnedpubkey(""),
		ProxyServiceName(""),
		ProxySSLAllowBeast(false),
		ProxyTLS13Ciphers(),
		ProxyTLSAuthType(""),
		ProxyTLSPassword(""),
		ProxyTLSUser(""),
		ProxyTLSv1(false),
		ProxyUser(""),
		Proxy(""),
		Proxy10(""),
		ProxyTunnel(false),
		Pubkey(""),
		Quote(false),
		RandomFile(""),
		Range(""),
		Raw(false),
		Referer(""),
		RemoteHeaderName(false),
		RemoteNameAll(false),
		RemoteName(false),
		RemoteTime(false),
		RequestTarget(false),
		Request(""),
		Resolve(),
		RetryConnrefused(false),
		RetryDelay(0),
		RetryMaxTime(0),
		Retry(0),
		SASLIR(false),
		ServiceName(""),
		ShowError(false),
		Silent(false),
		SOCKS4(""),
		SOCKS4a(""),
		SOCKS5Basic(false),
		SOCKS5GssAPIService(""),
		SOCKS5GssAPI(false),
		SOCKS5Hostname(""),
		SOCKS5(""),
		SpeedLimit(""),
		SpeedTime(0),
		SSLAllowBeast(false),
		SSLNoRevoke(false),
		SSLReqd(false),
		SSL(false),
		SSLv2(false),
		SSLv3(false),
		Stderr(false),
		StyledOutput(false),
		SuppressConnectHeaders(false),
		TCPFastOpen(false),
		TCPNoDelay(false),
		TelnetOption(""),
		TFTPBlkSize(0),
		TFTPNoOptions(false),
		TimeCond(""),
		TLSMax(""),
		TLS13Ciphers(),
		TLSAuthType(""),
		TLSPassword(false),
		TLSUser(""),
		TLSv10(false),
		TLSv11(false),
		TLSv12(false),
		TLSv13(false),
		TSLv1(false),
		TrEncoding(false),
		TraceASCII(""),
		TraceTime(false),
		Trace(""),
		UnixSocket(""),
		UploadFile(""),
		Url(""),
		UseASCII(false),
		UserAgent(""),
		User(""),
		Verbose(false),
		Version(false),
		WriteOut(""),
		XAttr(false),
	}

	sort.Sort(options)
	return options
}
