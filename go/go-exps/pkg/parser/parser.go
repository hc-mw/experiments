package ApacheLogParser

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
)

var placeholderPatterns = map[string]string{
	"%":     `%`,                                                 // Literal percent sign
	"%a":    `(?P<client_ip>\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})`, // Client IP address
	"%{c}a": `(?P<peer_ip>\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})`,   // Peer IP address
	"%A":    `(?P<local_ip>\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})`,  // Local IP address
	"%B":    `(?P<response_size>\d+)`,                            // Size of response in bytes
	"%b":    `(?P<response_size_clf>\d+|-)`,                      // Size of response in bytes (CLF format)
	//"%{VARNAME}C":   `(?P<cookie_value>.*?)`,                                  // Contents of cookie VARNAME
	"%D": `(?P<request_time_micro>\d+)`, // Time taken to serve the request (microseconds)
	//"%{VARNAME}e":   `(?P<env_var_value>.*?)`,                                 // Contents of environment variable VARNAME
	"%f":    `(?P<filename>.*?)`,         // Filename
	"%h":    `(?P<remote_hostname>\S+)`,  // Remote hostname
	"%{c}h": `(?P<peer_hostname>\S+)`,    // Underlying peer hostname
	"%H":    `(?P<request_protocol>\S+)`, // Request protocol
	//"%{VARNAME}i":   `(?P<request_header_value>.*?)`,                          // Contents of request header VARNAME
	"%k":    `(?P<keepalive_requests>\d+)`, // Number of keepalive requests handled
	"%l":    `(?P<remote_logname>\S+)`,     // Remote logname
	"%L":    `(?P<request_log_id>\S+)`,     // Request log ID
	"%{c}L": `(?P<connection_log_id>\S+)`,  // Connection log ID
	"%m":    `(?P<request_method>\S+)`,     // Request method
	//"%{VARNAME}n":   `(?P<note_value>.*?)`,                                    // request_timeContents of note VARNAME
	//"%{VARNAME}o":   `(?P<response_header_value>.*?)`,                         // Contents of response header VARNAME
	"%p":            `(?P<canonical_port>\d+)`,                                // Canonical port of the server
	"%{FORMAT}p":    `(?P<port_value>\S+)`,                                    // Canonical port, local port, or remote port
	"%P":            `(?P<process_id>\d+)`,                                    // Process ID of the child
	"%{FORMAT}P":    `(?P<process_thread_id>\S+)`,                             // Process ID or thread ID
	"%q":            `(?P<query_string>\S+)`,                                  // Query string
	"%r":            `(?P<method>\w+) (?P<url>.*?) HTTP\/(?P<version>\d\.\d)`, // First line of request
	"%R":            `(?P<handler_response>\S+)`,                              // Handler generating the response
	"%s":            `(?P<status>\d+)`,                                        // Status
	"%t":            `\[(?P<request_time>.*?)\]`,                              // Time the request was received
	"%{FORMAT}t":    `(?P<time_format_value>\S+)`,                             // Time in specified format
	"%T":            `(?P<request_time_seconds>\d+)`,                          // Time taken to serve the request (seconds)
	"%{UNIT}T":      `(?P<request_time_unit>\d+)`,                             // Time taken in specified unit
	"%u":            `(?P<remote_user>\S+)`,                                   // Remote user
	"%U":            `(?P<url_path>\S+)`,                                      // URL path requested
	"%v":            `(?P<server_name>\S+)`,                                   // ServerName of the server
	"%V":            `(?P<canonical_server_name>\S+)`,                         // Server name according to UseCanonicalName setting
	"%X":            `(?P<connection_status>[\+\-X])`,                         // Connection status
	"%I":            `(?P<bytes_received>\d+)`,                                // Bytes received
	"%O":            `(?P<bytes_sent>\d+)`,                                    // Bytes sent
	"%S":            `(?P<bytes_transferred>\d+)`,                             // Bytes transferred
	"%{VARNAME}^ti": `(?P<trailer_request_value>.*?)`,                         // Contents of trailer line(s) in the request
	"%{VARNAME}^to": `(?P<trailer_response_value>.*?)`,                        // Contents of trailer line(s) in the response
	"%>s":           `(?P<status>\d+)`,                                        // status code
}

// namedGroupContentRegex is regex to capture the content of the named group in the given regex itself
var namedGroupContentRegex = regexp.MustCompile(`\(\?P?\<.*?\>(.*?)\)`)
var perlCaptureGroupNameRegex = regexp.MustCompile(`\?\<(\w+)\>`)

func ParseLogFormat(format string) string {
	for placeholder, pattern := range placeholderPatterns {
		format = strings.ReplaceAll(format, placeholder, pattern)
	}

	return handleSpecialCase(format)
}

func handleSpecialCase(apacheFormat string) string {
	tokens := strings.Split(apacheFormat, " ")

	for i, token := range tokens {
		// handle edge cases
		edgeCases := []string{"}i", "}e", "}C", "}o", "}n"}
		if slices.ContainsFunc(edgeCases, func(s string) bool {
			return strings.Contains(token, s)
		}) {
			var groupName string
			// extract the groupName
			fI := strings.Index(token, "{")
			lI := strings.Index(token, "}")

			if fI != -1 && lI != -1 {
				groupName = formatGroupName(token[(fI + 1):lI])
				tokens[i] = token[:(fI-1)] + fmt.Sprintf("(?P<%s>.*?)", groupName) + token[(lI+2):]
			}
		}
	}

	return strings.Join(tokens, " ")
}

func formatGroupName(groupName string) string {
	return strings.ReplaceAll(groupName, "-", "_")
}

func SanitizeRegex(rgxStr string) string {
	// Replace the ?<> perl syntax capture grouped with ?P<> python syntax capture grouped
	// This is required as typescript don't support ?P<> syntax and user would require to use ?<> syntax.
	// However, our agent uses the ?P<> syntax. Hence, we need to perform this operation.
	rgxStr = perlCaptureGroupNameRegex.ReplaceAllStringFunc(rgxStr, func(match string) string {
		submatches := perlCaptureGroupNameRegex.FindStringSubmatch(match)
		if len(submatches) >= 2 {
			return "?P<" + submatches[1] + ">"
		}
		return match
	})

	// Remove the named captured group from the regex for the "if" and "value" field.
	rgxWithoutNamedCaptureGroup := namedGroupContentRegex.ReplaceAllStringFunc(rgxStr, func(match string) string {
		submatches := namedGroupContentRegex.FindStringSubmatch(match)
		if len(submatches) >= 2 {
			return fmt.Sprintf("(%s)", submatches[1])
		}

		return match
	})

	return rgxWithoutNamedCaptureGroup
}

func doubleEscapeRegex(rgxStr string) string {
	x := `dDwWsS.[]()+-^$|?*/`

	for _, c := range x {
		rgxStr = strings.ReplaceAll(rgxStr, `\`+string(c), `\\`+string(c))
	}

	return rgxStr
}
